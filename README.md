# CPU Scheduling Simulator

A high-performance command-line CPU Scheduling Simulator built for **Linux** environments. This tool simulates how an Operating System's CPU scheduler manages processes using various scheduling algorithms, calculating key performance metrics like waiting time, turnaround time, and CPU utilization.

## 🚀 Features

* **Multiple Algorithms Supported:**
  * First-Come, First-Served (FCFS)
  * Shortest Job First (SJF) - Non-preemptive
  * Shortest Remaining Time First (SRTF) - Preemptive SJF
  * Round Robin (RR) with custom Time Quantum
  * Priority Scheduling (Preemptive & Non-preemptive)
* **Detailed Evaluation:** Generates a real-time Gantt chart visualization in the terminal.
* **Performance Analysis:** Calculates Average Waiting Time (AWT) and Average Turnaround Time (ATT) for comparison.
* **File I/O:** Supports reading process workflows directly from text configuration files.

## 🛠️ Prerequisites

Ensure your Linux system has the standard GNU Compiler Collection (GCC) tools installed:

```bash
sudo apt update
sudo apt install build-essential
```

## 📂 Project Structure

```text
├── src/
│   ├── main.cpp         # Entry point and CLI logic
│   ├── scheduler.cpp    # Core scheduling algorithms
│   └── scheduler.h      # Headers and data structures
├── input.txt            # Sample process dataset
├── Makefile             # Build automation script
└── README.md            # Documentation
```

## ⚙️ How to Build

A `Makefile` is included for easy compilation. Run the following commands in your Linux terminal:

1. **Compile the project:**
   ```bash
   make
   ```
2. **Clean build files (optional):**
   ```bash
   make clean
   ```

*(If a Makefile is not used, compile manually using: `g++ -std=c++11 src/*.cpp -o simulator`)*

## 🚀 Running the Simulator

Execute the binary from the terminal. The program expects a process configuration file as an argument:

```bash
./simulator input.txt
```

### Input File Format (`input.txt`)
Provide process details in the following space-separated format: `[Process_ID] [Arrival_Time] [Burst_Time] [Priority]`

```text
1 0 8 3
2 1 4 1
3 2 9 2
4 3 5 4
```

## 📊 Sample Output

Upon execution, the terminal displays the scheduling sequence and metrics:

```text
Gantt Chart:

| P1 (0-8) | P2 (8-12) | P4 (12-17) | P3 (17-26) |

=======================================================
Process ID | Arrival | Burst | Waiting Time | Turnaround
=======================================================
P1         | 0       | 8     | 0            | 8
P2         | 1       | 4     | 7            | 11
P3         | 2       | 9     | 15           | 24
P4         | 3       | 5     | 9            | 14
=======================================================

Average Waiting Time: 7.75 ms
Average Turnaround Time: 14.25 ms
```

## 👥 Authors

* **Your Name** - *Initial Work* - [YourGitHub](https://github.com)
