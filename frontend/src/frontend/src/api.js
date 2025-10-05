const API_BASE = import.meta.env.DEV 
  ? '/api' 
  : `${window.location.origin}/api`;

export const api = {
  getStatus: async () => {
    try {
      const response = await fetch(`${API_BASE}/status`);
      if (!response.ok) throw new Error('Network response was not ok');
      return await response.json();
    } catch (error) {
      console.error('API Error:', error);
      throw error;
    }
  },
  
  processPrompt: async (prompt) => {
    try {
      const response = await fetch(`${API_BASE}/process`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ prompt })
      });
      
      if (!response.ok) throw new Error('Network response was not ok');
      return await response.json();
    } catch (error) {
      console.error('API Error:', error);
      throw error;
    }
  }
};