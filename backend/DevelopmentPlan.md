# Development Plan for the Backend

## TODO + Dev Roadmap:

### Phase 1: Backend, Server Logic, Server Hosting Framework
- [ ] Define the web server API
- [x] Get PostgreSQL running locally according to the DB design doc
  - [x] DB Design Doc: https://dbdiagram.io/d/ChocolateHosting-69c1929efb2db18e3be642aa

- [x] Fill in the game_version table in the DB with a Python script
  - [x] Parse manifest.json
  - [x] Get server download links for each version
  - [x] Add them to DB

- [ ] Create a basic MCServerManager server for the EC2 instance
  - [ ] Basic functions:
    - [ ] Create server
      - [ ] Download server.jar with the URL in the DB
      - [ ] Agree to EULA
      - [ ] Somehow send commands to the opened console
      - [ ] Update servers table in DB

    - [ ] Stop server
      - [ ] Send simple stop command using

    - [ ] Delete server







## Architecture

There should be two different backend systems and a database:
- Minecraft Server Manager (MCSM)
  - The backend service that creates the different MC servers, sets them up and assigns them UUIDs, saves configs, starts/stops servers, and sends operator commands
  - Each compute device (probably EC2 instances) will have an instance of this MCSM

- App Server
  - The app server that connects to the appropriate MC server manager and creates and manages the user authentication and accounts, provides the frontend apps with a more abstract API, and acts like a middleman in between the frontend apps and the MC server manager

- Databaseo
  - User accounts:
    - Simply holds the user's email or phone number. JWT Tokens will be used for authentication
  - Compute device table:
    - Stores references to the different EC2 instances, their ip addresses, and useful metadata
    - Fields/columns:
      - DeviceID : UUID
      - MaxServerCount : Int
      - RunningServerCount : Int
      - PerformanceTier : Int
  - Server list:
    - All the different Minecraft servers
    - Fields/columns:
      - ServerID : UUID
      - OwnerID : UUID
      - ComputeDeviceID: UUID
      - SubscriptionTier: 
  - Subscriptions (TBD):
    - The different tiers available for each server

Both servers will be written in Go + the Gin framework.

## MC Server Manager (MCSM)

### Functions

- Create Server:
  - Makes the container containing a new MC server given the game version id and a config file
  - Creates a new world or copies over a provided one uses the
- Delete Server:
- Start/Stop Server:
- Update World:
- Update Config:

### Considerations

Subdomain Routing:
- Each player can set a custom subdomain for each server they create, so SRV records can be used to allow a player to automatically join the right server that's hosted

Networking:
- Servers can be on different EC2 instances, there might be a lookup table on the database, and a table for storing the remaining slots of each EC2 instance

Optimization and Cost Management:
- These are things I may have to worry about when I get more users:
  - Moving around servers based on average activity to different EC2 instances so that I don't have multiple instances running suboptimal numbers of game servers

## App Server

### API Functions

Servers:
- POST : CreateServer(userID, ServerConfig) : Server
- DELETE : DeleteServer(serverID) :

- PUT : StartServer(serverID) : ServerStatus
- PUT : StopServer(serverID) : ServerStatus
- PUT : EditServer(serverID, ServerInfo) :
- PUT : EditServerConfig(serverID, ServerConfig) :

Accounts/User Auth:
- POST : Signup(user email | phone number, OTP) : AuthToken
- GET : Login(user email | phone number, OTP) : AuthToken
- DELETE : Logout(user email | phone number) :
- PUT : ChangeEmail(user email) : AuthToken
- PUT : ChangePhone(user phone number) : AuthToken

- GET : GetServers(userID) : Server[]
- 


## Core feature list

- Bedrock + Java servers available
- Server Options (server.properties)
- Custom subdomain
- Limits:
  - Player Count (Higher player counts reserved for different subscription tiers)
  - Viewing Distance (Limited to make sure server performance is ideal and predictable)

### Optimizations and Defaults

- Default configs that makes sense
  - Disable spawn chunk protection
  - Presets for survival with friends and a creative building world
  - Enable rcon for server info on the app dashboard
- Paper servers where available