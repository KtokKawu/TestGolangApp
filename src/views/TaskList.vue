<template>
  <div class="task-list">
    <h2>My Tasks</h2>
    
    <div v-if="loading" class="loading">Loading tasks...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else-if="tasks.length === 0" class="no-tasks">
      No tasks found. Create a new one!
    </div>
    
    <div v-else class="tasks">
      <div v-for="task in tasks" :key="task.id" class="task-item">
        <div class="task-content">
          <h3>{{ task.title }}</h3>
          <p>{{ task.description }}</p>
          <div class="task-status" :class="task.status">
            Status: {{ task.status }}
          </div>
        </div>
        <div class="task-actions">
          <button @click="editTask(task)" class="btn edit">Edit</button>
          <button @click="confirmDelete(task)" class="btn delete">Delete</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TaskService from '@/services/TaskService'

export default {
  name: 'TaskList',
  
  data() {
    return {
      tasks: [],
      loading: true,
      error: null
    }
  },
  
  created() {
    this.fetchTasks()
  },
  
  methods: {
    async fetchTasks() {
      this.loading = true
      this.error = null
      
      try {
        const response = await TaskService.getTasks()
        this.tasks = response.data
      } catch (err) {
        this.error = 'Error loading tasks: ' + (err.response?.data || err.message)
        console.error(err)
      } finally {
        this.loading = false
      }
    },
    
    editTask(task) {
      this.$router.push(`/edit/${task.id}`)
    },
    
    async confirmDelete(task) {
      if (confirm(`Are you sure you want to delete "${task.title}"?`)) {
        try {
          await TaskService.deleteTask(task.id)
          // Remove from local array
          this.tasks = this.tasks.filter(t => t.id !== task.id)
        } catch (err) {
          alert('Error deleting task: ' + (err.response?.data || err.message))
          console.error(err)
        }
      }
    }
  }
}
</script>

<style scoped>
.task-list {
  text-align: left;
}

.task-item {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 15px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.task-content {
  flex: 1;
}

.task-content h3 {
  margin-top: 0;
  margin-bottom: 8px;
}

.task-content p {
  color: #666;
  margin-bottom: 8px;
}

.task-status {
  display: inline-block;
  padding: 5px 10px;
  border-radius: 4px;
  font-size: 14px;
  font-weight: bold;
}

.task-status.pending {
  background-color: #ffecb3;
  color: #ffa000;
}

.task-status.in-progress {
  background-color: #bbdefb;
  color: #1976d2;
}

.task-status.completed {
  background-color: #c8e6c9;
  color: #388e3c;
}

.task-actions {
  display: flex;
  gap: 8px;
}

.btn {
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}

.btn.edit {
  background-color: #42b983;
  color: white;
}

.btn.delete {
  background-color: #f44336;
  color: white;
}

.loading, .error, .no-tasks {
  text-align: center;
  padding: 20px;
  color: #666;
}

.error {
  color: #f44336;
}
</style>
