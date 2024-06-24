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
    useEffect(() => {
        
        const fetchData = async () => {
            if(token !== null){
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
        fetchData();
    }, [token]);


    const donutChartData1 = {
        labels: [
            'Potholes',
            'Malfunctioning Streetlights',
            'Damaged Traffic Signs',
            'Road Damage',
            'Blocked or Clogged Drainage',
            'Other'
        ],
        datasets: [{
            label: 'Total',
            data: [10, 15, 20, 5, 6, 2],
            backgroundColor: [
                'rgb(38, 70, 83)',
                'rgb(42, 157, 143)',
                'rgb(138, 177, 125)',
                'rgb(233, 196, 106)',
                'rgb(244, 162, 97)',
                'rgb(226, 71, 32)'
            ],
            hoverOffset: 4
        }]
    };

    const donutChartData2 = {
        labels: [
            'Reported',
            'Being Assessed',
            'Being Resolved',
            'Resolved',
            'Closed'
        ],
        datasets: [{
            label: 'Total',
            data: [10, 15, 20, 5, 6],
            backgroundColor: [
                'rgb(38, 70, 83)',
                'rgb(42, 157, 143)',
                'rgb(138, 177, 125)',
                'rgb(233, 196, 106)',
                'rgb(244, 162, 97)'
            ],
            hoverOffset: 4
        }]
    };

    const lineLabels = [
        'January',
        'February',
        'March',
        'April',
        'May',
        'June',
        'July',
        'August',
        'September',
        'October',
        'November',
        'December'
    ]
    const lineData = {
        labels: lineLabels,
        datasets: [{
            label: 'Issues',
            data: [65, 59, 80, 81, 56, 55, 40, 66, 90, 150, 100, 87],
            fill: false,
            borderColor: 'rgb(75, 192, 192)',
            tension: 0.1
        }]
    };

    return (
        <>
            {error ? (
                <ErrorHandler error={error} />
            ) : (
                <>
                    <CustomAuthorityNavbar setToken={setToken} t={t}></CustomAuthorityNavbar>
                    <Row className="justify-content-center mt-4 dashboard-row row">
                        <Col md={12} className='justify-content-center'>
                            <h1 className='text-center' style={{ color: 'white' }}>
                                {t('authpanel')}
                            </h1>
                        </Col>
                    </Row>
                    <Container fluid className="dashboard-container">
                        <Row className="justify-content-center mt-4 dashboard-row">
                            <Col xs={12} md={5} >
                                <Container className='h-auto'>
                                    <Typography sx={{ fontSize: 20, marginTop: 3 }} color="text.primary" gutterBottom>
                                        {t('totalIssues')}
                                    </Typography>
                                    <Line
                                        data={lineData}
                                        options={{
                                            scales: {
                                                y: {
                                                    beginAtZero: true
                                                }
                                            },
                                            title: {
                                                display: true,
                                                text: t('typesIssues')
                                            },
                                            legend: {
                                                position: 'bottom',
                                            },
                                            responsive: true,

                                        }}
                                    />
                                </Container>
                            </Col>
                        </Row>
                        <Row className="justify-content-center mt-4 dashboard-row">
                            <Col xs={12} md={4}>
                                <Container>
                                    <Typography sx={{ fontSize: 20, marginTop: 3 }} color="text.primary" gutterBottom>
                                        {t('typesIssues')}
                                    </Typography>
                                    <Doughnut
                                        data={donutChartData1}
                                        options={{
                                            plugins: {
                                                title: {
                                                    display: true,
                                                    text: t('last7days')
                                                },
                                                legend: {
                                                    position: 'bottom',
                                                }
                                            },
                                            responsive: true,
                                        }}
                                    />
                                </Container>
                            </Col>
                            <Col xs={12} md={4}>
                                <Container>
                                    <Typography sx={{ fontSize: 20, marginTop: 3 }} color="text.primary" gutterBottom>
                                        {t('statusIssues')}
                                    </Typography>
                                    <Doughnut
                                        data={donutChartData2}
                                        options={{
                                            plugins: {
                                                title: {
                                                    display: true,
                                                    text: t('last7days')
                                                },
                                                legend: {
                                                    position: 'bottom',
                                                }
                                            },
                                            responsive: true,
                                        }}
                                    />
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