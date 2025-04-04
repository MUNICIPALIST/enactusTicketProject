FROM node:18-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# Устанавливаем глобально serve для раздачи статики
RUN npm install -g serve

EXPOSE 5173

# Используем serve для раздачи собранного приложения
CMD ["serve", "-s", "dist", "-l", "5173"]
