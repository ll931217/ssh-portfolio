# Projects

I will break down my projects into two categories, professional projects and hobby projects. Professional projects are projects that I have worked on in a professional setting, whereas hobby projects are projects that I have worked on in my free time.

## Professional Projects

Professional projects will each be categorised by the company I worked for at the time of development. I will include a brief description of the project, the technologies used, my role in the project, what contributions I have made for the project and what I have learned developing the project.

### [Gyro Systems Inc.](https://www.gyro.com.tw) (Tech Lead)

- Erack Controller and Server: This project was for our first semi-conductor factory client. The project consists of a server that will keep track of all the erack's in the factory. An erack (Electronic Rack) is a shelf to store FOUP's in the factory. The data of the eracks received will then be displayed on either the monitor of the erack or the user can connect to the server and retrieve the website.
  - Technologies used:
    - Node.js (Express.js)
    - Nuxt.js (Vue.js)
    - MariaDB
    - BootstrapVue
    - Tibco Rendesvous: an industry grade communication protocol with a publish-subscribe model. That emphasis on high performance, low latency and stability.
    - Java: This was used to create a middleware between the Tibco Rendesvous and the server since Tibco Rendesvous only provides Java SDK.
  - My role:
    - I was the sole developer in this project. My manager worked on the erack's software. I worked on
      - The interface displayed on the monitor on each erack
      - The server to connect all eracks in the factory
      - The middleware to connect the Tibco Rendesvous and the server
    - I was also the maintainer of the system until the system was stable enough to hand over to the client
    - I had to travel overseas to the client, which means directly interact with the client.
  - What I learned:
    - I learned how to work with a client directly
    - I learned how to work with a publish-subscribe communication protocol
- Material Control System (MCS): This was the main project and software product of Gyro.

  - What does this system do?
    - Monitor the materials in the factory
    - View the status of all machines connected to the system
    - View the location of the materials, such as if it is on an AMR (Autonomous Mobile Robot) or on an erack.
    - Includes dashboard to view the statistics of the factory
    - Controls the machines connected to the system
  - Technologies used:
    - Flask (Python + Jinja2)
    - jQuery
    - Bootstrap
    - MariaDB
  - My role:
    - I was the creator and main developer in this project. I worked on the front-end and back-end of the project.
  - What I learned:
    - How to improve the performance of queries in a database. Since we have a lot of data being stored in the database from the machines, I had to learn how to optimise the queries to make it faster. (Future me: I should have used a different database for logging these events such as InfluxDB instead of MariaDB, since InfluxDB is designed for time series data)
    - Loading a team of developers to help me with the project. I had to learn how to delegate tasks and manage the team to make sure the project is on track. This also includes member onboarding, code reviews and making sure the code is up to standard.

- ITRI's MCS: This project is very similar to the previous project however this was in collaboration with ITRI (Industrial Technology Research Institute). This project was to create a Material Control System for ITRI's AMRs but using Gyro's eRacks. For this project I had more control on what technologies would be used. I chose these technologies because I could easily spin up a server and have a front-end up and running quickly.

  - Technologies used:
    - Flask (Python)
    - Vuejs
    - BootstrapVue
    - MSSQL: This was chosen because ITRI's server
    - Docker: Package the system to ITRI
  - My role:
    - I was the main developer in this project. I worked on the front-end and back-end of the project.
    - I had to work with ITRI's engineers to complete with project.
  - What I learned:
    - How to work with a client that is a government organisation
    - How to work with a client that has a lot of requirements
    - How to work with a client that has a lot of changes in requirements

### Houzz (Software Engineer)

Every project I worked on at Houzz was to improve the GraphQL API. I was part of the API team and my main role was to improve the performance of the API. Using TypeScript as the main language, we also used PHP in some cases as the main REST API was written in PHP.

- GraphQL query cost analyzer, this works by providing the default cost of certain datatypes in the schema. This feature can be implemented with rate-limiting to prevent abuse of the system.
  - The costs are determined by the following factors:
    - Computed field
    - Backend service calls
    - Lists
- Improved Graphene's (Our in-house GraphQL server framework) image size from 99% to 51% which helped the company save on image storage costs.

During my time at Houzz, I truly learned a lot of skills such as how to write cleaner code, best development practices, great documentation and also improving my communication skills with my team and other teams.

### [LeapIn HR](https://leapin.co) (Back-end Engineer) (Current company)

I work in this company as a back-end engineer, I worked with our front-end engineer to develop the system we have today to revolutionise the HR industry.

- Technologies used:
  - Express.js (Node.js)
  - PostgreSQL
  - passport.js
  - redis cache
  - Sentry
    - Otel
    - Log management
  - Docker
  - AWS
    - CloudWatch
    - S3
    - RDS
    - ECS
    - AWS Copilot CLI
  - GitHub Actions
  - Newman running postman test collection

### Freelance Projects

- Project management: A freelance project I worked on to generate a QR Code that can be printed and pasted onto products in stores for customers to scan and get more detailed information about the product. At the same time, get advertisements that can direct customers to similar products.

  - Technologies used:
    - Nuxt.js (Vue.js)
    - i18n: For internationalisation
    - Vuetify
    - xlsx: To generate the QR Codes
    - Apollo (GraphQL)
    - MongoDB
    - JWT (JSON Web Tokens)

- [Uhuru](https://github.com/ll931217/uhuru): A website to keep record of all African businesses in South Africa. The project was discontinued after my client abandoned it.
  - Technologies used:
    - LAMP stack
    - PHPMailer

## Hobby Projects

- [Vue-TreeView](https://baoge.dev/vue-treeview): This was a side project that I developed to help work on the "Erack Controller and Server". The project was to create a tree view component that can be used in the project. I modified it further to make it more user-friendly and added more features to it.
  - Technologies used:
    - Vue.js
    - Font Awesome
  - What I learned:
    - Further improved my understanding of recursion
- [Python JSON Validator](https://github.com/ll931217/python-json-validator): This was a side project that I developed to help work on the "Material Control System". The project was to create a JSON validator that can be used to validate the JSON data that is being sent to the server.
  - Technologies used:
    - Python
  - What I learned:
    - How to validate JSON data
- [kancli](https://github.com/ll931217/kancli): A kanban board in the terminal. This was initially a tutorial for Charm's bubbletea TUI framework, but I improved it to include saving the board to a file.
- [Wedding Questionaire](https://github.com/ll931217/wedding-questionnaire): A little QA game made for my wedding. This system includes a questionaire that guests can answer and the system will keep track of who answered the questionaire and who has not. The system will also display the results of the questionaire to the guests.
  - Technologies used:
    - Next.js (React)
    - Tailwind
    - Docker
- [Wedding Invitation](https://baoge.dev/wedding-invitation): A simple instagram story-like digital wedding invitation, it helped us save on some wedding costs.
  - Technologies used:
    - React
    - Github Pages
- [Sunnyvale](https://sunnyvale.baoge.dev): A ordering and order management system for my wife's shop.
  - Technologies used:
    - Sveltekit
    - Line SDK
    - Coolify: Hosted site and pocketbase on there
    - AWS S3 for pocketbase backups and file storage
