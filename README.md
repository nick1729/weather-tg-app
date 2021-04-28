# weather-tg-app
Simple telegram weather application.
<p><code>git clone https://github.com/nick1729/weather-tg-app.git</code></p>

<br>Create new bot on telegram using <a href="https://telegram.me/BotFather" rel="nofollow">@BotFather</a> and copy your token to <code>/config/config.json</code> configuration file.</br>
<br>Register on the site <a href="https://openweathermap.org/" rel="nofollow">openweathermap.org</a> openweathermap.org (free) and copy your api key to configuration file too.</br>

<br>Build the project <code>go build .</code> or only run <code>go run .</code></br>

<br>Available commands:</br>
<p><code>/weather</code> - show current weather</p>
<p><code>/city [city name]</code> - set name of the city</p>
<p><code>/coordinates [37.62, 55.75]</code> - set city by nearest coordinates</p>
<p><code>/units [metric, imperial]</code> - set measurement units</p>
<p><code>/lang [ar, cz, de, en, fr, it, ja, kr, nl, pt, ru, sp, tr, ua]</code> - set language</p>

<br>Also, the list of commands is available via <code>/help</code></br>
