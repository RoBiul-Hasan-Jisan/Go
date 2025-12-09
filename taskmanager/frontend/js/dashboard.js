// Dashboard JavaScript
const apiBaseUrl = 'http://localhost:8080/api';
let currentTasks = [];

// Initialize dashboard
document.addEventListener('DOMContentLoaded', () => {
    checkAuth();
    loadUserInfo();
    loadTasks();
    setupEventListeners();
    showSection('tasks');
});

// Authentication check
function checkAuth() {
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = '/login';
    }
}

// Logout function
function logout() {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    window.location.href = '/';
}

// Load user information
async function loadUserInfo() {
    try {
        const response = await fetch(`${apiBaseUrl}/user`, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });
        
        if (response.ok) {
            const user = await response.json();
            document.getElementById('username').textContent = user.username;
            
            // Update profile section
            const profileHtml = `
                <div class="profile-info">
                    <div class="info-item">
                        <strong>Username:</strong>
                        <span>${user.username}</span>
                    </div>
                    <div class="info-item">
                        <strong>Email:</strong>
                        <span>${user.email}</span>
                    </div>
                    <div class="info-item">
                        <strong>Member Since:</strong>
                        <span>${new Date(user.created_at).toLocaleDateString()}</span>
                    </div>
                </div>
            `;
            document.getElementById('profile-info').innerHTML = profileHtml;
        }
    } catch (error) {
        console.error('Failed to load user info:', error);
    }
}

// Load tasks from API
async function loadTasks() {
    try {
        const response = await fetch(`${apiBaseUrl}/tasks`, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });
        
        if (response.ok) {
            currentTasks = await response.json();
            displayTasks(currentTasks);
            updateStatistics();
        }
    } catch (error) {
        console.error('Failed to load tasks:', error);
        document.getElementById('tasks-container').innerHTML = `
            <div class="error">
                <i class="fas fa-exclamation-circle"></i>
                Failed to load tasks. Please try again.
            </div>
        `;
    }
}

// Display tasks in grid
function displayTasks(tasks) {
    const container = document.getElementById('tasks-container');
    
    if (tasks.length === 0) {
        container.innerHTML = `
            <div class="empty-state">
                <i class="fas fa-clipboard-list"></i>
                <h3>No tasks found</h3>
                <p>Create your first task to get started!</p>
                <button onclick="showSection('new')" class="btn-primary">
                    <i class="fas fa-plus"></i> Create Task
                </button>
            </div>
        `;
        return;
    }
    
    container.innerHTML = tasks.map(task => `
        <div class="task-card priority-${task.priority}">
            <div class="task-header">
                <h3 class="task-title">${task.title}</h3>
                <span class="task-priority priority-${task.priority}">
                    ${task.priority.charAt(0).toUpperCase() + task.priority.slice(1)}
                </span>
            </div>
            
            <div class="task-description">
                ${task.description || 'No description provided.'}
            </div>
            
            <div class="task-meta">
                <span class="task-status status-${task.status}">
                    <i class="fas fa-${getStatusIcon(task.status)}"></i>
                    ${task.status.replace('_', ' ').toUpperCase()}
                </span>
                ${task.due_date ? `
                    <span class="task-due">
                        <i class="far fa-calendar"></i>
                        Due: ${new Date(task.due_date).toLocaleDateString()}
                    </span>
                ` : ''}
            </div>
            
            <div class="task-footer">
                <span class="task-date">
                    Created: ${new Date(task.created_at).toLocaleDateString()}
                </span>
                <div class="task-actions">
                    <button onclick="editTask(${task.id})" class="btn-icon" title="Edit">
                        <i class="fas fa-edit"></i>
                    </button>
                    <button onclick="deleteTask(${task.id})" class="btn-icon" title="Delete">
                        <i class="fas fa-trash"></i>
                    </button>
                    <button onclick="toggleTaskStatus(${task.id}, '${task.status}')" class="btn-icon" title="Toggle Status">
                        <i class="fas fa-check"></i>
                    </button>
                </div>
            </div>
        </div>
    `).join('');
}

// Get status icon
function getStatusIcon(status) {
    switch(status) {
        case 'pending': return 'clock';
        case 'in_progress': return 'spinner';
        case 'completed': return 'check-circle';
        default: return 'question';
    }
}

// Create new task
document.getElementById('taskForm')?.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const formData = {
        title: document.getElementById('title').value,
        description: document.getElementById('description').value,
        priority: document.getElementById('priority').value,
        status: document.getElementById('status').value,
        due_date: document.getElementById('due_date').value || null
    };
    
    try {
        const response = await fetch(`${apiBaseUrl}/tasks`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
            body: JSON.stringify(formData)
        });
        
        if (response.ok) {
            const newTask = await response.json();
            currentTasks.unshift(newTask);
            displayTasks(currentTasks);
            updateStatistics();
            showSection('tasks');
            document.getElementById('taskForm').reset();
            
            // Show success message
            showNotification('Task created successfully!', 'success');
        }
    } catch (error) {
        console.error('Failed to create task:', error);
        showNotification('Failed to create task. Please try again.', 'error');
    }
});

