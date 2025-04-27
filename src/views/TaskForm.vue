<template>
  <div class="task-form">
    <h2>{{ isEditing ? 'Edit Task' : 'Create New Task' }}</h2>
    
    <div v-if="loading" class="loading">Loading task data...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    
    <form v-else @submit.prevent="saveTask">
      <div class="form-group">
        <label for="title">Title</label>
        <input
          id="title"
          v-model="task.title"
          type="text"
          required
          placeholder="Enter task title"
        />
      </div>
      
      <div class="form-group">
        <label for="description">Description</label>
        <textarea
          id="description"
          v-model="task.description"
          placeholder="Enter task description"
          rows="4"
        ></textarea>
      </div>
      
      <div class="form-group">
        <label for="status">Status</label>
        <select id="status" v-model="task.status">
          <option value="pending">Pending</option>
          <option value="in-progress">In Progress</option>
          <option value="completed">Completed</option>
        </select>
      </div>
      
      <div class="form-actions">
        <button type="button" @click="$router.push('/')" class="btn cancel">
          Cancel
        </button>
        <button type="submit" class="btn save">
          {{ isEditing ? 'Update Task' : 'Create Task' }}
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import TaskService from '@/services/TaskService'

export default {
  name: 'TaskForm',
  props: {
    id: {
      type: String,
      default: null
    }
  },
  
  data() {
    return {
      task: {
        title: '',
        description: '',
        status: 'pending'
      },
      loading: false,
      error: null
    }
  },
  
  computed: {
    isEditing() {
      return !!this.id
    }
  },
  
  async created() {
    if (this.isEditing) {
      await this.fetchTask()
    }
  },
  
  methods: {
    async fetchTask() {
      this.loading = true
      this.error = null
      
      try {
        const response = await TaskService.getTask(this.id)
        this.task = response.data
      } catch (err) {
        this.error = 'Error loading task: ' + (err.response?.data || err.message)
        console.error(err)
      } finally {
        this.loading = false
      }
    },
    
    async saveTask() {
      try {
        if (this.isEditing) {
          await TaskService.updateTask(this.task)
        } else {
          await TaskService.createTask(this.task)
        }
        this.$router.push('/')
      } catch (err) {
        this.error = 'Error saving task: ' + (err.response?.data || err.message)
        console.error(err)
      }
    }
  }
}
</script>

<style scoped>
.task-form {
  max-width: 600px;
  margin: 0 auto;
  text-align: left;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: bold;
}

.form-group input,
.form-group textarea,
.form-group select {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 30px;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
  font-size: 16px;
}

.btn.save {
  background-color: #42b983;
  color: white;
}

.btn.cancel {
  background-color: #f5f5f5;
  color: #333;
}

.loading, .error {
  text-align: center;
  padding: 20px;
  color: #666;
}

.error {
  color: #f44336;
}
</style>
