```sql
 
 CREATE TABLE users{
    id UUID PRIMARY KEY,
    name    VARCHAR(300) NOT NULL,
    email   VARCHAR(300) NOT NULL,
    phone   VARCHAR(20),

    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 }

 CREATE TABLE notebooks{
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),

    icon        VARCHAR(50),
    name        VARCHAR(300) NOT NULL,
    image       VARCHAR(900),
    description VARCHAR(900),

    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP  
 }

 CREATE TABLE meta_contents {
    id  UUID PRIMARY KEY,
    notebook_id UUID REFERENCES notebooks(id),
    
    icon        VARCHAR(50),
    name        VARCHAR(300) NOT NULL,
    user_id UUID NOT NULL,
    
 
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP
 }
 CREATE TABLE meta_tag_contents {
    content_id  UUID REFERENCES meta_contents(id),
    notebook_id UUID REFERENCES notebooks(id),
    
    tag_id      UUID NOT NULL REFERENCES tags(id),
 
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP
 }

  CREATE TABLE tags {
    id  UUID PRIMARY KEY,
    
    name        VARCHAR(300) NOT NULL,
    color       VARCHAR(6) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    
 
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP
 }


 CREATE TABLE nodes_contents{
    id  UUID PRIMARY KEY,
    content_id UUID REFERENCES meta_contents(id),
    
    user_id UUID NOT NULL,
    notebook_id UUID NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP
 }
-- SELECT * FROM nodes_contents c LEFT JOIN meta_tag_contents t on c.content_id = t.content_id


```

# Criar banco
```
docker-compose exec postgres psql -U postgres -c "CREATE DATABASE surfbook_dev;

```