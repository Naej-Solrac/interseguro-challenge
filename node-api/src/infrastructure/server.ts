// Infraestructura: Configuración del servidor Express
import express from 'express';
import swaggerUi from 'swagger-ui-express';
import { swaggerSpec } from './swagger/swagger.config';
import matrixStatsRoutes from './matrixStats.routes';

const app = express();

app.use(express.json());

// Documentación Swagger
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerSpec));

// Ruta para obtener el spec en JSON
app.get('/api-docs.json', (_req, res) => {
  res.setHeader('Content-Type', 'application/json');
  res.send(swaggerSpec);
});

app.use('/api', matrixStatsRoutes);

// Middleware de manejo de rutas no encontradas
app.use((_req, res) => {
  res.status(404).json({ error: 'Ruta no encontrada' });
});

export default app;
