import React from "react";
import { Modal, Button, Carousel, Form } from "react-bootstrap";
import PropTypes from 'prop-types';
import { useTranslation } from "react-i18next";
import { useAxiosWithToken } from "../../utils/api";
import { toast } from "react-toastify";
import { useAuth } from "../../hooks/AuthContext";
import { useState, useEffect } from "react";

function IssueDetailsModal({ issue, show, handleClose }) {
  const { userInfo } = useAuth();
  const { t } = useTranslation();
  const [options, setOptions] = useState([]);
  const api = useAxiosWithToken();
  const [formData, setFormData] = useState({
    name: issue?.name || "",
    report_type_id: issue?.report_type_id || "",
    description: issue?.description || "",
    report_status_id: issue?.report_status_id || "",
    latitude: issue?.latitude || 0,
    longitude: issue?.longitude || 0,
  });

  const handleStatusChange = (event) => {
    setFormData({
      ...formData,
      report_status_id: parseInt(event.target.value, 10),
    });
  };

  useEffect(() => {
    debugger;
    const fetchStatusOptions = async () => {
      try {
        const response = await api.get("reportStatus");
        setOptions(response.data.report_types);
      } catch (error) {
        console.error("Error fetching status options:", error);
        toast.error("Failed to fetch status options");
      }
    };
    if (userInfo.role === 1) {
      fetchStatusOptions();
    }

  }, [api]);

  const handleSubmit = async (event) => {
    event.preventDefault();

    const dataToSend = { report_status_id: formData.report_status_id };

    try {
      const response = await api.put(`admin/report/${issue.id}`, dataToSend);

      if (response.status === 200) {
        toast.success("Issue updated successfully");
        handleClose();
      } else {
        toast.error("Failed to update issue");
        console.error("Error:", response.data);
      }
    } catch (error) {
      toast.error("An error occurred during submission:", error);
    }
  };
  return (
    <Modal show={show} onHide={handleClose}>
      {issue && (
        <>
          <Modal.Header closeButton>
            <Modal.Title>{issue.name}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <p>
              <b>{t('description')}:</b> {issue.description}
            </p>
            {userInfo?.role !== 1 ? (
              <p>
                <b>{t('status')}:</b> {issue.report_status.name}
              </p>
            ) : (
              <p>
                <Form.Group controlId="report_status_id" className="mb-3">
                  <Form.Label>{t("status")}</Form.Label>
                  {options.length > 0 ? (
                    <Form.Select
                      value={formData.report_status_id}
                      onChange={handleStatusChange}
                    >
                      <option value={issue.report_status.id}>{issue.report_status.name}</option>
                      {options.map((option) => (
                        <option key={option.id} value={option.id}>
                          {option.name}
                        </option>
                      ))}
                    </Form.Select>
                  ) : (
                    <p>Loading status options...</p> 
                  )}
                </Form.Group>
              </p>
            )}
            <p>
              <b>Latitude:</b> {issue.latitude}
            </p>
            <p>
              <b>Longitude:</b> {issue.longitude}
            </p>
            <p style={{ color: 'white' }}>
              <b>{t('photos')}:</b>
              {issue?.photos && issue.photos.length > 0 ? (
                <Carousel className="mt-3" variant="dark" interval={null}>
                  {issue.photos.map((photo) => (
                    <Carousel.Item key={photo.id}>
                      <img
                        className="d-block w-100"
                        src={photo.Image}
                        alt={`Issue Photo ${photo.id}`}
                      />
                    </Carousel.Item>
                  ))}
                </Carousel>
              ) : (
                <p className="no-photos-message">{t('noPhotos')}</p>
              )}
            </p>
            {t('createdBy')}: {issue.user.first_name} {issue.user.last_name}

          </Modal.Body>
          <Modal.Footer>
          <Button variant="primary" onClick={handleSubmit}>
              {t('Save')}
            </Button>
            <Button variant="secondary" onClick={handleClose}>
              {t('close')}
            </Button>
          </Modal.Footer>
        </>
      )}
    </Modal>
  );
}

IssueDetailsModal.propTypes = {
  issue: PropTypes.object,
  show: PropTypes.bool.isRequired,
  handleClose: PropTypes.func.isRequired,
};

export default IssueDetailsModal;
