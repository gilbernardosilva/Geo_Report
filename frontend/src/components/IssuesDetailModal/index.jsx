import React from "react";
import { Modal, Button, Carousel } from "react-bootstrap";
import PropTypes from 'prop-types';
import { useTranslation } from "react-i18next";

function IssueDetailsModal({ issue, show, handleClose }) {
    const {t} = useTranslation();
    return (
      <Modal show={show} onHide={handleClose}>
        {issue && ( // Only render if issue is not null or undefined
          <>
            <Modal.Header closeButton>
              <Modal.Title>{issue.name}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
              <p>
                <b>Description:</b> {issue.description}
              </p>
              <p>
                <b>Status:</b> {issue.report_status.name}
              </p>
              <p>
                <b>Latitude:</b> {issue.latitude}
              </p>
              <p>
                <b>Longitude:</b> {issue.longitude}
              </p>
              <p style={{ color: 'white' }}>
            <b>{t('photos')}:</b> 
            {issue?.photos && issue.photos.length > 0 ? (
                <Carousel variant="dark" interval={null}>
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
                <p>{t('noPhotos')}</p> 
            )}
        </p>
              Created by: {issue.user.first_name} {issue.user.last_name}

            </Modal.Body>
            <Modal.Footer>
              <Button variant="secondary" onClick={handleClose}>
                Close
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
