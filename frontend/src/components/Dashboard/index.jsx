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
import { useAuth } from '../../hooks/AuthContext';
import { useAxiosWithToken } from '../../utils/api'
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css"; 

function Dashboard({ }) {

    const [previewImages, setPreviewImages] = useState([]);
    const { userInfo, token } = useAuth();
    const [userPosition, setUserPosition] = useState(null);
    const [showModal, setShowModal] = useState(false);
    const [clickedCoords, setClickedCoords] = useState(null);
    const [options, setOptions] = useState([]);
    const [error, setError] = useState(null);
    const api = useAxiosWithToken();
    console.log(previewImages);

    const handleCloseModal = () => {
        setShowModal(false);
        setPreviewImages([]);
    };

    const [formData, setFormData] = useState({
        name: "",
        latitude: 0,
        longitude: 0,
        user_id: userInfo?.id || 0,
        description: "",
        report_type_id: 0,
        photos: [],
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
        const storedOptions = JSON.parse(localStorage.getItem("reportTypes"));
        if (storedOptions && storedOptions.length > 0) {
            setOptions(storedOptions);
        } else {
            const fetchOptions = async () => {
                try {
                    const response = await api.get("reportTypes");
                    if (response.status === 200) {
                        setOptions(response.data.report_types);
                        localStorage.setItem("reportTypes", JSON.stringify(response.data.report_types));
                    } else {
                        setError("Error fetching options");
                    }
                } catch (err) {
                    setError("Error fetching options");
                }
            };

            fetchOptions();
        }

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
          // Handling for text inputs
          setFormData({
            ...formData,
            [event.target.id]: event.target.id === 'report_type_id' ? parseInt(event.target.value) : event.target.value,
            user_id : userInfo?.id || 0,

          });
        }
      };
    

    const handleSubmit = async (event) => {
        debugger;
        event.preventDefault();
        try {
          const response = await api.post("/report", formData); 
          if (response.status === 200) {
            toast.success(t("reportSentSuccessfully")); 
            setShowModal(false);
          } else {
            toast.error(t("errorSendingReport"));
        }
        } catch (error) {
            toast.error(t("errorSendingReport"));
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
                                <Form.Control disabled required
                                    type="text"
                                    value={formData.latitude}
                                    onChange={handleChange}
                                />
                            </Form.Group>

                            <Form.Group controlId="longitude" className="mb-3">
                                <Form.Label>Longitude</Form.Label>
                                <Form.Control disabled required
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