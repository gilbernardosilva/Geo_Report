import React, { useState, useEffect } from 'react';
import './index.css';
import { Container, Modal, Form, Button, Carousel } from 'react-bootstrap';
import { MapContainer, TileLayer, Marker, Popup, useMapEvents } from 'react-leaflet'
import L from "leaflet";
import CustomNavbar from './../Navbar';
import "leaflet/dist/leaflet.css";
import { useTranslation } from "react-i18next";
import iconGreen from "../../img/icons/marker_green.svg"
import iconBlue from "../../img/icons/marker_blue.svg"
import iconRed from "../../img/icons/marker_red.svg"
import iconOrange from "../../img/icons/marker_orange.svg"
import icon from "leaflet/dist/images/marker-icon.png";
import iconShadow from "leaflet/dist/images/marker-shadow.png";
import { useAuth } from '../../hooks/AuthContext';
import { useAxiosWithToken } from '../../utils/api'
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import IssueDetailsModal from '../IssuesDetailModal';

const statusColorMap = {
    0: iconOrange, // Pending
    1: iconBlue,   // In Progress
    2: iconGreen,  // Completed
    3: iconRed,    // Rejected
};


function Dashboard() {
    const [filterType, setFilterType] = useState('');
    const [previewImages, setPreviewImages] = useState([]);
    const { userInfo } = useAuth();
    const [issues, setIssues] = useState([]);
    const [userPosition, setUserPosition] = useState(null);
    const [showModal, setShowModal] = useState(false);
    const [options, setOptions] = useState([]);
    const [statuses, setStatuses] = useState([]);
    const api = useAxiosWithToken();
    const [selectedIssue, setSelectedIssue] = useState(null);
    const [showIssueModal, setIssueShowModal] = useState(false);


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
                setFormData({
                    latitude: e.latlng.lat,
                    longitude: e.latlng.lng,
                });
                setShowModal(true);
            },
        });
        return null;
    }
    const handleFilterChange = (event) => {
        setFilterType(event.target.value);
        fetchIssues(event.target.value);
        console.log(event.target.value);
    };

    const fetchIssues = async (statusId) => {
        try {
            const params = {
                user_id: userInfo.id,
                report_status_id: statusId

            };

            const response = await api.get("report/user/" + userInfo.id, { params });
            if (response.status === 200) {
                setIssues(response.data.reports);
            } else {
                toast.error("Error fetching issues");
            }
        } catch (err) {
            toast.error("Error fetching issues");
        }
    };

    useEffect(() => {

        if (userInfo) {
            fetchIssues();
        }

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
                        toast.error("Error fetching options");
                    }
                } catch (err) {
                    toast.error("Error fetching options");
                }
            };

            fetchOptions();
        }


        const fetchStatus = async () => {
            try {
                const response = await api.get("reportStatus");
                if (response.status === 200) {
                    const filteredStatus = response.data.report_status.filter(
                        (status) => status.id !== 5)
                    setStatuses(filteredStatus);
                } else {
                    toast.error("Error fetching status");
                }
            } catch (err) {
                toast.error("Error fetching status");
            }
        };

        fetchStatus();

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
    }, [userInfo, api]);

    const customIcon = (statusID) => {
        const isCustomIcon = statusID in statusColorMap;
        return L.icon({
            iconUrl: statusColorMap[statusID] || icon,
            shadowUrl: iconShadow,
            iconSize: isCustomIcon ? [50, 82] : [25, 41],
            iconAnchor: [12, 41],
            popupAnchor: [1, -34],
            shadowSize: [41, 41],
        });
    };

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
        debugger;
        event.preventDefault();
        try {
            const response = await api.post("/report", formData);
            if (response.status === 200) {
                toast.success(t("reportSentSuccessfully"));
                setPreviewImages([]);
                setShowModal(false);
                fetchIssues();
            } else {
                toast.error(t("errorSendingReport"));
            }
        } catch (error) {
            toast.error(t("errorSendingReport"));
        }
    };

    const { t } = useTranslation();

    const handleCloseIssueModal = () => {
        setIssueShowModal(false);
    };

    const handleViewDetails = (issue) => {
        setSelectedIssue(issue);
        setIssueShowModal(true);
    };
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
                        <Marker position={userPosition} icon={customIcon(50)}>
                            <Popup>
                                {t("youAreHere")} (Lat: {userPosition[0]}, Lng: {userPosition[1]})
                            </Popup>
                        </Marker>
                    )}
                    {issues.map((issue) => (
                        <Marker
                            key={issue.id}
                            position={[issue.latitude, issue.longitude]}
                            icon={customIcon(issue.report_status.id)}
                            eventHandlers={{
                                click: (e) => {
                                    handleViewDetails(issue)
                                }
                            }}
                        >
                        </Marker>
                    ))}
                    <MapEvents />
                </MapContainer>
                <Form.Group controlId="report_status_id"
                    style={{ position: "absolute", top: 80, right: 50, zIndex: 1000 }}
                >
                    <Form.Select value={filterType} onChange={handleFilterChange}>
                        <option value="">{t("all")}</option>
                        {statuses.map((option) => (
                            <option key={option.id} value={option.id}>
                                {option.status}
                            </option>
                        ))}
                    </Form.Select>
                </Form.Group>
                <IssueDetailsModal
                    issue={selectedIssue}
                    show={showIssueModal}
                    handleClose={handleCloseIssueModal}
                />
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