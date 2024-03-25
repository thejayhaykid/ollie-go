FROM node:lts-alpine AS runtime
WORKDIR /app

RUN corepack enable

COPY . .

RUN yarn install
RUN yarn run build

ENV HOST=0.0.0.0
ENV PORT=4321
EXPOSE 4321
CMD node ./dist/server/entry.mjs