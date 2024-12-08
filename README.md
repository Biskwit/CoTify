# CoTify Proxy Service (Draft)

Proxy service that add a Chain of Thought logic for any openAI compatible endpoints (like Groq <3)

## Curl Example

```bash
curl http://localhost:3000/chat/completions -s \
-H "Content-Type: application/json" \
-H "Authorization: Bearer gsk_..." \
-d '{
"model": "llama3-8b-8192",
"messages": [{
    "role": "user",
    "content": "Explain the importance of fast language models"
}]
}'
```

# Thanks

Inspired by [g1](https://github.com/bklieger-groq/g1)
