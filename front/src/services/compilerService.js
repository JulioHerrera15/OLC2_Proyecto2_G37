import CryptoJS from 'crypto-js';

class CompilerService {
  constructor() {
    this.baseURL = '/api';
    this.cache = new Map();
    this.cacheTTL = 5 * 60 * 1000; // 5 minutos como en el legacy
    this.executionSemaphore = 0;
    this.maxConcurrentExecutions = 3; // Como en el legacy
  }

  // Genera hash del código para cache
  generateCodeHash(code) {
    return CryptoJS.SHA256(code).toString();
  }

  // Verifica si el cache es válido
  isCacheValid(cacheKey) {
    const cached = this.cache.get(cacheKey);
    if (!cached) return false;
    
    return (Date.now() - cached.timestamp) < this.cacheTTL;
  }

  // Control de concurrencia
  async acquireSemaphore() {
    while (this.executionSemaphore >= this.maxConcurrentExecutions) {
      await new Promise(resolve => setTimeout(resolve, 100));
    }
    this.executionSemaphore++;
  }

  releaseSemaphore() {
    this.executionSemaphore--;
  }

  // Ejecutar código
  async executeCode(code) {
    // Advertencia para código grande
    if (code.length > 50 * 1024) {
      console.warn('Código muy grande, puede tardar...');
    }

    await this.acquireSemaphore();
    
    try {
      const response = await fetch(`${this.baseURL}/execute`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ code }),
        signal: AbortSignal.timeout(30000) // 30 segundos como en el legacy
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const result = await response.json();
      return result;
    } catch (error) {
      console.error('Error ejecutando código:', error);
      throw error;
    } finally {
      this.releaseSemaphore();
    }
  }

  // Generar reporte de errores
  async generateErrorReport(code) {
    const cacheKey = `errors_${this.generateCodeHash(code)}`;
    
    if (this.isCacheValid(cacheKey)) {
      console.log('📋 Cache hit: Usando reporte de errores en cache');
      return this.cache.get(cacheKey).data;
    }

    try {
      // CORREGIDO: Usar el endpoint del servidor
      const response = await fetch(`${this.baseURL}/report/errors`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ code }),
        signal: AbortSignal.timeout(30000)
      });

      if (!response.ok) {
        throw new Error(`Error generando reporte de errores: ${response.status}`);
      }

      const result = await response.json();
      
      // Guardar en cache
      this.cache.set(cacheKey, {
        data: result.htmlContent,
        timestamp: Date.now()
      });

      return result.htmlContent;
    } catch (error) {
      console.error('Error generando reporte de errores:', error);
      throw error;
    }
  }

  // Generar tabla de símbolos
  async generateSymbolReport(code) {
    const cacheKey = `symbols_${this.generateCodeHash(code)}`;
    
    if (this.isCacheValid(cacheKey)) {
      console.log('🔤 Cache hit: Usando tabla de símbolos en cache');
      return this.cache.get(cacheKey).data;
    }

    try {
      const response = await fetch(`${this.baseURL}/report/symbols`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ code }),
        signal: AbortSignal.timeout(30000)
      });

      if (!response.ok) {
        throw new Error(`Error generando tabla de símbolos: ${response.status}`);
      }

      const result = await response.json();
      
      // Guardar en cache
      this.cache.set(cacheKey, {
        data: result.htmlContent,
        timestamp: Date.now()
      });

      return result.htmlContent;
    } catch (error) {
      console.error('Error generando tabla de símbolos:', error);
      throw error;
    }
  }

  // Generar CST
  async generateCSTReport(code) {
    const cacheKey = `cst_${this.generateCodeHash(code)}`;
    
    if (this.isCacheValid(cacheKey)) {
      console.log('🌳 Cache hit: Usando árbol CST en cache');
      return this.cache.get(cacheKey).data;
    }

    try {
      const response = await fetch(`${this.baseURL}/cst`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ 
          code,
          grammar: 'Language.g4',
          startRule: 'program'
        }),
        signal: AbortSignal.timeout(60000) // 60 segundos para CST
      });

      if (!response.ok) {
        throw new Error(`Error generando CST: ${response.status}`);
      }

      const result = await response.json();
      
      // Guardar en cache
      this.cache.set(cacheKey, {
        data: result.htmlContent,
        timestamp: Date.now()
      });

      return result.htmlContent;
    } catch (error) {
      console.error('Error generando CST:', error);
      throw error;
    }
  }

  // Limpiar cache
  clearCache() {
    this.cache.clear();
    console.log('🧹 Cache limpiado completamente');
  }

  // Invalidar cache específico
  invalidateCache(codeHash) {
    const keysToDelete = [];
    for (const key of this.cache.keys()) {
      if (key.includes(codeHash)) {
        keysToDelete.push(key);
      }
    }
    keysToDelete.forEach(key => this.cache.delete(key));
    if (keysToDelete.length > 0) {
      console.log(`🗑️ Cache invalidado: ${keysToDelete.length} entradas eliminadas`);
    }
  }

  // Obtener estadísticas del cache
  getCacheStats() {
    return {
      size: this.cache.size,
      maxSize: 100, // Límite arbitrario
      memoryUsage: JSON.stringify([...this.cache.entries()]).length,
      activeExecutions: this.executionSemaphore
    };
  }
}

export default new CompilerService();