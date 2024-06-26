import React, { useEffect, useState } from 'react';
import './index.css';
import { Container, Modal, Form, Button, Carousel } from 'react-bootstrap';
import { MapContainer, TileLayer, Marker, Popup, useMapEvents } from 'react-leaflet'
import L from "leaflet";
import "leaflet/dist/leaflet.css";
import { useTranslation } from "react-i18next";
import icon from "leaflet/dist/images/marker-icon.png";
import iconShadow from "leaflet/dist/images/marker-shadow.png";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import IssueDetailsModal from '../IssuesDetailModal';
import "leaflet/dist/leaflet.css";

import { useAuth } from "./../../hooks/AuthContext.jsx";
import { useAxiosWithToken } from "./../../utils/api.js";
import CustomAuthorityNavbar from './../AuthorityNavbar';
import PropTypes from 'prop-types';


function AuthorityDashboard({ setToken }) {
    const [previewImages, setPreviewImages] = useState([]);
    const { t } = useTranslation();
    const [error, setError] = useState(null);
    const { token } = useAuth();
    const api = useAxiosWithToken();
    const [issues, setIssues] = useState([]);
    const [userPosition, setUserPosition] = useState(null);
    const [showModal, setShowModal] = useState(false);
    const [options, setOptions] = useState([]);
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
        user_id: token?.id || 0,
        description: "",
        report_type_id: 0,
        photos: [],
    });

    useEffect(() => {

        const fetchData = async () => {
            if (token !== null) {
                try {
                    const response = await api.get("authority");
                    console.log(response);
                    if (response.status === 200) {
                        console.log("has authorization");
                    } else {
                        setError(response);
                    }
                } catch (error) {
                    setError(error);
                }
            }
        };

        const fetchIssues = async () => {
            try {
                const response = await api.get("admin/report");
                if (response.status === 200) {
                    setIssues(response.data.reports);
                    console.log(response.data.reports)
                } else {
                    toast.error("Error fetching issues");
                }
            } catch (err) {
                toast.error("Error fetching issues");
            }
        };


        fetchData();
        fetchIssues();

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
    }, [token, api]);

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
            setFormData({
                ...formData,
                [event.target.id]: event.target.id === 'report_type_id' ? parseInt(event.target.value) : event.target.value,
                user_id: token?.id || 0,

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

    const handleCloseIssueModal = () => {
        setIssueShowModal(false);
    };
    const handleViewDetails = (issue) => {
        setSelectedIssue(issue);
        setIssueShowModal(true);
    };
    
    return (
    <>
    <CustomAuthorityNavbar t={t}></CustomAuthorityNavbar>
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
                    {issues.map((issue) => (
                        <Marker
                            key={issue.id}
                            position={[issue.latitude, issue.longitude]}
                            icon={customIcon}
                            eventHandlers={{
                                click: (e) => {
                                    handleViewDetails(issue)
                                }
                            }}
                        >
                        </Marker>
                    ))}
                </MapContainer>
                <IssueDetailsModal
                    issue={selectedIssue}
                    show={showIssueModal}
                    handleClose={handleCloseIssueModal}
                />
            </Container>
    </>
    )
}

AuthorityDashboard.propTypes = {
    setToken: PropTypes.func.isRequired,
}

export default AuthorityDashboard;