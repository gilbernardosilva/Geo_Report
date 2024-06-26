# Geo Report API

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report](https://github.com/keisaki/Geo_Report_API/actions/workflows/go.yml/badge.svg)](https://github.com/keisaki/Geo_Report_API/actions/workflows/go.yml)  

Geo Report API is a backend system for managing and tracking location-based issue reports. It provides a RESTful API for users to report problems, attach images, and view the status of their reports.

## Features

- **User Registration and Authentication:** Secure user accounts with JWT-based authentication.
- **Role-Based Access Control (RBAC):** Different permissions for regular users and authorities.
- **Geolocation-Based Reporting:** Users can report issues with precise location data.
- **Image Uploads:**  Allows users to attach images to their reports.
- **Issue Tracking:**  Keeps track of the status of reported issues.
- **Authority Dashboard:** Provides a dashboard for authorities to manage and resolve reported issues.

## Installation and Setup

1. **Clone the repository:**

2. **Set up the .env variables:**

```bash
DB_HOST=
DB_PORT=
DB_DATABASE=
DB_USERNAME=
DB_PASSWORD=
JWT_PRIVATE_KEY=
TOKEN_TTL=3600  # Expiration time in seconds (e.g., 1 hour)
```
