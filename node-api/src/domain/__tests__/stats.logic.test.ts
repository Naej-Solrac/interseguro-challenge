import { calculateMatrixStatistics } from '../stats.logic';

describe('calculateMatrixStatistics', () => {
  it('calcula correctamente las estadísticas de una matriz', () => {
    const matrix = [
      [1, 2],
      [3, 4]
    ];
    const result = calculateMatrixStatistics(matrix);
    expect(result).toEqual({
      max: 4,
      min: 1,
      avg: 2.5,
      sum: 10,
      isDiagonal: false
    });
  });

  it('detecta matriz diagonal correctamente', () => {
    const matrix = [
      [5, 0, 0],
      [0, 7, 0],
      [0, 0, 9]
    ];
    const result = calculateMatrixStatistics(matrix);
    expect(result.isDiagonal).toBe(true);
  });

  it('lanza error si la matriz no es válida', () => {
    expect(() => calculateMatrixStatistics([])).toThrow();
    expect(() => calculateMatrixStatistics([1 as any,2 as any,3 as any])).toThrow();
    expect(() => calculateMatrixStatistics([[1,2],[3]])).toThrow();
    expect(() => calculateMatrixStatistics([[1,2],[null as any,4]])).toThrow();
  });
});