// Edit task
async function editTask(taskId) {
    const task = currentTasks.find(t => t.id === taskId);
    if (!task) return;
    
    const modalBody = document.getElementById('modal-body');
    modalBody.innerHTML = `
        <h3>Edit Task</h3>
        <form id="editTaskForm">
            <div class="form-group">
                <label>Title</label>
                <input type="text" id="edit-title" value="${task.title}" required>
            </div>
            <div class="form-group">
                <label>Description</label>
                <textarea id="edit-description" rows="4">${task.description || ''}</textarea>
            </div>
            <div class="form-row">
                <div class="form-group">
                    <label>Priority</label>
                    <select id="edit-priority">
                        <option value="low" ${task.priority === 'low' ? 'selected' : ''}>Low</option>
                        <option value="medium" ${task.priority === 'medium' ? 'selected' : ''}>Medium</option>
                        <option value="high" ${task.priority === 'high' ? 'selected' : ''}>High</option>
                    </select>
                </div>
                <div class="form-group">
                    <label>Status</label>
                    <select id="edit-status">
                        <option value="pending" ${task.status === 'pending' ? 'selected' : ''}>Pending</option>
                        <option value="in_progress" ${task.status === 'in_progress' ? 'selected' : ''}>In Progress</option>
                        <option value="completed" ${task.status === 'completed' ? 'selected' : ''}>Completed</option>
                    </select>
                </div>
            </div>
            <div class="form-group">
                <label>Due Date</label>
                <input type="date" id="edit-due_date" value="${task.due_date ? task.due_date.split('T')[0] : ''}">
            </div>
            <div class="form-actions">
                <button type="button" onclick="closeModal()" class="btn-secondary">Cancel</button>
                <button type="submit" class="btn-primary">Save Changes</button>
            </div>
        </form>
    `;
    
    document.getElementById('editTaskForm').addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const updatedData = {
            title: document.getElementById('edit-title').value,
            description: document.getElementById('edit-description').value,
            priority: document.getElementById('edit-priority').value,
            status: document.getElementById('edit-status').value,
            due_date: document.getElementById('edit-due_date').value || null
        };
        
        try {
            const response = await fetch(`${apiBaseUrl}/tasks/${taskId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                },
                body: JSON.stringify(updatedData)
            });
            
            if (response.ok) {
                const updatedTask = await response.json();
                const index = currentTasks.findIndex(t => t.id === taskId);
                currentTasks[index] = updatedTask;
                displayTasks(currentTasks);
                updateStatistics();
                closeModal();
                showNotification('Task updated successfully!', 'success');
            }
        } catch (error) {
            console.error('Failed to update task:', error);
            showNotification('Failed to update task. Please try again.', 'error');
        }
    });
    
    openModal();
}

// Delete task
async function deleteTask(taskId) {
    if (!confirm('Are you sure you want to delete this task?')) return;
    
    try {
        const response = await fetch(`${apiBaseUrl}/tasks/${taskId}`, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });
        
        if (response.ok) {
            currentTasks = currentTasks.filter(task => task.id !== taskId);
            displayTasks(currentTasks);
            updateStatistics();
            showNotification('Task deleted successfully!', 'success');
        }
    } catch (error) {
        console.error('Failed to delete task:', error);
        showNotification('Failed to delete task. Please try again.', 'error');
    }
}

// Toggle task status
async function toggleTaskStatus(taskId, currentStatus) {
    let newStatus;
    switch(currentStatus) {
        case 'pending': newStatus = 'in_progress'; break;
        case 'in_progress': newStatus = 'completed'; break;
        default: newStatus = 'pending';
    }
    
    try {
        const response = await fetch(`${apiBaseUrl}/tasks/${taskId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
            body: JSON.stringify({ status: newStatus })
        });
        
        if (response.ok) {
            const updatedTask = await response.json();
            const index = currentTasks.findIndex(t => t.id === taskId);
            currentTasks[index] = updatedTask;
            displayTasks(currentTasks);
            updateStatistics();
            showNotification('Task status updated!', 'success');
        }
    } catch (error) {
        console.error('Failed to update task status:', error);
        showNotification('Failed to update task status. Please try again.', 'error');
    }
}

// Update statistics
function updateStatistics() {
    const pending = currentTasks.filter(t => t.status === 'pending').length;
    const inProgress = currentTasks.filter(t => t.status === 'in_progress').length;
    const completed = currentTasks.filter(t => t.status === 'completed').length;
    const total = currentTasks.length;
    
    document.getElementById('pending-count').textContent = pending;
    document.getElementById('progress-count').textContent = inProgress;
    document.getElementById('completed-count').textContent = completed;
    document.getElementById('total-count').textContent = total;
    
    // Update chart if it exists
    updateChart(pending, inProgress, completed);
}

