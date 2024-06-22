import React, { useState, useEffect } from 'react';
import './index.css';
import { Container, Modal, Form, Button, Carousel } from 'react-bootstrap';
import { MapContainer, TileLayer, Marker, Popup, useMapEvents } from 'react-leaflet'
import L from "leaflet";
import CustomNavbar from './../Navbar';
import "leaflet/dist/leaflet.css";
import { useTranslation } from "react-i18next";
import icon from "leaflet/dist/images/marker-icon.png";
import iconShadow from "leaflet/dist/images/marker-shadow.png";



function Dashboard({ }) {

    const [previewImages, setPreviewImages] = useState([]);
    const [userPosition, setUserPosition] = useState(null);
    const [showModal, setShowModal] = useState(false);
    const [clickedCoords, setClickedCoords] = useState(null);


    const handleCloseModal = () => {
        setShowModal(false);
        setPreviewImages([]);
        setFormData({
            latitude: "",
            longitude: "",
            image: "",
            type: "",
        });
    };

    const [formData, setFormData] = useState({
        latitude: '',
        longitude: '',
        type: '',
        image: [],
    });

    function MapEvents() {
        const map = useMapEvents({
            click: (e) => {
                setClickedCoords([e.latlng.lat, e.latlng.lng]);
                setFormData({
                    latitude: e.latlng.lat,
                    longitude: e.latlng.lng,
                });
                setShowModal(true);
            },
        });
        return null;
    }

    useEffect(() => {
        function getLocation() {
            if (navigator.geolocation) {
                navigator.geolocation.watchPosition(showPosition, showError);
            } else {
                console.log("Geolocation is not supported by this browser.");
            }
        }

        function showPosition(position) {
            setUserPosition([position.coords.latitude, position.coords.longitude]);
        }

        function showError(error) {
            console.error("Error getting geolocation:", error);
        }

        getLocation();
    }, []);

    const customIcon = new L.Icon({
        iconUrl: icon,
        shadowUrl: iconShadow,
        iconSize: [25, 41],
        iconAnchor: [12, 41],
    });


    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            // Send formData to your backend here
            const response = await fetch('your_api_endpoint', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(formData)
            });

            if (response.ok) {
                // Handle success (e.g., close modal, update map)
                setShowModal(false);
            } else {
                // Handle error
            }
        } catch (error) {
            console.error('Error submitting form:', error);
        }
    };

    const { t } = useTranslation();
    return (
        <>
            <CustomNavbar t={t}></CustomNavbar>
            <Container fluid>

                <MapContainer center={[41.1561925, -8.5171909]} zoom={13} scrollWheelZoom={true}>
                    <TileLayer
                        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                    />
                    {userPosition && (
                        <Marker position={userPosition} icon={customIcon}>
                            <Popup>
                                {t("youAreHere")} (Lat: {userPosition[0]}, Lng: {userPosition[1]})
                            </Popup>
                        </Marker>
                    )}
                    <MapEvents />
                </MapContainer>
                <Modal className="modal-container" show={showModal} onHide={handleCloseModal}>
                    <Modal.Header closeButton>
                        <Modal.Title>Report Issue</Modal.Title> {/* Or customize the title */}
                    </Modal.Header>
                    <Modal.Body>
                        <Form onSubmit={handleSubmit}>

                            <Form.Group controlId="problem">
                                <Form.Label>Type of problem</Form.Label>
                                <Form.Select
                                    type="text"
                                    value={formData.type}
                                    onChange={(e) => setFormData({ ...formData, type: e.target.value })}
                                />
                            </Form.Group>

                            <Form.Group controlId="latitude">
                                <Form.Label>Latitude</Form.Label>
                                <Form.Control disabled
                                    type="text"
                                    value={formData.latitude}
                                    onChange={(e) => setFormData({ ...formData, latitude: e.target.value })}
                                />
                            </Form.Group>

                            <Form.Group controlId="longitude">
                                <Form.Label>Longitude</Form.Label>
                                <Form.Control disabled
                                    type="text"
                                    value={formData.longitude}
                                    onChange={(e) => setFormData({ ...formData, longitude: e.target.value })}
                                />
                            </Form.Group>

                            <Form.Group controlId="formFileMultiple" className="mb-3">
                                <Form.Label>Images</Form.Label>
                                <Form.Control type="file" multiple capture="camera" accept="image/*"
                                    onChange={(e) => {
                                        const files = e.target.files;
                                        const images = [];

                                        for (let i = 0; i < files.length; i++) {
                                            images.push(URL.createObjectURL(files[i]));
                                        }

                                        setPreviewImages(images);
                                        setFormData({ ...formData, image: e.target.value });
                                    }} />
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

                            <Button variant="primary" type="submit">
                                Submit
                            </Button>
                        </Form>
                    </Modal.Body>
                </Modal>
            </Container>
        </>
    );
}

export default Dashboard;