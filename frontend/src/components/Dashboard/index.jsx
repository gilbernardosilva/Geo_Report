import React, { useState, useEffect } from 'react';
import './index.css';
import { Container } from 'react-bootstrap';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet'
import L from "leaflet";
import CustomNavbar from './../Navbar';
import "leaflet/dist/leaflet.css";
import { useTranslation } from "react-i18next";
import icon from "leaflet/dist/images/marker-icon.png";
import iconShadow from "leaflet/dist/images/marker-shadow.png";


function Dashboard({ setToken }) {

    const [userPosition, setUserPosition] = useState(null);

    useEffect(() => {
        function getLocation() {
            if (navigator.geolocation) {
                navigator.geolocation.watchPosition(showPosition, showError); // Use watchPosition to track changes
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

    const { t } = useTranslation();
    return (
        <>
            <CustomNavbar setToken={setToken} t={t}></CustomNavbar>
            <Container fluid>

                <MapContainer
                    center={userPosition || [41.1561925, -8.5171909]}
                    zoom={13}
                    scrollWheelZoom={true}>
                    <TileLayer
                        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                    />
                    {userPosition && ( // Render the marker only if userPosition is available
                        <Marker position={userPosition} icon={customIcon}>
                            <Popup>
                                You are here! (Lat: {userPosition[0]}, Lng: {userPosition[1]})
                            </Popup>
                        </Marker>
                    )}
                </MapContainer>

            </Container>
        </>
    );
}

export default Dashboard;