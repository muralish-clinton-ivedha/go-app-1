# Golang HTTP Server with Clean Architecture
This repository template provides a framework for generating a Golang HTTP server with a clean architecture. Clean architecture emphasizes separation of concerns and maintainability by organizing code into distinct layers, making it easier to understand, test, and modify.

## Deployment Instructions

To deploy this server, follow these steps:

1. Add GitHub Secret Key in Actions:
- Navigate to the repository's GitHub page.
- Go to the "Settings" tab.
- Select "Secrets" from the sidebar.
- Add a new secret named SERVICE_ACCOUNT_KEY.
- Set the value to the JSON key of your Google Cloud Platform (GCP) service account.
- For more details, refer to the documentation on advanced authentication for Google Container Registry.

2. Add the secret in the github workflow file
- Navigate to .github/workflows/deployment.yml
- Locate the key service_account_key
- Update the file to add the value as 

3. Trigger the workflow manually
