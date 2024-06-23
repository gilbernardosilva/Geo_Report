import React, { useState, useEffect } from "react";
import { Card, Container, Row, Col, Button, Carousel } from "react-bootstrap";
import { useAuth } from "../../hooks/AuthContext";
import { useAxiosWithToken } from "../../utils/api.js";
import CustomNavbar from './../Navbar';
import { useTranslation } from "react-i18next";
import cardIcon from './../../img/icons/cardIssues.svg'
import IssueDetailsModal from "../IssuesDetailModal/index.jsx";
import "./index.css";


function MyIssues() {
    const { userInfo } = useAuth();
    const [issues, setIssues] = useState([]);
    const [error, setError] = useState(null);
    const api = useAxiosWithToken();
    const [selectedIssue, setSelectedIssue] = useState(null);
    const [showModal, setShowModal] = useState(false);
    useEffect(() => {
        // Create some sample issue data
        const mockIssues = [
            {
                id: 1,
                title: "Pothole on Main Street",
                description:
                    "Large pothole causing damage to vehicles. Located near 123 Main Street.",
                status: "Pending",
                imageUrls: ["https://via.placeholder.com/300"], // Sample image URL
            },
            {
                id: 2,
                title: "Broken Streetlight on Elm Street",
                description: "Streetlight not working at the corner of Elm and Oak Street.",
                status: "In Progress",
            },
        ];

        setIssues(mockIssues); // Update the state with the mock data
    }, []);

    const handleViewDetails = (issue) => {
        setSelectedIssue(issue); // Store the selected issue in state
        setShowModal(true); // Show the modal
    };

    const handleCloseModal = () => setShowModal(false);
    const { t } = useTranslation();

    return (
        <>
            <CustomNavbar t={t}></CustomNavbar>
            <h2 className="text-center mt-2" style={{ color: 'white', fontSize: '2rem' }}>My Issues</h2>
            <div className="align-items-center">
                {issues.map((issue) => (
                    <Card key={issue.id} className="dark-card">
                        <Card.Header className="d-flex justify-content-between align-items-center">
                            <img src={cardIcon} alt="" className="img-fluid logo-icon" />
                            <Button disabled variant="info">
                                {issue.status}
                            </Button>
                        </Card.Header>
                        <Card.Body>
                            <Row>
                                <Col>
                                    {/* Issue Details */}
                                    <Card.Title>
                                        {t("name")}: {issue.title || "Not provided"}
                                    </Card.Title>
                                    <Card.Text>
                                        {t("type")}: {issue.type || "Not provided"}
                                    </Card.Text>
                                </Col>

                                <Col className="d-flex justify-content-end">
                                    <Button variant="primary" onClick={() => handleViewDetails(issue)}>
                                        {t("viewDetails")}
                                    </Button>
                                </Col>
                            </Row>
                        </Card.Body>
                        <Card.Footer style={{ color: "grey" }} className="text-center">
                            {issue.timestamp || "2 days ago"}
                        </Card.Footer>
                    </Card>

                ))}

            </div>
            <IssueDetailsModal
                issue={selectedIssue}
                show={showModal}
                handleClose={handleCloseModal}
            />
        </>
    );
}

export default MyIssues;
