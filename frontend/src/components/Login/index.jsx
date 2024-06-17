import useLogin from './../../hooks/useLogin.js'
import user_icon from './../../img/icons/person.png'
import email_icon from './../../img/icons/email.png'
import password_icon from './../../img/icons/password.png'
import address_icon from './../../img/icons/address.svg'
import information_icon from './../../img/icons/information.svg'
import './index.css';
import logo from './../../img/logo/geo_report_white_text.svg'
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min";
import { Container, Row, Col, Modal, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import LanguageSwitcher from '../LanguageSwitcher/index.jsx'; 



const Login = ({ setToken }) => {
    const { email, username, setUsername, handleSubmit, isValidEmail, isValidPassword, password, action, setAction, setFirstName, setLastName, firstName, lastName, handleEmailChange, handlePasswordChange,
        handleShowPasswordModal, showPasswordModal, handleClosePasswordModal, showWrongPasswordModal, setShowWrongPasswordModal } = useLogin(
            setToken
        );
    const { t } = useTranslation();


    return (
        <Container fluid className="bg-black">
            <LanguageSwitcher />
            <Row className="vh-100">
                <Col sm={12} md={6} className="d-flex align-items-center justify-content-center">
                    <img src={logo} alt="Your Logo" className="img-fluid" />
                </Col>
                <Col sm={12} md={6} className="d-flex align-items-center justify-content-center flex-column">
                    <div className="card login-signup-forgot" style={{
                        display: 'flex', justifyContent: 'center', maxWidth: '650px', minWidth: '400px',
                    }}>
                        <div className="header">
                            <div className="text">{t(action)}</div>
                            <div className="underline"></div>
                        </div>
                        <form>
                            <div className="inputs">
                                {action === "Login" ? <div></div> :
                                    <><div className="input">
                                        <img className="svg" src={address_icon} alt="Last Name Icon" />
                                        <input type="text" className="form-control" value={firstName} onChange={(e) => setFirstName(e.target.value)} required />
                                        <label className="placeholderx">{t('firstName')}</label>
                                    </div>
                                        <div className="input">
                                            <img className="svg" src={address_icon} alt="Last Name Icon" />
                                            <input type="text" className="form-control" value={lastName} onChange={(e) => setLastName(e.target.value)} required />
                                            <label className="placeholderx">{t('lastName')}</label>
                                        </div>
                                        <div className="input">
                                            <img src={email_icon} alt="E-mail Icon" />
                                            <input
                                                type="email"
                                                className={`form-control ${!isValidEmail ? "invalid-input" : ""}`}
                                                value={email}
                                                onChange={handleEmailChange}
                                                required
                                            />
                                            <label className="placeholderx">{t('email')}</label>
                                        </div></>}
                                <div className="input">
                                    <img src={user_icon} alt="User Icon" />
                                    <input type="text" className="form-control" value={username} onChange={(e) => setUsername(e.target.value)} required />
                                    <label className="placeholderx">{t('username')}</label>
                                </div>

                                <div className="input">
                                    {!isValidPassword && action === "Sign Up" && (
                                        <img src={information_icon} alt="Information Icon" onClick={handleShowPasswordModal} style={{ cursor: 'pointer' }} />
                                    )}
                                    <img src={password_icon} alt="Password Icon" />
                                    <input
                                        type="password"
                                        className={`form-control ${!isValidPassword && action === "Sign Up" ? "invalid-input" : ""}`}
                                        value={password}
                                        onChange={handlePasswordChange}
                                        required
                                    />
                                    <label className="placeholderx">{t('password')}</label>
                                </div>
                            </div>
                        </form>
                        {action === "Sign Up" ? <div></div> :
                            <div className="forgot-password"> {t('forgotPassword')} <span> {t('clickHere')} </span></div>}
                        <div className="submit-container">
                            <div className={action === "Login" ? "submit gray" : "submit"} onClick={action === "Sign Up" ? handleSubmit : () => { setAction("Sign Up") }}>{t('signUp')}</div>
                            <div className={action === "Sign Up" ? "submit gray" : "submit"} onClick={action === "Login" ? handleSubmit : () => { setAction("Login") }} >{t('login')}</div>
                        </div>
                    </div>
                </Col>
            </Row>
            <Modal show={showPasswordModal} onHide={handleClosePasswordModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Password Requirements</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <p>{t('passwordRequirementsBody')}</p>
                    <ul>
                        <li>{t('lowercaseLetter')}</li>
                        <li>{t('uppercaseLetter')}</li>
                        <li>{t('number')}</li>
                        <li>{t('specialCharacter')}</li>
                    </ul>
                </Modal.Body>
            </Modal>
            <Modal show={showWrongPasswordModal} onHide={() => setShowWrongPasswordModal(false)}>
                <Modal.Header closeButton>
                    <Modal.Title>{t('incorrectPassword')}</Modal.Title>
                </Modal.Header>
                <Modal.Body className="text-center">
                    <p className="text-danger">{t('incorrectPasswordMessage')}</p>
                    <p>{t('tryAgain')}</p>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={() => setShowWrongPasswordModal(false)}>
                        Close
                    </Button>
                </Modal.Footer>
            </Modal>
        </Container>
    );

};

export default Login;