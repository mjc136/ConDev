<h1>Wa-Tor Simulation</h1>
<p>
  This project simulates a simplified Wa-Tor world ecosystem using Go and the Ebiten game engine. The simulation showcases interactions between fish and sharks on a toroidal grid.
</p>

<hr>

<h2>Features</h2>
<ul>
  <li><strong>Dynamic Ecosystem</strong>: Sharks hunt fish, fish reproduce, and entities move randomly.</li>
  <li><strong>Customisable Parameters</strong>: Adjust shark starvation, breeding cycles, and grid size.</li>
</ul>

<hr>

<h2>Simulation Parameters</h2>
<table>
  <thead>
    <tr>
      <th>Parameter</th>
      <th>Description</th>
      <th>Default Value</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>xdim</code></td>
      <td>Width of the simulation grid.</td>
      <td>100</td>
    </tr>
    <tr>
      <td><code>ydim</code></td>
      <td>Height of the simulation grid.</td>
      <td>100</td>
    </tr>
    <tr>
      <td><code>NumFish</code></td>
      <td>Initial number of fish.</td>
      <td>100</td>
    </tr>
    <tr>
      <td><code>NumShark</code></td>
      <td>Initial number of sharks.</td>
      <td>15</td>
    </tr>
    <tr>
      <td><code>fishBreed</code></td>
      <td>Steps required for fish to reproduce.</td>
      <td>50</td>
    </tr>
    <tr>
      <td><code>sharkBreed</code></td>
      <td>Steps required for sharks to reproduce.</td>
      <td>100</td>
    </tr>
    <tr>
      <td><code>sharkStarve</code></td>
      <td>Steps before a shark starves without eating.</td>
      <td>75</td>
    </tr>
    <tr>
      <td><code>foodEnergy</code></td>
      <td>Energy gained by a shark after eating a fish.</td>
      <td>100</td>
    </tr>
  </tbody>
</table>

<hr>

<h2>How to Run</h2>
<ol>
  <li>Clone this repository:
    <pre><code>git clone https://github.com/username/Wa-Tor.git
cd Wa-Tor</code></pre>
  </li>
  <li>Install dependencies:
    <pre><code>go get github.com/hajimehoshi/ebiten/v2</code></pre>
  </li>
  <li>Run the program:
    <pre><code>go run Wa-Tor.go</code></pre>
  </li>
</ol>


<h2>How to Set Up a Virtual Environment and Install Dependencies</h2>
<ul>
  <li><strong>Step 1: Create a Virtual Environment</strong>
    <ul>
      <li>On Windows:
        <pre><code>python -m venv venv</code></pre>
      </li>
      <li>On macOS and Linux:
        <pre><code>python3 -m venv venv</code></pre>
      </li>
    </ul>
  </li>

  <li><strong>Step 2: Activate the Virtual Environment</strong>
    <ul>
      <li>On Windows:
        <pre><code>venv\Scripts\activate</code></pre>
      </li>
      <li>On macOS and Linux:
        <pre><code>source venv/bin/activate</code></pre>
      </li>
    </ul>
  </li>

  <li><strong>Step 3: Install Dependencies from <code>requirements.txt</code></strong>
    <ul>
      <li>Run the following command to install all dependencies:
        <pre><code>pip install -r requirements.txt</code></pre>
      </li>
    </ul>
  </li>
</ul>

