import logo from './../../img/logo/geo_report_no_text.png'
import { Container, Navbar, Offcanvas } from 'react-bootstrap';
import React, { useState } from 'react';
import home from './../../img/icons/home.svg'
import issues from './../../img/icons/issues.svg'
import logoutIcon from './../../img/icons/logout.svg'
import profile from './../../img/icons/profile.svg'
import settings from './../../img/icons/settings.svg'
import { Link } from 'react-router-dom';
import LanguageSwitcher from '../LanguageSwitcher/index.jsx';
import { useAuth } from '../../hooks/AuthContext';


function CustomNavbar({t}) {

    const { logout, userInfo } = useAuth();
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);
    const handleLanguageChange = (newAction) => {
        console.log(newAction);
    };
    return (
        <Navbar expand={false} bg="dark" variant="dark">
            <Container fluid>

                <Navbar.Toggle aria-controls={`offcanvasNavbar-expand-false`} onClick={handleShow} />

                <Navbar.Brand href="#">
                    <Link to="/dashboard"   >
                        <img src={logo} alt="Your Logo" className="img-fluid logo-icon" />
                    </Link>
                </Navbar.Brand>
                <Navbar.Offcanvas
                    id={`offcanvasNavbar-expand-false`}
                    aria-labelledby={`offcanvasNavbarLabel-expand-false`}
                    show={show}
                    onHide={handleClose}
                >
                    <Offcanvas.Header closeButton className="bg-dark">
                        <Offcanvas.Title style={{ color: 'white', fontSize: '2rem' }} id={`offcanvasNavbarLabel-expand-lg`}>
                            {t("Welcome")} {userInfo?.firstName} 
                        </Offcanvas.Title>
                    </Offcanvas.Header>
                    <Offcanvas.Body className="bg-dark d-flex flex-column">
                        <div className="sidebar-content flex-column align-items-start flex-grow-1">
                            <Link to="/dashboard" className="navbar-link text-end btn d-flex align-items-center" style={{ textDecoration: 'none', color: 'white' }}>
                                <img src={home} alt="" className="img-fluid logo-icon me-2" />
                                <p  className='navbar-p hover-effect'>{t('home')}</p>
                            </Link>
                            <Link to="/issues" className="navbar-link text-end btn d-flex align-items-center" style={{ textDecoration: 'none', color: 'white' }}>
                                <img src={issues} alt="" className="img-fluid logo-icon me-2" />
                                <p  className='navbar-p hover-effect'>{t('myIssues')}</p>
                            </Link>
                            <Link to="/profile" className="navbar-link text-end btn d-flex align-items-center" style={{ textDecoration: 'none', color: 'white' }}>
                                <img src={profile} alt="" className="img-fluid logo-icon me-2" />
                                <p  className='navbar-p hover-effect'>{t("profile")}</p>
                            </Link>
                        </div>
                        <div className="text-center text-white mb-3">
                            <h3>{t('aboutGeo')}</h3>
                            <small>{t('about')}</small>
                        </div>
                        <div className="d-flex align-items-center justify-content-between w-100">

                            <Link to="/dashboard" className="sidebar-header">
                                <img src={settings} alt="" className="img-fluid logo-icon me-2" />
                            </Link>
                            <LanguageSwitcher onLanguageChange={handleLanguageChange}></LanguageSwitcher>
                            <Link to="/" className="sidebar-header">
                                <img src={logoutIcon} alt="" className="img-fluid logo-icon me-2" onClick={logout} />
                            </Link>
                        </div>
                    </Offcanvas.Body>
                </Navbar.Offcanvas>

            </Container>
        </Navbar>
    )
};

export default CustomNavbar