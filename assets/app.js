document.addEventListener('DOMContentLoaded', () => {
    const ansi_up = new AnsiUp();
    const themeKey = "wt_theme";
    const linesKey = "wt_max_lines";
    let logLinesMax = parseInt(localStorage.getItem(linesKey), 10) || 500;
    let autoScroll = true;
    let logLinesCount = 0;
    const ol = document.getElementById('list');
    const pauseBtn = document.getElementById('pauseBtn');
    let isPaused = false;
    let conn = null;

    const changeTheme = (t) => {
        const theme = t.value;
        document.body.className = `theme-${theme}`;
        localStorage.setItem(themeKey, theme);
    };

    const clearScreen = () => {
        ol.innerHTML = "";
        logLinesCount = 0;
        appendLogMessage("Clear screen");
    };

    const appendLogMessage = (msg, kind = "log") => {
        if (kind === "log" && isPaused) return;
        const li = document.createElement('li');
        li.className = kind;
        li.innerHTML = ansi_up.ansi_to_html(msg);
        ol.appendChild(li);
        logLinesCount++;
        clearExcessLogs();
        scrollToBottom();
    };

    const clearExcessLogs = () => {
        while (logLinesCount > logLinesMax && autoScroll) {
            ol.removeChild(ol.firstChild);
            logLinesCount--;
        }
    };

    const scrollToBottom = () => {
        if (autoScroll) {
            window.scrollTo(0, document.body.scrollHeight);
        }
    };

    const pause = () => {
        isPaused = !isPaused;
        pauseBtn.textContent = isPaused ? "continue (Space)" : "pause (Space)";
        appendLogMessage(isPaused ? "Paused log print" : "Continued log print", "sys");
    };

    const toggleFullScreen = () => {
        const elem = document.documentElement;
        if (!document.fullscreenElement) {
            elem.requestFullscreen();
        } else {
            document.exitFullscreen();
        }
    };

    window.addEventListener('scroll', () => {
        autoScroll = window.innerHeight + window.scrollY >= document.body.scrollHeight;
    });

    const changeMaxLines = (t) => {
        logLinesMax = parseInt(t.value, 10);
        localStorage.setItem(linesKey, logLinesMax);
        clearScreen();
    };

    const openSocket = () => {
        conn = new WebSocket(`ws://${window.location.hostname}:${window.location.port}/ws${window.location.search}`);

        conn.onopen = () => conn.send("test");
        conn.onclose = () => appendLogMessage("Connection closed", "sys");
        conn.onerror = () => appendLogMessage("Socket error", "sys");
        conn.onmessage = (evt) => appendLogMessage(JSON.parse(evt.data).msg);

        window.onbeforeunload = () => conn.close();
    };

    openSocket();

    document.onkeydown = (e) => {
        if (e.key === 'Escape') closeConn();
        if (e.key === ' ') pause();
        if (e.key === 'l' && e.ctrlKey) clearScreen();
    };
});
