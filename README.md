# goSpotify
[![GitHub license](https://img.shields.io/github/license/chaithanyaMarripati/goSpotify)](https://github.com/chaithanyaMarripati/goSpotify/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/chaithanyaMarripati/goSpotify)](https://github.com/chaithanyaMarripati/goSpotify/issues)
### goSpotify uses Oauth2 to authorize the users with their spotify account and returns the song/album currently being played

## rename the .env.example to .env and attach clientId and clientSecret after registering an app in spotify dev dashboard
# QuickStart
1. [Install Docker](https://docs.docker.com/engine/install/).
2. Run the build command
   ```
   docker build -t gospotify:latest .
   ```
3. Run the start command
   ```
   docker run -p 8080:8080 gospotify:latest
   ```
# Todo list 
these are the pending action items, list will be updated overtime
- [ ] Add state in the url parmas when the user is redirected, to protect against csrf attacks
- [ ] After successful authentication return a html page instead of text
- [ ] Add error page for errors during /authorize step 
- [ ] Add more test cases
- [ ] Add more functionality to the app
- [ ] Buy a domain and add it to the cloud run

# Deployment
CI/CD pipeline is configured in gcp, google cloud build processes the builds and deploys it to the cloud run as per cloudbuild.yaml file.
This app is currently deployed to cloud run
[app](https://gospotify-sjskww6rpa-el.a.run.app/)