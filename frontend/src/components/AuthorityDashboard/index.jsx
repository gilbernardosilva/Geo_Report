import React, { useEffect, useState } from 'react';
import './index.css';
import { Container, Form } from 'react-bootstrap';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet'
import L from "leaflet";
import { useTranslation } from "react-i18next";
import iconGreen from "../../img/icons/marker_green.svg"
import iconBlue from "../../img/icons/marker_blue.svg"
import iconRed from "../../img/icons/marker_red.svg"
import iconOrange from "../../img/icons/marker_orange.svg"
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


const statusColorMap = {
    0: iconOrange, // Pending
    1: iconBlue,   // In Progress
    2: iconGreen,  // Completed
    3: iconRed,    // Rejected
  };


function AuthorityDashboard({ setToken }) {
    const [statuses, setStatuses] = useState([]);
    const [filterType, setFilterType] = useState('');
    const { t } = useTranslation();
    const [error,setError] = useState(null);
    const { token } = useAuth();
    const api = useAxiosWithToken();
    const [issues, setIssues] = useState([]);
    const [userPosition, setUserPosition] = useState(null);
    const [options, setOptions] = useState([]); 
    const [selectedIssue, setSelectedIssue] = useState(null);
    const [showIssueModal, setIssueShowModal] = useState(false);

    const fetchIssues = async (statusId) => {
        try {
            const params = {
                report_status_id : statusId

            };
            const response = await api.get("admin/report", {params});
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


    const handleCloseIssueModal = () => {
        setIssueShowModal(false);
    };
    const handleViewDetails = (issue) => {
        setSelectedIssue(issue);
        setIssueShowModal(true);
    };
    
    const handleFilterChange = (event) => {
        setFilterType(event.target.value);
        fetchIssues(event.target.value);
        console.log(event.target.value);
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
                        <Marker position={userPosition} icon={customIcon(60)}>
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
                </MapContainer>
                <Form.Group controlId="report_status_id"
                    style={{ position: "absolute", top: 80, right:50, zIndex: 1000 }}
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
            </Container>
    </>
    )
}

AuthorityDashboard.propTypes = {
    setToken: PropTypes.func.isRequired,
}

export default AuthorityDashboard;