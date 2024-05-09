## Product API with CRUD Operations (Fiber, PostgreSQL)

This project implements a RESTful API for managing products using Fiber, a high-performance web framework for Go, and PostgreSQL, a powerful and versatile open-source relational database. It offers comprehensive CRUD (Create, Read, Update, Delete) functionality for product data, enabling your application to interact with product information effectively.

### Key Features:

- CRUD Operations: Create, retrieve, update, and delete product data.
- Fiber Framework: Leverages Fiber's speed and flexibility for efficient API development.
- PostgreSQL Database: Employs PostgreSQL for reliable and scalable data storage.

### API Endpoints:

/products (GET): Retrieves a list of all products.
/product/:id (GET): Fetches a specific product by its ID.
/create (POST): Creates a new product.
/product/:id/edit (PUT): Updates an existing product.
/product/:id (DELETE): Deletes a product.
