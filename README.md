`ip-malicious-db` is a Go service designed to load and manage malicious IPs and IP ranges by country. The service retrieves IP data from GitHub, specifically the [FireHOL IP blocklists](https://github.com/firehol/blocklist-ipsets), and stores it in a Neo4j database as nodes with the fields `id` (IP address) and `country` (ISO 3166-1 alpha-2 country code).

## 🚀 Features

- 🗂️ Fetches and stores malicious IPs and IP ranges categorized by country.
- 🌐 Retrieves data from the [FireHOL blocklist IPsets](https://github.com/firehol/blocklist-ipsets).
- 🛢️ Persists the data in a **Neo4j database** for advanced querying and integration.
- 📦 Provides a REST endpoint to load data into the database.



## 🛠️ Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/ip-malicious-db.git
   cd ip-malicious-db
   ```

2. **Set up Neo4j**:
   - Install Neo4j: [Neo4j Installation Guide](https://neo4j.com/docs/operations-manual/current/installation/)
   - Start the Neo4j database:
     ```bash
     neo4j start
     ```
   - Configure the database credentials in your environment:
     ```bash
     export NEO4J_URI="bolt://localhost:7687"
     export NEO4J_USER="neo4j"
     export NEO4J_PASSWORD="your_password"
     ```

3. **Build the service**:
   ```bash
   go build -o ip-malicious-db ./cmd
   ```

4. **Run the service**:
   ```bash
   ./ip-malicious-db
   ```



## 🔧 Usage

The service exposes a REST endpoint to fetch and load malicious IPs into the database.

### Endpoint: `/save-malicious-ip`

#### Method: `POST`

#### Description:
Fetches malicious IP data for all countries from GitHub and stores it in Neo4j.

#### Example Request:
```bash
curl -X POST http://localhost:8080/save-malicious-ip
```

### Data Storage in Neo4j

- **Nodes**:
  - `IP` nodes:
    - **Fields**:
      - `id`: Represents the IP address or IP range (primary key).
      - `country`: Represents the ISO 3166-1 alpha-2 code of the country.

There are no relationships between the nodes.


## 📚 Example Cypher Queries

### List All Malicious IPs
```cypher
MATCH (ip:IP)
RETURN ip.id
```

### List All Malicious IPs for a Specific Country
```cypher
MATCH (ip:IP {country: "us"})
RETURN ip.id
```



## 🌍 How It Works

1. **Fetch Data**: 
   - Downloads IP blocklist data for each country from the [FireHOL blocklist IPsets](https://github.com/firehol/blocklist-ipsets).
2. **Store in Neo4j**:
   - Creates `IP` nodes for each malicious IP or range, with the fields:
     - `id`: IP address.
     - `country`: Country code.



## 📚 Future Features

- 🌐 Additional REST API endpoints for querying malicious IPs.
- 📊 Analytics and visualization for malicious IP trends.
- 🛡️ Integration with real-time threat detection tools.



## 🤝 Contributions

Contributions are welcome! Please fork the repository, create a feature branch, and submit a pull request.



## 🛡️ License

This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for details.



## 🌟 Acknowledgments

Special thanks to the [FireHOL project](https://github.com/firehol/blocklist-ipsets) for providing the data that powers this service and to the Neo4j community for their database technology.
