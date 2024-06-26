import React, { useState, useEffect } from "react";
import { Card, Row, Col, Button } from "react-bootstrap";
import { useAxiosWithToken } from "../../utils/api.js";
import CustomNavbar from './../Navbar';
import { useTranslation } from "react-i18next";
import cardIcon from './../../img/icons/cardIssues.svg'
import IssueDetailsModal from "../IssuesDetailModal/index.jsx";
import { useAuth } from "../../hooks/AuthContext.jsx";
import { formatDistanceToNow } from 'date-fns';
import "./index.css";


function MyIssues() {
    const [issues, setIssues] = useState([]);
    const api = useAxiosWithToken();
    const [selectedIssue, setSelectedIssue] = useState(null);
    const [showModal, setShowModal] = useState(false);
    const { userInfo } = useAuth();
    const [error, setError] = useState(null);

    useEffect(() => {
        debugger;
        const fetchIssues = async () => {
          try {
            const response = await api.get("report/user/" + userInfo.id); 
            if (response.status === 200) {
              setIssues(response.data.reports);
              console.log(response.data.reports)
            } else {
              setError("Error fetching issues");
            }
          } catch (err) {
            setError("Error fetching issues");
          }
        };
    
        if (userInfo) {
          fetchIssues();
        }
      }, [userInfo]);

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
            <div className="d-flex flex-column align-items-center vh-100"> 

            <div className="d-flex justify-content-center flex-wrap gap-3">  
            {issues.map((issue) => (
                    <Card key={issue.id} className="dark-card" style={{ maxWidth: "80%" }}>
                        <Card.Header className="d-flex justify-content-between align-items-center">
                            <img src={cardIcon} alt="" className="img-fluid logo-icon" />
                            <Button disabled variant="info">
                                {issue.report_status.status}
                            </Button>
                        </Card.Header>
                        <Card.Body>
                            <Row>
                                <Col>
                                    <Card.Title>
                                        {t("name")}: {issue.name || "Not provided"}
                                    </Card.Title>
                                    <Card.Text>
                                        {t("type")}: {issue.report_type.name || "Not provided"}
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
                            {formatDistanceToNow(new Date(issue.report_date))+  " ago" || "2 days ago"}
                        </Card.Footer>
                    </Card>

                ))}

            </div>
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
