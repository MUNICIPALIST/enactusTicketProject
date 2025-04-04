FROM node:18-alpine as builder

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# Используем nginx для раздачи статики
FROM nginx:alpine

# Копируем собранные файлы в директорию nginx
COPY --from=builder /app/dist /usr/share/nginx/html/aidana

# Копируем конфигурацию nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 5173

CMD ["nginx", "-g", "daemon off;"]
