

  <h1>💪 GoFit - Workout Planner App</h1>
  <p>
    GoFit is a full-stack web application that helps users plan, organize, and manage gym workout routines.
    Built by Team 9, this project features a RESTful backend using Golang and a responsive frontend using React, Vite, and TailwindCSS.
  </p>

  <h2>📸 Preview</h2>
  <img src="./frontend/src/assets/gofit.png" alt="App Preview" style="width:100%; max-width: 600px;" />

  <h2>🧩 Features</h2>

  <h3>✅ Backend</h3>
  <ul>
    <li>Create and manage workout flows and exercises</li>
    <li>PostgreSQL integration with GORM</li>
    <li>Password hashing with validation</li>
    <li>Configurable via environment variables</li>
    <li>Dockerized for easy deployment</li>
    <li>Centralized logging support via Logstash</li>
    <li>Integration and unit tests</li>
  </ul>

  <h3>🌐 Frontend</h3>
  <ul>
    <li>Modern UI built with TailwindCSS</li>
    <li>Axios-based service layer</li>
    <li>Form validation with Zod</li>
    <li>Dynamic routing with React Router</li>
    <li>Global typography and responsive design</li>
    <li>Deployable to Vercel with clean URLs</li>
  </ul>

  <h2>📂 Folder Structure</h2>
  <pre>
.
├── backend/        # Golang API (REST)
│   ├── cmd/        # App entrypoint
│   ├── internal/   # Config, DB, handlers, models
│   ├── tests/      # Integration tests
│   ├── utils/      # Utility functions (e.g. hash password)
│   ├── Dockerfile
│   ├── Makefile
│   └── logstash.conf
├── frontend/       # React app (Vite + TailwindCSS)
│   ├── src/
│   │   ├── assets/
│   │   ├── components/
│   │   ├── layouts/
│   │   ├── lib/       # API client (axios)
│   │   ├── pages/
│   │   ├── schemas/   # Zod validation schemas
│   │   └── services/  # API service functions
│   ├── vite.config.ts
│   ├── vercel.json
├── docker-compose.yml
└── README.md
  </pre>

  <h2>🛠️ Getting Started</h2>

  <h3>⚙️ Backend</h3>
  <ol>
    <li>
      <strong>Setup</strong>
      <pre><code>
cd backend
cp .env-example .env
make build
make run
      </code></pre>
    </li>
    <li>
      <strong>Run Tests</strong>
      <pre><code>
make tests
      </code></pre>
    </li>
    <li>
      <strong>Docker</strong>
      <pre><code>
docker-compose up --build
      </code></pre>
    </li>
  </ol>

  <h3>🎨 Frontend</h3>
  <ol>
    <li>
      <strong>Setup</strong>
      <pre><code>
cd frontend
npm install
npm run dev
      </code></pre>
    </li>
    <li>
      <strong>Environment Variable</strong>
      <pre><code>
VITE_API_URL=http://localhost:8080
      </code></pre>
    </li>
  </ol>

  <h2>🔒 Security</h2>
  <ul>
    <li>Passwords are securely hashed with validation logic.</li>
    <li>Sensitive configs are stored via <code>.env</code> and <code>.env-example</code>.</li>
  </ul>

  <h2>🐳 Deployment</h2>
  <h3>Backend:</h3>
  <ul>
    <li>Uses <code>Dockerfile</code> and <code>docker-compose.yml</code></li>
    <li>Connects to PostgreSQL</li>
    <li>Logstash configuration included (<code>logstash.conf</code>)</li>
  </ul>

  <h2>👥 Contributing</h2>
  <p>Pull requests are welcome! For major changes, open an issue first to discuss what you’d like to change.</p>

  <h2>📄 License</h2>
  <p>MIT License</p>

  <h2>💪 Team</h2>
  <p>Made with effort by <a href="https://github.com/ProgramadoresSemPatria/Team-9" target="_blank">Team-9</a></p>

</body>
</html>



 <h2>Application Deploy</h2>
https://gofitpsp.vercel.app/
