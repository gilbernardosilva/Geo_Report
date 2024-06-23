import React from "react";
import { Modal, Button } from "react-bootstrap";

function IssueDetailsModal({ issue, show, handleClose }) {
    return (
      <Modal show={show} onHide={handleClose}>
        {issue && ( // Only render if issue is not null or undefined
          <>
            <Modal.Header closeButton>
              <Modal.Title>{issue.title}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
              <p>
                <b>Description:</b> {issue.description}
              </p>
              <p>
                <b>Status:</b> {issue.status}
              </p>
              {/* ... other details (e.g., location, date, images) */}
            </Modal.Body>
            <Modal.Footer>
              <Button variant="secondary" onClick={handleClose}>
                Close
              </Button>
            </Modal.Footer>
          </>
        )}
      </Modal>
    );
  }

export default IssueDetailsModal;
