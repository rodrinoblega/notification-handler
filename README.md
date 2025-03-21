# Notification system

![technology Golang](https://img.shields.io/badge/technology-Golang-blue.svg)

This system is capable of managing millions of users, ensuring efficient notification delivery and allowing for flexibility in staff management.

## Key points

- Reliability
  - Ensure notification delivery without message loss.
  - Retry handling in case of failures.
  - Logs and real-time monitoring.

- In-App Notification Center
  - Notification history per user.
  - Pagination and descending order by date.
  - Filters by notification type and date range.

- Flexible Template Management
  - Modify templates without depending on the technical team.
  - Support for dynamic placeholders.

## Running the app

You can run the app locally using Docker by executing the following command:

```docker compose up --build ```

This command will:
- Start a local PostgreSQL instance on port 5432.
- Make all the migration process using the files in /migrations.
- Start a kafka service instance on port 9092
- Create a topic called notifications
- Start an application that simulates the sending of an event.
- Start an application that acts as a consumer of the event and process a notification. It will also make endpoints available for history query and template management.
- Start a cron job that will run every 30 seconds looking for failed tasks to reprocess.

## More information

For more information you can visit the [tech documentation](https://docs.google.com/document/d/1pH9V_8-jkIA6V9bjg7Nci3Y3bozmmhhHf-2qrH3vbV0/edit?usp=sharing)

## Questions

* [rnoblega@gmail.com](rnoblega@gmail.com)


