<script setup>
import { onMounted, ref } from 'vue'
import { api } from './api'

const status = ref(null)
const processing = ref(false)
const result = ref(null)
const prompt = ref('')

onMounted(async () => {
  try {
    status.value = await api.getStatus()
  } catch (error) {
    status.value = { error: 'Failed to connect to backend' }
  }
})

const processPrompt = async () => {
  if (!prompt.value.trim()) return
  
  processing.value = true
  try {
    result.value = await api.processPrompt(prompt.value)
  } catch (error) {
    result.value = { error: 'Processing failed' }
  } finally {
    processing.value = false
  }
}
</script>

<template>
  <div class="container py-4">
    <h1>AI Code Studio Pro</h1>
    
    <div class="card mb-4">
      <div class="card-header">Backend Status</div>
      <div class="card-body">
        <div v-if="status">
          <p v-if="status.error" class="text-danger">{{ status.error }}</p>
          <div v-else>
            <p>Backend: <span class="text-success">{{ status.status }}</span></p>
            <p>Message: {{ status.message }}</p>
            <p>Time: {{ status.timestamp }}</p>
          </div>
        </div>
        <div v-else>Loading...</div>
      </div>
    </div>
    
    <div class="card">
      <div class="card-header">Process Prompt</div>
      <div class="card-body">
        <div class="mb-3">
          <textarea 
            v-model="prompt" 
            class="form-control" 
            rows="3"
            placeholder="Enter your prompt here..."></textarea>
        </div>
        
        <button 
          @click="processPrompt" 
          :disabled="processing || !prompt.trim()"
          class="btn btn-primary">
          {{ processing ? 'Processing...' : 'Process' }}
        </button>
        
        <div v-if="result" class="mt-3">
          <h5>Result:</h5>
          <pre class="bg-light p-3 rounded">{{ JSON.stringify(result, null, 2) }}</pre>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
body {
  background-color: #f8f9fa;
}
.card {
  border-radius: 1rem;
  box-shadow: 0 0.5rem 1rem rgba(0,0,0,0.1);
}
</style>