# FocusBrew 🍵⏳

![Go](https://img.shields.io/badge/Go-1.20+-blue)
![TUI](https://img.shields.io/badge/TUI-BubbleTea-yellowgreen)
![Cross-Platform](https://img.shields.io/badge/Platform-Linux%20|%20Windows%20|%20macOS-blue)

**FocusBrew** is a terminal-based Pomodoro Timer built using **Golang** and the **Bubble Tea** TUI framework. This project is designed to provide a lightweight and distraction-free productivity tool for developers and terminal enthusiasts.

---

## **Project Status**

🚀 **Active Development**:  
The core MVP features have been implemented, including the Pomodoro timer, session management, and keyboard shortcuts. Contributions and feedback are welcome!

---

## **Planned Features**

### **Core Features (MVP)**
- [x] **Pomodoro Timer**:
  - Default work and break durations (25 min work, 5 min short break, 15 min long break).
- [x] **Keyboard Shortcuts**:
  - Start, pause, reset, and quit the timer using intuitive shortcuts.
- [ ] **Session Tracking**:
  - Keep track of completed work sessions.

### **Future Enhancements**
- **Customization**:
  - Allow users to configure work, short break, and long break durations.
- **Task Association**:
  - Add tasks to each Pomodoro session for better tracking.
- **Progress Bar**:
  - Visualize timer progress with a terminal-based progress bar.
- **Daily/Weekly Analytics**:
  - Display time spent on tasks or sessions.
- **Desktop Notifications**:
  - Notify users when a session ends.
- **Dark Mode**:
  - Provide light and dark themes for better visual appeal.

---

## **Goals of the Project**

1. **Lightweight**:
   - A terminal-based productivity tool that doesn't rely on a GUI environment.
2. **Cross-Platform**:
   - Compatible with Linux, Windows, and macOS.
3. **Distraction-Free**:
   - Minimalist interface to maximize focus.

---

## **Tech Stack**

- **Language**: Golang
- **Framework**: [Bubble Tea](https://github.com/charmbracelet/bubbletea) for TUI development

---

## **Installation**

To build and run **FocusBrew**, ensure you have [Go](https://go.dev/dl/) installed (version 1.20+).

1. Clone the repository:
   ```bash
   git clone https://github.com/MaskedSyntax/FocusBrew.git
   cd focusbrew
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build and run:
   ```bash
   go build -o focusbrew
   ./focusbrew
   ```

---

## **Roadmap**

### Phase 1: Initial Development
- [x] Implement a basic Pomodoro timer with work and break intervals.
- [x] Add keyboard shortcuts for basic controls (start, pause, reset, quit).

### Phase 2: Feature Expansion
- [ ] Allow customization of timer durations.
- [x] Add progress visualization.
- [ ] Add session tracking.

### Phase 3: Advanced Features
- [ ] Introduce task management.
- [ ] Add notifications and analytics.

---

## **License**
This project will be licensed under the **MIT License** once development begins.

---

## **Contact**
- **Author**: [Aftaab Siddiqui](https://github.com/MaskedSyntax)
- **Repository**: [FocusBrew](https://github.com/MaskedSyntax/FocusBrew)

Feel free to reach out with ideas, feedback, or questions!
