const express = require('express');
const path = require('path');
const cors = require('cors');

const app = express();
const PORT = process.env.PORT || 8080;

app.use((req, res, next) => {
    res.setHeader("Access-Control-Allow-Origin", "*"); // или укажите конкретный домен
    res.setHeader("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS");
    res.setHeader("Access-Control-Allow-Headers", "Content-Type");
    next();
});

// Укажите папку с вашими статическими файлами (например, HTML, CSS, JS)
app.use(express.static(path.join(__dirname, 'public')));

// Отправка index.html при запросе на корень
app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'public', 'index.html'));
});

// Запуск сервера
app.listen(PORT, () => {
  console.log(`Сервер запущен на http://localhost:${PORT}`);
});