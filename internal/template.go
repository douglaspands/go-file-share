package internal

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <title>Go File Share</title>
    <style>
        /* Reset básico para evitar quebras de layout */
        * {
            box-sizing: border-box;
        }

        :root {
            --primary: #2563eb;
            --bg: #f8fafc;
            --container-bg: #ffffff;
            --text: #1e293b;
            --text-muted: #64748b;
            --border: #f1f5f9;
            --hover: #eff6ff;
            --breadcrumb-bg: #f1f5f9;
            --input-bg: rgba(255, 255, 255, 0.2);
            --icon-folder: #f59e0b;
            --icon-file: #94a3b8;
            --icon-back: #3b82f6;
        }

        body.dark-mode {
            --bg: #0f172a;
            --container-bg: #1e293b;
            --text: #f1f5f9;
            --text-muted: #94a3b8;
            --border: #334155;
            --hover: #334155;
            --breadcrumb-bg: #0f172a;
            --input-bg: rgba(0, 0, 0, 0.2);
            --icon-folder: #fbbf24;
            --icon-file: #cbd5e1;
        }

        body {
            font-family: -apple-system, system-ui, sans-serif;
            background: var(--bg);
            color: var(--text);
            margin: 0;
            padding: 15px;
            transition: background 0.3s, color 0.3s;
        }

        .container {
            max-width: 900px;
            margin: 0 auto;
            background: var(--container-bg);
            border-radius: 12px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            overflow: hidden;
            display: flex;
            flex-direction: column;
        }

        .header {
            padding: 20px;
            background: var(--primary);
            color: white;
        }

        .header-top {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 15px;
        }

        .header h2 { margin: 0; font-size: 1.2rem; }

        #search-input {
            width: 100%;
            padding: 12px;
            border: none;
            border-radius: 8px;
            background: var(--input-bg);
            color: white;
            font-size: 1rem;
            outline: none;
        }

        #search-input::placeholder { color: rgba(255, 255, 255, 0.7); }

        .theme-toggle {
            background: rgba(255, 255, 255, 0.2);
            border: none;
            color: white;
            padding: 8px 12px;
            border-radius: 8px;
            cursor: pointer;
            font-size: 0.8rem;
        }

        .breadcrumb {
            padding: 10px 20px;
            background: var(--breadcrumb-bg);
            font-size: 0.9rem;
            color: var(--text-muted);
            border-bottom: 1px solid var(--border);
        }

        /* AJUSTE NA SEÇÃO DE UPLOAD */
        .upload-section {
            padding: 20px;
            background: var(--container-bg);
            border-bottom: 2px dashed var(--border);
            width: 100%;
        }

        .upload-form {
            display: flex;
            flex-direction: column; /* Mobile first: um embaixo do outro */
            gap: 12px;
        }

        input[type="file"] {
            width: 100%;
            padding: 10px;
            background: var(--bg);
            border: 1px solid var(--border);
            border-radius: 8px;
            color: var(--text);
            font-size: 0.85rem;
        }

        .upload-btn {
            background: var(--primary);
            color: white;
            border: none;
            padding: 12px;
            border-radius: 8px;
            cursor: pointer;
            font-weight: bold;
            font-size: 0.9rem;
            width: 100%; /* Ocupa tudo no mobile */
            transition: background 0.2s;
        }

        .upload-btn:hover {
            filter: brightness(1.1);
        }

        /* Responsividade para Telas Maiores */
        @media (min-width: 600px) {
            .upload-form {
                flex-direction: row; /* Volta para o lado a lado */
                align-items: center;
            }
            .upload-btn {
                width: auto;
                min-width: 140px;
                padding: 10px 20px;
            }
        }

        /* Lista de Arquivos */
        .list-item {
            display: flex;
            align-items: center;
            padding: 15px 20px;
            text-decoration: none;
            color: inherit;
            border-bottom: 1px solid var(--border);
            transition: 0.2s;
        }

        .list-item:last-child { border-bottom: none; }
        .list-item.hidden { display: none; }
        .list-item:active { background: var(--hover); }

        @media (min-width: 768px) {
            .list-item:hover { background: var(--hover); }
        }

        .icon-svg {
            width: 28px;
            height: 28px;
            margin-right: 15px;
            flex-shrink: 0;
        }

        .icon-folder { fill: var(--icon-folder); }
        .icon-file { fill: var(--icon-file); }
        .icon-back-arrow { fill: none; stroke: var(--icon-back); stroke-width: 2.5; stroke-linecap: round; stroke-linejoin: round; }

        .info { flex-grow: 1; min-width: 0; }
        .name { font-weight: 600; word-break: break-all; overflow-wrap: break-word; }
        .details { font-size: 0.8rem; color: var(--text-muted); }

        .progress-container {
            width: 100%;
            background-color: var(--bg);
            border-radius: 8px;
            margin-top: 15px;
            display: none;
            overflow: hidden;
            border: 1px solid var(--border);
        }

        .progress-bar {
            width: 0%;
            height: 10px;
            background-color: var(--primary);
            transition: width 0.2s;
        }

        .progress-text { font-size: 0.75rem; text-align: center; margin-top: 6px; font-weight: bold; }
    </style>
