## Database design

This package contains the models used by Gorm in the database. A side effect is
that this package essentially acts as the database design spec.

This document exists as way to both document the current database and collaborate
on future design changes.

### MVP database design

To start with, the database will need to have tables for the following:

* Users
* Suggestions
* PlatformIdentities, e.g. Discord identity, IRC identity, etc.

These three tables should give us the ability to launch an MVP on which we can
iterate.

The User table will need the following fields:

| Field Name | Description       | Reason for inclusion |
| ---------- | ----------------- | -------------------- |
| ID         | Primary Key/internal ID |                |

The User table will have a one-to-many relationship with the PlatformIdentities
table will need the following fields:

| Field Name | Description | Reason for inclusion |
| ---------- | ----------------- | -------------------- |
| Platform   | Platform name     |                      |
| Nick       | User's name on platform | |
| UserID     | Foreign key for User.ID

The Suggestions table will need the following fields:

| Field Name | Description | Reason for inclusion |
| ---------- | ----------------- | -------------------- |
| Suggester  | Foreign key User.ID |                      |
| Title      | Suggested title   | |
| Show       | Show ID set with !show start | |
