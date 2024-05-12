FROM node:20-alpine AS base
WORKDIR /app/frontend

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

RUN pnpm add -g serve

COPY package.json pnpm-lock.yaml ./
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile

ARG BACKEND_PORT

COPY . .
RUN pnpm build

CMD ["sh", "-c", "serve ./build -l ${FRONTEND_PORT}"]
