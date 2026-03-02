/**
 * Dark Mode Toggle
 */

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

/**
 * Upload File
 */

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
            progressText.innerText = `Uploading: ${percent.toFixed(0)}%`;
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
    xhr.open('POST', `/files${window.location.pathname}`, true);
    xhr.send(formData);
});