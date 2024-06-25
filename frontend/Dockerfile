# Stage 1: Build React app
FROM node:14 as build

WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install --silent  # Use --silent to suppress npm output for cleaner logs

# Copy the rest of the application code
COPY . .

# Build the React app
RUN npm run build

# Stage 2: Serve React app using nginx
FROM nginx:alpine

# Remove default nginx website
RUN rm -rf /usr/share/nginx/html/*

# Copy built files to nginx default location
COPY --from=build /app/build /usr/share/nginx/html

# Copy nginx configuration
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8000
CMD ["nginx", "-g", "daemon off;"]
