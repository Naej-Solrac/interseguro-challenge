// Infraestructura: Configuración del servidor Express
import express from 'express';
import matrixStatsRoutes from './matrixStats.routes';

const app = express();

app.use(express.json());
app.use('/api', matrixStatsRoutes);

// Middleware de manejo de rutas no encontradas
app.use((_req, res) => {
  res.status(404).json({ error: 'Ruta no encontrada' });
});

export default app;
