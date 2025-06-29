import { useState } from 'react';
import compilerService from '../services/compilerService';

export const useReports = () => {
  const [isGeneratingReport, setIsGeneratingReport] = useState(false);
  const [reportProgress, setReportProgress] = useState('');

  const generateErrorReport = async (code) => {
    if (!code.trim() || code === '// Escribe tu código Vlang aquí...') {
      alert('⚠️ No hay código para analizar');
      return;
    }

    setIsGeneratingReport(true);
    setReportProgress('Analizando errores...');
    
    try {
      // CAMBIADO: Llamada directa al backend para obtener datos JSON
      const response = await fetch('http://localhost:8080/execute', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code })
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json();
      setReportProgress('¡Reporte de errores generado!');
      setTimeout(() => setReportProgress(''), 2000);
      
      return {
        success: data.success,
        output: data.output,
        errors: data.errors || [],
        symbols: data.symbols || []
      };
      
    } catch (error) {
      console.error('Error generando reporte de errores:', error);
      setReportProgress('');
      
      if (error.message.includes('fetch')) {
        throw new Error('No se pudo conectar con el backend. ¿Está corriendo el servidor?');
      } else {
        throw new Error(`Error generando reporte de errores: ${error.message}`);
      }
    } finally {
      setIsGeneratingReport(false);
    }
  };

  const generateSymbolReport = async (code) => {
    if (!code.trim() || code === '// Escribe tu código Vlang aquí...') {
      throw new Error('No hay código para analizar');
    }

    setIsGeneratingReport(true);
    setReportProgress('Extrayendo símbolos...');
    
    try {
      // CAMBIADO: Llamada directa al backend para obtener datos JSON
      const response = await fetch('http://localhost:8080/execute', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code })
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json();
      setReportProgress('¡Tabla de símbolos generada!');
      setTimeout(() => setReportProgress(''), 2000);
      
      return {
        success: data.success,
        symbols: data.symbols || []
      };
      
    } catch (error) {
      console.error('Error generando tabla de símbolos:', error);
      setReportProgress('');
      
      if (error.message.includes('fetch')) {
        throw new Error('No se pudo conectar con el backend. ¿Está corriendo el servidor?');
      } else {
        throw new Error(`Error generando tabla de símbolos: ${error.message}`);
      }
    } finally {
      setIsGeneratingReport(false);
    }
  };

  const generateCSTReport = async (code) => {
    if (!code.trim() || code === '// Escribe tu código Vlang aquí...') {
      alert('⚠️ No hay código para analizar');
      return;
    }

    setIsGeneratingReport(true);
    setReportProgress('Generando árbol CST...');
    
    try {
      const htmlContent = await compilerService.generateCSTReport(code);
      openHTMLInNewTab(htmlContent, '🌳 Árbol CST - V→ARM64');
      setReportProgress('¡Árbol CST generado!');
      
      // Limpiar mensaje después de 2 segundos
      setTimeout(() => setReportProgress(''), 2000);
    } catch (error) {
      console.error('Error generando árbol CST:', error);
      setReportProgress('');
      
      if (error.message.includes('fetch')) {
        alert('❌ Error de conexión: No se pudo conectar con el backend. ¿Está corriendo el servidor?');
      } else if (error.message.includes('CST')) {
        alert('❌ Error generando CST: El servicio ANTLR Lab puede estar temporalmente no disponible.');
      } else {
        alert(`❌ Error generando árbol CST: ${error.message}`);
      }
    } finally {
      setIsGeneratingReport(false);
    }
  };

  // Generar CST para componente React
  const generateCST = async (code) => {
    if (!code.trim() || code === '// Escribe tu código Vlang aquí...') {
      throw new Error('No hay código para analizar');
    }

    setIsGeneratingReport(true);
    setReportProgress('Generando árbol CST interactivo...');
    
    try {
      // Usar el endpoint directo para obtener solo el SVG
      const response = await fetch('http://localhost:8080/cst', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code })
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json();
      
      if (data.success) {
        setReportProgress('¡Árbol CST generado exitosamente!');
        setTimeout(() => setReportProgress(''), 2000);
        return data.htmlContent; // SVG content
      } else {
        throw new Error(data.error || 'Error generando CST');
      }
    } catch (error) {
      console.error('Error generando CST:', error);
      setReportProgress('');
      
      if (error.message.includes('fetch')) {
        throw new Error('No se pudo conectar con el backend. ¿Está corriendo el servidor?');
      } else if (error.message.includes('CST')) {
        throw new Error('El servicio ANTLR Lab puede estar temporalmente no disponible.');
      } else {
        throw new Error(`Error generando árbol CST: ${error.message}`);
      }
    } finally {
      setIsGeneratingReport(false);
    }
  };

  const generateAllReports = async (code) => {
    if (!code.trim() || code === '// Escribe tu código Vlang aquí...') {
      alert('⚠️ No hay código para analizar');
      return;
    }

    setIsGeneratingReport(true);
    setReportProgress('Generando todos los reportes...');
    
    try {
      // Generar todos los reportes en paralelo
      const [errorReport, symbolReport, cstReport] = await Promise.allSettled([
        compilerService.generateErrorReport(code),
        compilerService.generateSymbolReport(code),
        compilerService.generateCSTReport(code)
      ]);

      // Abrir cada reporte que se generó exitosamente
      if (errorReport.status === 'fulfilled') {
        openHTMLInNewTab(errorReport.value, '📋 Reporte de Errores');
      }
      
      if (symbolReport.status === 'fulfilled') {
        openHTMLInNewTab(symbolReport.value, '🔤 Tabla de Símbolos');
      }
      
      if (cstReport.status === 'fulfilled') {
        openHTMLInNewTab(cstReport.value, '🌳 Árbol CST');
      }

      setReportProgress('¡Todos los reportes generados!');
      setTimeout(() => setReportProgress(''), 2000);
      
    } catch (error) {
      console.error('Error generando reportes:', error);
      setReportProgress('');
      alert(`❌ Error generando reportes: ${error.message}`);
    } finally {
      setIsGeneratingReport(false);
    }
  };

  return {
    isGeneratingReport,
    reportProgress,
    generateErrorReport,
    generateSymbolReport,
    generateCSTReport,
    generateCST,
    generateAllReports
  };
};