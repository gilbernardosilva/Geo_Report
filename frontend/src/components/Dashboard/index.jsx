import React, { useState } from 'react';
import './index.css';
import { Container, Col } from 'react-bootstrap';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet'
import CustomNavbar from './../Navbar';
import "leaflet/dist/leaflet.css";
import { useTranslation } from "react-i18next";



function Dashboard({ setToken }) {

    const { t } = useTranslation();
    return (
        <>
            <CustomNavbar setToken={setToken} t={t}></CustomNavbar>
            <Container fluid>

                        <MapContainer
                            center={[51.505, -0.09]}
                            zoom={13}
                            scrollWheelZoom={true}>
                            <TileLayer
                                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                            />
                            <Marker position={[51.505, -0.09]}>
                                <Popup>
                                    A pretty CSS3 popup. <br /> Easily customizable.
                                </Popup>
                            </Marker>
                        </MapContainer>

            </Container>
        </>
    );
}

export default Dashboard;