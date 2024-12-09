<h1>Wa-Tor Simulation</h1>
<p>
  This project simulates a simplified Wa-Tor world ecosystem using <strong>Go</strong> and the <strong>Ebiten</strong> game engine. The simulation showcases interactions between fish and sharks on a toroidal grid.
</p>

<hr>

<h2>Features</h2>
<ul>
  <li><strong>Dynamic Ecosystem</strong>: Sharks hunt fish, fish reproduce, and entities move randomly.</li>
  <li><strong>Customisable Parameters</strong>: Adjust shark starvation, breeding cycles, grid size, and the number of threads.</li>
  <li><strong>Threaded Execution</strong>: Supports single-threaded and multi-threaded execution for performance comparison.</li>
</ul>

<hr>

<h2>How to Run</h2>
<ol>
  <li>
    <strong>Clone this repository:</strong>
    <pre><code>git clone https://github.com/username/Wa-Tor.git
cd Wa-Tor</code></pre>
  </li>
  <li>
    <strong>Install dependencies:</strong>
    <pre><code>go get github.com/hajimehoshi/ebiten/v2</code></pre>
  </li>
  <li>
    <strong>Run the program:</strong>
    <pre><code>go run Wa-Tor.go</code></pre>
  </li>
</ol>

<hr>

<h2>Running Simulations with Different Thread Counts</h2>
<p>You can run the simulation with different thread counts for performance comparison. Use the following commands:</p>
<ul>
  <li><strong>Single-threaded:</strong>
    <pre><code>go run Wa-Tor.go -threads=1</code></pre>
  </li>
  <li><strong>Two-threaded:</strong>
    <pre><code>go run Wa-Tor.go -threads=2</code></pre>
  </li>
  <li><strong>Four-threaded:</strong>
    <pre><code>go run Wa-Tor.go -threads=4</code></pre>
  </li>
  <li><strong>Eight-threaded:</strong>
    <pre><code>go run Wa-Tor.go -threads=8</code></pre>
  </li>
</ul>

<hr>

<h2>How to Set Up a Virtual Environment and Install Dependencies</h2>
<p>If you are analysing data using Python (e.g., for plotting TPS comparisons), you can set up a virtual environment and install dependencies:</p>

<ol>
  <li>
    <strong>Create a Virtual Environment</strong>
    <ul>
      <li>On Windows:
        <pre><code>python -m venv venv</code></pre>
      </li>
      <li>On macOS and Linux:
        <pre><code>python3 -m venv venv</code></pre>
      </li>
    </ul>
  </li>

  <li>
    <strong>Activate the Virtual Environment</strong>
    <ul>
      <li>On Windows:
        <pre><code>venv\Scripts\activate</code></pre>
      </li>
      <li>On macOS and Linux:
        <pre><code>source venv/bin/activate</code></pre>
      </li>
    </ul>
  </li>

  <li>
    <strong>Install Dependencies</strong>
    <p>Install required Python packages from the <code>requirements.txt</code> file:</p>
    <pre><code>pip install -r requirements.txt</code></pre>
  </li>
</ol>

<hr>

<h2>Plotting the Results</h2>
<p>After running the simulations, you can plot and compare the TPS data using the provided Jupyter Notebook.</p>

<ol>
  <li><strong>Activate the virtual environment (if not already activated):</strong>
    <pre><code>venv\Scripts\activate   # On Windows
source venv/bin/activate  # On macOS and Linux</code></pre>
  </li>
  <li><strong>Run the Jupyter Notebook:</strong>
    <pre><code>jupyter notebook tps_analysis.ipynb</code></pre>
  </li>
</ol>