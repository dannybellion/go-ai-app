import React, { useState, useEffect, useRef } from 'react';
import './app.css';

const API_URL = 'http://localhost:8080';

function App() {
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const messagesEndRef = useRef(null);

  useEffect(() => {
    fetchHistory();
  }, []);

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  };

  const fetchHistory = async () => {
    try {
      const response = await fetch(`${API_URL}/history`);
      if (!response.ok) {
        throw new Error('Failed to fetch history');
      }
      const data = await response.json();
      
      // Transform history into message format
      const historyMessages = data.flatMap(item => [
        { role: 'user', content: item.prompt },
        { role: 'assistant', content: item.response }
      ]);
      
      setMessages(historyMessages);
    } catch (err) {
      console.error('Error fetching history:', err);
      setError('Could not load chat history. Please try again later.');
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!input.trim()) return;
    
    const userMessage = { role: 'user', content: input };
    setMessages(prev => [...prev, userMessage]);
    setInput('');
    setLoading(true);
    setError(null);
    
    try {
      const response = await fetch(`${API_URL}/chat`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ prompt: input }),
      });
      
      if (!response.ok) {
        throw new Error('Failed to get response');
      }
      
      const data = await response.json();
      const aiMessage = { role: 'assistant', content: data.response };
      
      setMessages(prev => [...prev, aiMessage]);
    } catch (err) {
      console.error('Error:', err);
      setError('Something went wrong. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const formatMessage = (content) => {
    return content.split('\n').map((line, i) => (
      <span key={i}>
        {line}
        <br />
      </span>
    ));
  };

  return (
    <div className="app">
      <header className="header">
        <h1>AI Chat Assistant</h1>
      </header>
      
      <main className="chat-container">
        <div className="messages">
          {messages.length === 0 && !loading && !error && (
            <div className="empty-state">
              <p>No messages yet. Start a conversation!</p>
            </div>
          )}
          
          {messages.map((msg, index) => (
            <div 
              key={index} 
              className={`message ${msg.role === 'user' ? 'user-message' : 'ai-message'}`}
            >
              <div className="message-bubble">
                {formatMessage(msg.content)}
              </div>
            </div>
          ))}
          
          {loading && (
            <div className="message ai-message">
              <div className="message-bubble loading">
                <div className="typing-indicator">
                  <span></span>
                  <span></span>
                  <span></span>
                </div>
              </div>
            </div>
          )}
          
          {error && (
            <div className="error-message">
              <p>{error}</p>
            </div>
          )}
          
          <div ref={messagesEndRef} />
        </div>
      </main>
      
      <footer className="input-area">
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Type a message..."
            disabled={loading}
          />
          <button 
            type="submit" 
            disabled={loading || !input.trim()}
          >
            Send
          </button>
        </form>
      </footer>
    </div>
  );
}

export default App;