// Update chart
function updateChart(pending, inProgress, completed) {
    const ctx = document.getElementById('tasksChart');
    if (!ctx) return;
    
    const chart = new Chart(ctx.getContext('2d'), {
        type: 'doughnut',
        data: {
            labels: ['Pending', 'In Progress', 'Completed'],
            datasets: [{
                data: [pending, inProgress, completed],
                backgroundColor: [
                    '#f59e0b',  // Orange
                    '#3b82f6',  // Blue
                    '#10b981'   // Green
                ],
                borderWidth: 2,
                borderColor: '#fff'
            }]
        },
        options: {
            responsive: true,
            plugins: {
                legend: {
                    position: 'bottom'
                }
            }
        }
    });
}

// Section navigation
function showSection(sectionId) {
    // Hide all sections
    document.querySelectorAll('.content-section').forEach(section => {
        section.classList.remove('active');
    });
    
    // Remove active from all nav links
    document.querySelectorAll('.sidebar-nav a').forEach(link => {
        link.classList.remove('active');
    });
    
    // Show selected section
    document.getElementById(`${sectionId}-section`).classList.add('active');
    
    // Activate corresponding nav link
    document.getElementById(`nav-${sectionId}`).classList.add('active');
    
    // Update page title
    const titles = {
        'tasks': 'My Tasks',
        'new': 'New Task',
        'profile': 'My Profile',
        'stats': 'Statistics'
    };
    document.getElementById('page-title').textContent = titles[sectionId];
}

// Modal functions
function openModal() {
    document.getElementById('taskModal').style.display = 'block';
}

function closeModal() {
    document.getElementById('taskModal').style.display = 'none';
}

// Notification
function showNotification(message, type) {
    const notification = document.createElement('div');
    notification.className = `notification ${type}`;
    notification.innerHTML = `
        <i class="fas fa-${type === 'success' ? 'check-circle' : 'exclamation-circle'}"></i>
        <span>${message}</span>
    `;
    
    document.body.appendChild(notification);
    
    setTimeout(() => {
        notification.classList.add('show');
    }, 10);
    
    setTimeout(() => {
        notification.classList.remove('show');
        setTimeout(() => {
            document.body.removeChild(notification);
        }, 300);
    }, 3000);
}

// Setup event listeners
function setupEventListeners() {
    // Navigation
    document.querySelectorAll('.sidebar-nav a').forEach(link => {
        link.addEventListener('click', (e) => {
            e.preventDefault();
            const sectionId = link.id.replace('nav-', '');
            showSection(sectionId);
        });
    });
    
    // New task button
    document.getElementById('new-task-btn')?.addEventListener('click', () => {
        showSection('new');
    });
    
    // Filter buttons
    document.querySelectorAll('.filter-btn').forEach(btn => {
        btn.addEventListener('click', () => {
            document.querySelectorAll('.filter-btn').forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            
            const filter = btn.dataset.filter;
            let filteredTasks = currentTasks;
            
            if (filter !== 'all') {
                filteredTasks = currentTasks.filter(task => task.status === filter);
            }
            
            displayTasks(filteredTasks);
        });
    });
    
    // Search
    document.getElementById('search')?.addEventListener('input', (e) => {
        const searchTerm = e.target.value.toLowerCase();
        const filteredTasks = currentTasks.filter(task => 
            task.title.toLowerCase().includes(searchTerm) ||
            task.description?.toLowerCase().includes(searchTerm)
        );
        displayTasks(filteredTasks);
    });
    
    // Modal close
    document.querySelector('.close')?.addEventListener('click', closeModal);
    window.addEventListener('click', (e) => {
        if (e.target === document.getElementById('taskModal')) {
            closeModal();
        }
    });
}

// Add notification styles
const style = document.createElement('style');
style.textContent = `
    .notification {
        position: fixed;
        top: 20px;
        right: 20px;
        padding: 1rem 1.5rem;
        border-radius: 8px;
        display: flex;
        align-items: center;
        gap: 0.75rem;
        box-shadow: 0 4px 12px rgba(0,0,0,0.15);
        transform: translateX(100%);
        transition: transform 0.3s ease;
        z-index: 10000;
        max-width: 300px;
    }
    
    .notification.show {
        transform: translateX(0);
    }
    
    .notification.success {
        background: #10b981;
        color: white;
        border-left: 4px solid #059669;
    }
    
    .notification.error {
        background: #ef4444;
        color: white;
        border-left: 4px solid #dc2626;
    }
    
    .notification i {
        font-size: 1.2rem;
    }
`;
document.head.appendChild(style);