</head>

<body>
    <svg style="display: none;">
        <symbol id="icon-filled-folder" viewBox="0 0 24 24">
            <path d="M20 4H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm0 14H4V8h16v10zm-2-12H6V6h12v2z" />
            <path d="M2 6C2 4.9 2.9 4 4 4h6l2 2h8c1.1 0 2 .9 2 2v10c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6z" />
        </symbol>
        <symbol id="icon-filled-file" viewBox="0 0 24 24">
            <path d="M14 2H6c-1.1 0-1.99.9-1.99 2L4 20c0 1.1.89 2 1.99 2H18c1.1 0 2-.9 2-2V8l-6-6zm2 16H8v-2h8v2zm0-4H8v-2h8v2zm-3-5V3.5L18.5 9H13z" />
        </symbol>
        <symbol id="icon-back-thick" viewBox="0 0 24 24">
            <polyline points="9 14 4 9 9 4"></polyline>
            <path d="M20 20v-7a4 4 0 0 0-4-4H4"></path>
        </symbol>
    </svg>

    <div class="container">
        <div class="header">
            <div class="header-top">
                <h2>Go File Server</h2>
                <button class="theme-toggle" onclick="toggleTheme()" id="theme-btn">🌙 Dark</button>
            </div>
            <div class="search-container">
                <input type="text" id="search-input" placeholder="Search files..." onkeyup="filterFiles()">
            </div>
        </div>

        <div class="breadcrumb">📍 Current path: {{.CurrentPath}}</div>

        <div class="upload-section">
            <form id="upload-form" class="upload-form">
                <input type="file" id="file-input" name="file" required>
                <button type="submit" class="upload-btn">📤 Upload</button>
            </form>
            <div id="progress-container" class="progress-container">
                <div id="progress-bar" class="progress-bar"></div>
            </div>
            <div id="progress-text" class="progress-text"></div>
        </div>

        <div id="file-list">
            {{range .Files}}
            <a href="{{.Path}}" class="list-item" data-name="{{.Name}}">
                {{if eq .Name ".. (Back)"}}
                <svg class="icon-svg icon-back-arrow"><use href="#icon-back-thick" /></svg>
                {{else if .IsDir}}
                <svg class="icon-svg icon-folder"><use href="#icon-filled-folder" /></svg>
                {{else}}
                <svg class="icon-svg icon-file"><use href="#icon-filled-file" /></svg>
                {{end}}

                <div class="info">
                    <div class="name">{{.Name}}</div>
                    {{if not .IsDir}}<div class="details">{{.Size}}</div>{{end}}
                </div>
            </a>
            {{end}}
        </div>
    </div>

    <script>
        const body = document.body; 
        const btn = document.getElementById('theme-btn'); 
        const savedTheme = localStorage.getItem('theme');
        
        if (savedTheme === 'dark' || (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)) { 
            body.classList.add('dark-mode'); btn.innerText = '☀️ Light'; 
        }

        function toggleTheme() { 
            body.classList.toggle('dark-mode'); 
            const isDark = body.classList.contains('dark-mode'); 
            localStorage.setItem('theme', isDark ? 'dark' : 'light'); 
            btn.innerText = isDark ? '☀️ Light' : '🌙 Dark'; 
        }

        function filterFiles() { 
            const query = document.getElementById('search-input').value.toLowerCase(); 
            const items = document.querySelectorAll('.list-item'); 
            items.forEach(item => { 
                const name = item.getAttribute('data-name').toLowerCase(); 
                if (name.includes('.. (back)')) { item.classList.remove('hidden'); return; } 
                if (name.includes(query)) { item.classList.remove('hidden'); } else { item.classList.add('hidden'); } 
            }); 
        }

        const uploadForm = document.getElementById('upload-form'); 
        const fileInput = document.getElementById('file-input'); 
        const progressBar = document.getElementById('progress-bar'); 
        const progressContainer = document.getElementById('progress-container'); 
        const progressText = document.getElementById('progress-text');

        uploadForm.addEventListener('submit', e => {
            e.preventDefault(); 
            const file = fileInput.files[0]; 
            if (!file) return; 
            
            const formData = new FormData(); 
            formData.append('file', file); 
            const xhr = new XMLHttpRequest();
            
            xhr.upload.addEventListener('progress', e => { 
                if (e.lengthComputable) { 
                    const percent = (e.loaded / e.total) * 100; 
                    progressContainer.style.display = 'block'; 
                    progressBar.style.width = percent.toFixed(0) + '%'; 
                    progressText.innerText = ` + "`" + `Uploading: ${percent.toFixed(0)}%` + "`" + `; 
                } 
            });
            
            xhr.onload = () => { 
                if (xhr.status >= 200 && xhr.status < 300 || xhr.status === 303) { 
                    progressText.innerText = "✅ Upload Complete! Reloading..."; 
                    setTimeout(() => location.reload(), 1000); 
                } else { 
                    alert("Error during upload."); 
                    progressContainer.style.display = 'none'; 
                } 
            };
            xhr.open('POST', window.location.pathname, true); 
            xhr.send(formData);
        });
    </script>
</body>
</html>
`
