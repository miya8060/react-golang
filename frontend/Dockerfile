FROM node:20-alpine

# pnpmのインストール
RUN corepack enable && corepack prepare pnpm@latest --activate

WORKDIR /app

# pnpm-lock.yamlをコピー
COPY pnpm-lock.yaml ./

# パッケージのインストール
COPY package.json ./
RUN pnpm install --frozen-lockfile

COPY . .

EXPOSE 5173

CMD ["pnpm", "dev"]