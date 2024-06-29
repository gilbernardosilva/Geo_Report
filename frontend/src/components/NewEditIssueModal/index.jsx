import React, { useState, useEffect } from "react";
import { Modal, Button, Form, Carousel } from "react-bootstrap";
import { MapContainer, TileLayer } from "react-leaflet";
import { useTranslation } from "react-i18next";
import L from "leaflet";
import "leaflet/dist/leaflet.css";
import { toast } from "react-toastify";
import { useAxiosWithToken } from "../../utils/api";
import LocationMarker from "../LocationMarker";
import { useAuth } from "../../hooks/AuthContext";


function NewEditIssueModal({ issue, show, editMode, handleClose }) {
    const { userInfo } = useAuth();
    const { t } = useTranslation();
    const [options, setOptions] = useState([]);
    const api = useAxiosWithToken();
    const [formData, setFormData] = useState({
        name: "",
        report_type_id: "",
        description: "",
        latitude: 0,
        longitude: 0,
    });

    const [previewImages, setPreviewImages] = useState([]);

    const [mapPosition, setMapPosition] = useState(
        issue
            ? { lat: issue.latitude, lng: issue.longitude }
            : { lat: 0, lng: 0 }
    );

    useEffect(() => {
        if (issue) {
            setFormData(issue);
            setMapPosition({
                lat: issue.latitude,
                lng: issue.longitude,
            });
        }
    }, [issue]);

    const handleChange = (event) => {
        if (event.target.files) {
            const selectedFiles = Array.from(event.target.files);

            const imageUrls = [];
            selectedFiles.forEach((file, index) => {
                const reader = new FileReader();
                reader.onload = (e) => {
                    imageUrls.push(e.target.result);

                    if (imageUrls.length === selectedFiles.length) {
                        setPreviewImages(imageUrls);

                        setFormData({
                            ...formData,
                            photos: imageUrls,
                        });
                    }
                };
                reader.readAsDataURL(file);
            });
        } else {
            setFormData({
                ...formData,
                [event.target.id]: event.target.id === 'report_type_id' ? parseInt(event.target.value) : event.target.value,
                user_id: userInfo?.id || 0,

            });
        }
    };

    const handleSubmit = async (event) => {
        event.preventDefault();



        try {
            const endpoint = editMode ? `report` : "admin/report";
            const method = editMode ? api.put : api.post;

            const response = await method(endpoint, formData);

            if (response.status === 200) {
                toast.success(editMode ? "Issue updated successfully" : "Issue reported successfully");
                setFormData({ name: "", description: "" });
                setPreviewImages([]);
                handleClose();
            } else {
                toast.error(editMode ? "Failed to update issue" : "Failed to report issue");
                console.error("Error:", response.data);
            }
        } catch (error) {
            toast.error("An error occurred during submission:", error);
        }
    };

    useEffect(() => {
        if (options.length === 0) {
            const fetchOptions = async () => {
                try {
                    const response = await api.get("reportTypes");
                    if (response.status === 200) {
                        setOptions(response.data.report_types);
                        localStorage.setItem("reportTypes", JSON.stringify(response.data.report_types));
                    } else {
                        toast.error("Error fetching options");
                    }
                } catch (err) {
                    toast.error("Error fetching options");
                }
            };

            fetchOptions();
        }
    });


    return (
        <Modal className="modal-container" show={show} onHide={handleClose}>
            <Modal.Header closeButton>
                <Modal.Title>{t("reportIssue")}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={handleSubmit}>
                    <Form.Group controlId="name" className="mb-3">
                        <Form.Label>{t("title")}</Form.Label>
                        <Form.Control required
                            type="text"
                            value={formData.name}
                            onChange={handleChange}
                        />
                    </Form.Group>
                    <Form.Group controlId="report_type_id" className="mb-3">
                        <Form.Label>{t("typeProblem")}</Form.Label>
                        <Form.Select required
                            value={formData.report_type_id}
                            onChange={handleChange}
                        >
                            <option value="">{t("selectAnOption")}</option>
                            {options.map((option) => (
                                <option key={option.id} value={option.id}>
                                    {option.name}
                                </option>
                            ))}
                        </Form.Select>
                    </Form.Group>

                    <Form.Group controlId="latitude" className='mb-3'>
                        <Form.Label>Latitude</Form.Label>
                        <Form.Control required
                            type="text"
                            value={formData.latitude}
                            onChange={handleChange}
                        />
                    </Form.Group>

                    <Form.Group controlId="longitude" className="mb-3">
                        <Form.Label>Longitude</Form.Label>
                        <Form.Control required
                            type="text"
                            value={formData.longitude}
                            onChange={handleChange}
                        />
                    </Form.Group>

                    <Form.Group controlId="description" className="mb-3">
                        <Form.Label>{t("description")}</Form.Label>
                        <Form.Control
                            as="textarea"
                            rows={3}
                            value={formData.description}
                            onChange={handleChange}
                        />
                    </Form.Group>

                    <Form.Group controlId="formFileMultiple" className="mb-3">
                        <Form.Label>Images</Form.Label>
                        <Form.Control type="file" multiple capture="camera" accept="image/*"
                            onChange={handleChange} />
                    </Form.Group>

                    {previewImages.length > 0 && (
                        <Carousel className="mb-2">
                            {previewImages.map((image, index) => (
                                <Carousel.Item key={index}>
                                    <img
                                        className="d-block w-100 img-fluid"
                                        src={image}
                                        alt={`Preview ${index + 1}`}
                                    />
                                </Carousel.Item>
                            ))}
                        </Carousel>
                    )}
                    <MapContainer
                        center={[formData.latitude || 0, formData.longitude || 0]}
                        zoom={13}
                        scrollWheelZoom={false}
                        style={{ height: "300px", width: "100%" }}
                    >
                        <TileLayer
                            attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                        />
                        {mapPosition && (
                            <LocationMarker setFormData={setFormData} formData={formData} setPosition={setMapPosition} />
                        )}
                    </MapContainer>
                    <Button variant="primary" type="submit">
                        Submit
                    </Button>
                </Form>
            </Modal.Body>
        </Modal>
    );
}

export default NewEditIssueModal;