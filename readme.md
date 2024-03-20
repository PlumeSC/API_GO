# Football League Management System

## Overview
This Football League Management System is a comprehensive backend service developed in Go (Golang), incorporating RESTful APIs to manage football leagues, seasons, teams, players, and live match updates efficiently. Utilizing bcrypt for secure password handling, JWT for authentication, and external FOOTBALL API for real-time data, this project offers a robust platform for football league management.

## Features
- **Secure User Authentication**: Implements bcrypt for password hashing and JWT for maintaining secure sessions.
- **Dynamic Season and League Creation**: Checks for the existence of seasons and leagues in the database; if not present, it fetches necessary data from the FOOTBALL API to create them, including standings tables.
- **Player Management**: Retrieves player data from the FOOTBALL API, ensuring that only new players are added to the database with their respective statistics.
- **Match Management**: Facilitates the creation of matches for new seasons and allows for live score updates which in turn automatically update standings.
- **Standings and Matches Retrieval**: Provides endpoints to fetch current standings and match details, which can be filtered by team name or round.
- **Live Score Updates**: Supports live score tracking and automatic standings updates post-match conclusion.

## Technologies Used
- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT, bcrypt
- **External APIs**: FOOTBALL API for fetching real-time league, team, and player data

## Getting Started

### Prerequisites
- Go (Version 1.15 or later)
- PostgreSQL database

### Installation
1. Clone the repository:
