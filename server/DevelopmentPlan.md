# Development Plan for the Backend


## Architecture

There should be two different backend systems and a database:
- Minecraft Server Manager (MCSM)
  - The backend service that creates the different MC servers, sets them up and assigns them UUIDs, saves configs, starts/stops servers, and sends operator commands
  - Each compute device (probably EC2 instances) will have an instance of an MCSM

- Web Server
  - The web server that connects to the appropriate MC server manager and creates and manages the user authentication and accounts, provides the frontend apps with a more abstract API, and acts like a middleman in between the frontend apps and the MC server manager

- Database
  - User accounts:
    - Simply holds the user's email or phone number. JWT Tokens will be used to 

Both servers will be written in Go + the Gin framework.

## MC Server Manager (MCSM)

### Functions

- Create Server:
  - Makes the container containing a new MC server given the game version id and a config file
  - Creates a new world or copies over a provided one (uses the
- Delete Server:
- Start/Stop Server:
- Update World:
- Update Config:

### Example Workflows

User Creates a Server:
1. User makes a request to create a new MC server to the Web Server
2. The Web Server finds  an EC2 instance with an available slot
3. Server forwards that request to the MCSM of that EC2 instance
4. MCSM creates the new server and returns the UUID for that server
5. Web server updates the database with that UUID and EC2 instance identifier, sets up the SRV record, and shows the user a success message

### Considerations

Subdomain Routing:
- Each player can set a custom subdomain for each server they create, so SRV records can be used to allow a player to automatically join the right server that's hosted

Networking:
- Servers can be on different EC2 instances, there might be a lookup table on the database, and a table for storing the remaining slots of each EC2 instance

Optimization and Cost Management:
- These are things I may have to worry about when I get more users:
  - Moving around servers based on average activity to different EC2 instances so that I don't have multiple instances running suboptimal numbers of game servers

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