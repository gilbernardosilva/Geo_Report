import React, { useEffect, useState } from 'react';
import './index.css';
import { Container, Row, Col } from 'react-bootstrap';
import Typography from '@mui/material/Typography';
import "leaflet/dist/leaflet.css";
import { useTranslation } from 'react-i18next';

import {
    Chart as ChartJS,
    ArcElement,
    Tooltip,
    Legend,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title
} from 'chart.js'
import { Doughnut, Line } from 'react-chartjs-2'
import { useAuth } from "./../../hooks/AuthContext.jsx";
import { useAxiosWithToken } from "./../../utils/api.js";
import CustomAuthorityNavbar from './../AuthorityNavbar';
import ErrorHandler from './../ErrorHandler/index.jsx';
import PropTypes from 'prop-types';
import axios from 'axios';

ChartJS.register(
    ArcElement,
    Tooltip,
    Legend,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title
)

function AuthorityDashboard({ setToken }) {
    const { t } = useTranslation();
    const [error, setError] = useState(null);
    const { token } = useAuth();
    const api = useAxiosWithToken();

    const [statusData, setStatusData] = useState({
        labels: [],
        datasets: [{
            label: 'Total',
            data: [],
            backgroundColor: [],
            hoverOffset: 4
        }]
    });

    const [typesData, setTypesData] = useState({
        labels: [],
        datasets: [{
            label: 'Total',
            data: [],
            backgroundColor: [],
            hoverOffset: 4
        }]
    });

    const [loading, setLoading] = useState(true);

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

                    const statusResponse = await api.get('report/status/chart');
                    setStatusData(statusResponse.data);
                    console.log(statusResponse.data)

                    const typesResponse = await api.get('report/types/chart');
                    setTypesData(typesResponse.data);
                    console.log(typesData)

                } catch (error) {
                    setError(error);
                    console.error("Error fetching data:", error);
                } finally {
                    setLoading(false);
                }
            }
        };
        fetchData();
    }, [token]);


    return (
        <>
            {error ? (
                <ErrorHandler error={error} />
            ) : (
                <>
                    <CustomAuthorityNavbar setToken={setToken} t={t}></CustomAuthorityNavbar>
                    <Row style={{ backgroundColor: '#181b1d' }}>
                        <Col md={12} className='justify-content-center'>
                            <h1 className='text-center' style={{ color: 'white' }}>
                                {t('authpanel')}
                            </h1>
                        </Col>
                    </Row>
                    <Container fluid className="dashboard-container">
                        <Row className="justify-content-center mt-4 dashboard-row">
                            <Col xs={12} md={4}>
                                <Container>
                                    <Typography sx={{ fontSize: 20, marginTop: 3 }} color="text.primary" gutterBottom>
                                        {t('typesIssues')}
                                    </Typography>
                                    {typesData.length > 0 && (
                                        <Doughnut
                                            data={typesData}
                                            options={{
                                                plugins: {
                                                    legend: {
                                                        position: "bottom",
                                                    },
                                                },
                                                responsive: true,
                                            }}
                                        />
                                    )}
                                </Container>
                            </Col>
                            <Col xs={12} md={4}>
                                <Container>
                                    <Typography sx={{ fontSize: 20, marginTop: 3 }} color="text.primary" gutterBottom>
                                        {t('statusIssues')}
                                    </Typography>
                                    {statusData.length > 0 && (

                                    <Doughnut
                                        data={statusData}
                                        options={{
                                            plugins: {
                                                legend: {
                                                    position: 'bottom',
                                                }
                                            },
                                            responsive: true,
                                    }}
                                    />
                                )}
                                </Container>
                            </Col>
                        </Row>
                    </Container>
                </>

            )}
        </>
    );
}

AuthorityDashboard.propTypes = {
    setToken: PropTypes.func.isRequired,
}

export default AuthorityDashboard;