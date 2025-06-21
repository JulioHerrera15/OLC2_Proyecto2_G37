import React from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { X, Bug, Hash, AlertTriangle, CheckCircle, Calendar } from 'lucide-react';

// Variantes de animación sofisticadas
const backdropVariants = {
  hidden: { opacity: 0 },
  visible: {
    opacity: 1,
    transition: {
      duration: 0.4,
      ease: "easeOut"
    }
  },
  exit: {
    opacity: 0,
    transition: {
      duration: 0.3,
      ease: "easeIn"
    }
  }
};

const modalVariants = {
  hidden: { 
    opacity: 0, 
    scale: 0.8, 
    y: 100,
    rotateX: -15
  },
  visible: {
    opacity: 1,
    scale: 1,
    y: 0,
    rotateX: 0,
    transition: {
      type: "spring",
      stiffness: 200,
      damping: 25,
      duration: 0.6,
      staggerChildren: 0.1
    }
  },
  exit: {
    opacity: 0,
    scale: 0.8,
    y: 100,
    rotateX: 15,
    transition: {
      duration: 0.4,
      ease: "easeIn"
    }
  }
};

const headerVariants = {
  hidden: { y: -50, opacity: 0 },
  visible: {
    y: 0,
    opacity: 1,
    transition: {
      type: "spring",
      stiffness: 300,
      damping: 30,
      delay: 0.1
    }
  }
};

const contentVariants = {
  hidden: { opacity: 0, y: 30 },
  visible: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 0.5,
      delay: 0.2,
      staggerChildren: 0.1
    }
  }
};

const cardVariants = {
  hidden: { opacity: 0, x: -30, scale: 0.95 },
  visible: {
    opacity: 1,
    x: 0,
    scale: 1,
    transition: {
      type: "spring",
      stiffness: 200,
      damping: 20
    }
  }
};

const errorItemVariants = {
  hidden: { opacity: 0, x: -20, scale: 0.9 },
  visible: {
    opacity: 1,
    x: 0,
    scale: 1,
    transition: {
      type: "spring",
      stiffness: 150,
      damping: 15
    }
  }
};

const tableRowVariants = {
  hidden: { opacity: 0, y: 20 },
  visible: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 0.3,
      ease: "easeOut"
    }
  }
};

const buttonVariants = {
  initial: { scale: 1 },
  hover: {
    scale: 1.05,
    transition: {
      type: "spring",
      stiffness: 400,
      damping: 10
    }
  },
  tap: {
    scale: 0.95,
    transition: {
      duration: 0.1
    }
  }
};

const iconVariants = {
  initial: { rotate: 0 },
  hover: {
    rotate: 360,
    transition: {
      duration: 0.6,
      ease: "easeInOut"
    }
  }
};

const floatingVariants = {
  initial: { y: 0 },
  animate: {
    y: [-2, 2, -2],
    transition: {
      duration: 2,
      repeat: Infinity,
      ease: "easeInOut"
    }
  }
};

const ReportViewer = ({ reportData, reportType, isDarkMode, onClose }) => {
  // Clases dinámicas basadas en el tema
  const themeClasses = {
    background: isDarkMode 
      ? 'bg-gradient-to-br from-slate-900 via-gray-900 to-slate-800' 
      : 'bg-gradient-to-br from-gray-50 via-blue-50 to-indigo-50',
    textPrimary: isDarkMode ? 'text-white' : 'text-gray-900',
    textSecondary: isDarkMode ? 'text-gray-400' : 'text-gray-600',
    textAccent: isDarkMode ? 'text-emerald-400' : 'text-emerald-600',
    panel: isDarkMode 
      ? 'bg-black/40 backdrop-blur-3xl border-emerald-500/20' 
      : 'bg-white/40 backdrop-blur-3xl border-emerald-400/30',
    panelHeader: isDarkMode 
      ? 'bg-gradient-to-r from-black/30 to-emerald-500/5 border-emerald-500/20'
      : 'bg-gradient-to-r from-white/50 to-emerald-400/10 border-emerald-400/30',
  };

  const getReportConfig = () => {
    switch (reportType) {
      case 'errors':
        return {
          title: 'Reporte de Errores y Warnings',
          emoji: '🐛',
          icon: Bug,
          gradient: isDarkMode 
            ? 'from-red-500 to-red-600' 
            : 'from-red-500 to-red-600',
          bgGradient: isDarkMode 
            ? 'from-red-500/10 to-red-600/10' 
            : 'from-red-50 to-red-100',
          color: 'red'
        };
      case 'symbols':
        return {
          title: 'Tabla de Símbolos',
          emoji: '🔤',
          icon: Hash,
          gradient: isDarkMode 
            ? 'from-blue-500 to-blue-600' 
            : 'from-blue-500 to-blue-600',
          bgGradient: isDarkMode 
            ? 'from-blue-500/10 to-blue-600/10' 
            : 'from-blue-50 to-blue-100',
          color: 'blue'
        };
      default:
        return {
          title: 'Reporte',
          emoji: '📋',
          icon: Bug,
          gradient: isDarkMode 
            ? 'from-gray-500 to-gray-600' 
            : 'from-gray-500 to-gray-600',
          bgGradient: isDarkMode 
            ? 'from-gray-500/10 to-gray-600/10' 
            : 'from-gray-50 to-gray-100',
          color: 'gray'
        };
    }
  };

  const config = getReportConfig();
  const IconComponent = config.icon;

  const renderErrorReport = () => {
    const hasErrors = reportData.errors && reportData.errors.length > 0;
    const statusInfo = hasErrors 
      ? { 
          bg: isDarkMode ? 'bg-red-500/10 border-red-500/20' : 'bg-red-100 border-red-200', 
          text: isDarkMode ? 'text-red-400' : 'text-red-800', 
          emoji: '❌', 
          status: 'Errores encontrados' 
        }
      : { 
          bg: isDarkMode ? 'bg-green-500/10 border-green-500/20' : 'bg-green-100 border-green-200', 
          text: isDarkMode ? 'text-green-400' : 'text-green-800', 
          emoji: '✅', 
          status: 'Sin errores' 
        };

    return (
      <motion.div 
        className="space-y-6"
        variants={contentVariants}
        initial="hidden"
        animate="visible"
      >
        {/* Status Card */}
        <motion.div 
          className={`${statusInfo.bg} ${statusInfo.text} p-4 rounded-xl border-l-4 ${hasErrors ? 'border-l-red-500' : 'border-l-green-500'} border transition-all duration-500`}
          variants={cardVariants}
          whileHover={{ scale: 1.02, transition: { duration: 0.2 } }}
        >
          <div className="flex items-center gap-3">
            <motion.span 
              className="text-2xl"
              variants={floatingVariants}
              initial="initial"
              animate="animate"
            >
              {statusInfo.emoji}
            </motion.span>
            <motion.div
              initial={{ opacity: 0, x: -20 }}
              animate={{ opacity: 1, x: 0 }}
              transition={{ delay: 0.3, duration: 0.5 }}
            >
              <p className="font-semibold">Estado: {statusInfo.status}</p>
              <p className="text-sm opacity-80">
                {hasErrors ? `${reportData.errors.length} error(es) encontrado(s)` : 'Código analizado correctamente'}
              </p>
            </motion.div>
          </div>
        </motion.div>

        {/* Errors Detail */}
        <AnimatePresence>
          {hasErrors && (
            <motion.div 
              className={`${isDarkMode ? 'bg-red-500/5 border-red-500/20' : 'bg-red-50 border-red-200'} rounded-xl p-4 border transition-all duration-500`}
              variants={cardVariants}
              initial="hidden"
              animate="visible"
              exit={{ opacity: 0, scale: 0.95, transition: { duration: 0.3 } }}
            >
              <motion.h3 
                className={`font-semibold ${isDarkMode ? 'text-red-400' : 'text-red-800'} mb-4 flex items-center gap-2 transition-all duration-500`}
                initial={{ opacity: 0, y: -10 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.2 }}
              >
                <motion.div variants={iconVariants} whileHover="hover">
                  <AlertTriangle className="w-5 h-5" />
                </motion.div>
                Detalles de Errores
              </motion.h3>
              <motion.div 
                className="space-y-3"
                variants={contentVariants}
              >
                {reportData.errors.map((error, index) => (
                  <motion.div 
                    key={index} 
                    className={`${isDarkMode ? 'bg-black/20 border-red-500/30' : 'bg-white border-red-200'} p-4 rounded-lg border shadow-sm transition-all duration-500`}
                    variants={errorItemVariants}
                    whileHover={{ 
                      scale: 1.02, 
                      boxShadow: isDarkMode 
                        ? "0 8px 32px rgba(239, 68, 68, 0.1)" 
                        : "0 8px 32px rgba(239, 68, 68, 0.15)",
                      transition: { duration: 0.2 }
                    }}
                  >
                    <div className="flex items-start gap-3">
                      <motion.div 
                        className={`${isDarkMode ? 'bg-red-500/20 text-red-400' : 'bg-red-100 text-red-600'} p-2 rounded-lg transition-all duration-500`}
                        initial={{ scale: 0 }}
                        animate={{ scale: 1 }}
                        transition={{ 
                          delay: 0.1 * index, 
                          type: "spring", 
                          stiffness: 200, 
                          damping: 15 
                        }}
                      >
                        <span className="font-bold text-sm">#{index + 1}</span>
                      </motion.div>
                      <motion.div 
                        className="flex-1"
                        initial={{ opacity: 0, x: -10 }}
                        animate={{ opacity: 1, x: 0 }}
                        transition={{ delay: 0.2 + 0.1 * index }}
                      >
                        <div className="flex items-center gap-2 mb-2">
                          <motion.span 
                            className="bg-red-500 text-white px-2 py-1 rounded text-xs font-medium"
                            initial={{ opacity: 0, scale: 0.8 }}
                            animate={{ opacity: 1, scale: 1 }}
                            transition={{ delay: 0.3 + 0.1 * index }}
                          >
                            {error.type}
                          </motion.span>
                          <span className={`${themeClasses.textSecondary} text-sm transition-all duration-500`}>
                            Línea {error.line}, Columna {error.column}
                          </span>
                        </div>
                        <p className={`${themeClasses.textPrimary} font-medium transition-all duration-500`}>
                          {error.message}
                        </p>
                        {error.severity && (
                          <motion.span 
                            className={`inline-block mt-2 px-2 py-1 rounded text-xs font-medium ${
                              error.severity === 'high' 
                                ? (isDarkMode ? 'bg-red-500/20 text-red-400' : 'bg-red-100 text-red-800')
                                : error.severity === 'medium' 
                                ? (isDarkMode ? 'bg-yellow-500/20 text-yellow-400' : 'bg-yellow-100 text-yellow-800')
                                : (isDarkMode ? 'bg-blue-500/20 text-blue-400' : 'bg-blue-100 text-blue-800')
                            } transition-all duration-500`}
                            initial={{ opacity: 0, y: 10 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ delay: 0.4 + 0.1 * index }}
                          >
                            Severidad: {error.severity}
                          </motion.span>
                        )}
                      </motion.div>
                    </div>
                  </motion.div>
                ))}
              </motion.div>
            </motion.div>
          )}
        </AnimatePresence>
      </motion.div>
    );
  };

  const renderSymbolTable = () => {
    const hasSymbols = reportData.symbols && reportData.symbols.length > 0;

    const getSymbolBadge = (type) => {
      const baseClasses = 'transition-all duration-500';
      switch (type.toLowerCase()) {
        case 'variable':
          return isDarkMode ? `${baseClasses} bg-blue-500/20 text-blue-400` : `${baseClasses} bg-blue-100 text-blue-800`;
        case 'function':
          return isDarkMode ? `${baseClasses} bg-green-500/20 text-green-400` : `${baseClasses} bg-green-100 text-green-800`;
        case 'struct':
          return isDarkMode ? `${baseClasses} bg-yellow-500/20 text-yellow-400` : `${baseClasses} bg-yellow-100 text-yellow-800`;
        case 'class':
          return isDarkMode ? `${baseClasses} bg-purple-500/20 text-purple-400` : `${baseClasses} bg-purple-100 text-purple-800`;
        default:
          return isDarkMode ? `${baseClasses} bg-gray-500/20 text-gray-400` : `${baseClasses} bg-gray-100 text-gray-800`;
      }
    };

    return (
      <motion.div 
        className="space-y-6"
        variants={contentVariants}
        initial="hidden"
        animate="visible"
      >
        {/* Summary Card */}
        <motion.div 
          className={`${isDarkMode ? 'bg-blue-500/10 text-blue-400 border-blue-500/20' : 'bg-blue-50 text-blue-800 border-blue-200'} p-4 rounded-xl border-l-4 border-l-blue-500 border transition-all duration-500`}
          variants={cardVariants}
          whileHover={{ scale: 1.02, transition: { duration: 0.2 } }}
        >
          <div className="flex items-center gap-3">
            <motion.span 
              className="text-2xl"
              variants={floatingVariants}
              initial="initial"
              animate="animate"
            >
              📊
            </motion.span>
            <motion.div
              initial={{ opacity: 0, x: -20 }}
              animate={{ opacity: 1, x: 0 }}
              transition={{ delay: 0.3, duration: 0.5 }}
            >
              <p className="font-semibold">Resumen de Símbolos</p>
              <p className="text-sm opacity-80">
                {hasSymbols ? `${reportData.symbols.length} símbolo(s) encontrado(s)` : 'No se encontraron símbolos'}
              </p>
            </motion.div>
          </div>
        </motion.div>

        {/* Symbols Table */}
        <AnimatePresence>
          {hasSymbols ? (
            <motion.div 
              className={`${isDarkMode ? 'bg-black/20 border-gray-700/50' : 'bg-white border-gray-200'} rounded-xl border overflow-hidden transition-all duration-500`}
              variants={cardVariants}
              initial="hidden"
              animate="visible"
              exit={{ opacity: 0, scale: 0.95, transition: { duration: 0.3 } }}
            >
              <div className="overflow-x-auto">
                <table className="w-full">
                  <motion.thead 
                    className={`${isDarkMode ? 'bg-gray-800/50' : 'bg-gray-50'} transition-all duration-500`}
                    initial={{ opacity: 0, y: -20 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ delay: 0.2 }}
                  >
                    <tr>
                      {['Nombre', 'Tipo', 'Ámbito', 'Línea', 'Columna'].map((header, index) => (
                        <motion.th 
                          key={header}
                          className={`px-6 py-4 text-left text-xs font-medium ${themeClasses.textSecondary} uppercase tracking-wider transition-all duration-500`}
                          initial={{ opacity: 0, y: -10 }}
                          animate={{ opacity: 1, y: 0 }}
                          transition={{ delay: 0.3 + index * 0.1 }}
                        >
                          {header}
                        </motion.th>
                      ))}
                    </tr>
                  </motion.thead>
                  <motion.tbody 
                    className={`${isDarkMode ? 'bg-black/10 divide-gray-700/50' : 'bg-white divide-gray-200'} divide-y transition-all duration-500`}
                    variants={contentVariants}
                  >
                    {reportData.symbols.map((symbol, index) => (
                      <motion.tr 
                        key={index} 
                        className={`${isDarkMode ? 'hover:bg-gray-800/30' : 'hover:bg-gray-50'} transition-colors duration-300`}
                        variants={tableRowVariants}
                        whileHover={{ 
                          backgroundColor: isDarkMode ? 'rgba(31, 41, 55, 0.3)' : 'rgba(249, 250, 251, 1)',
                          scale: 1.01,
                          transition: { duration: 0.2 }
                        }}
                      >
                        <td className="px-6 py-4 whitespace-nowrap">
                          <motion.div 
                            className="flex items-center"
                            initial={{ opacity: 0, x: -10 }}
                            animate={{ opacity: 1, x: 0 }}
                            transition={{ delay: 0.1 * index }}
                          >
                            <span className={`text-sm font-medium ${themeClasses.textPrimary} font-mono transition-all duration-500`}>
                              {symbol.name}
                            </span>
                          </motion.div>
                        </td>
                        <td className="px-6 py-4 whitespace-nowrap">
                          <motion.span 
                            className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${getSymbolBadge(symbol.type)}`}
                            initial={{ opacity: 0, scale: 0.8 }}
                            animate={{ opacity: 1, scale: 1 }}
                            transition={{ delay: 0.2 + 0.1 * index }}
                            whileHover={{ scale: 1.1 }}
                          >
                            {symbol.type}
                          </motion.span>
                        </td>
                        <td className={`px-6 py-4 whitespace-nowrap text-sm ${themeClasses.textPrimary} transition-all duration-500`}>
                          {symbol.scope || 'global'}
                        </td>
                        <td className={`px-6 py-4 whitespace-nowrap text-sm ${themeClasses.textSecondary} font-mono transition-all duration-500`}>
                          {symbol.line}
                        </td>
                        <td className={`px-6 py-4 whitespace-nowrap text-sm ${themeClasses.textSecondary} font-mono transition-all duration-500`}>
                          {symbol.column}
                        </td>
                      </motion.tr>
                    ))}
                  </motion.tbody>
                </table>
              </div>
            </motion.div>
          ) : (
            <motion.div 
              className={`${isDarkMode ? 'bg-gray-800/30 border-gray-700/50' : 'bg-gray-50 border-gray-200'} rounded-xl p-8 text-center border transition-all duration-500`}
              variants={cardVariants}
              initial="hidden"
              animate="visible"
            >
              <motion.div 
                className={`${themeClasses.textSecondary} mb-4 transition-all duration-500`}
                variants={floatingVariants}
                initial="initial"
                animate="animate"
              >
                <Hash className="w-16 h-16 mx-auto opacity-50" />
              </motion.div>
              <motion.p 
                className={`${themeClasses.textPrimary} font-medium transition-all duration-500`}
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.3 }}
              >
                No se encontraron símbolos
              </motion.p>
              <motion.p 
                className={`${themeClasses.textSecondary} text-sm mt-2 transition-all duration-500`}
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.4 }}
              >
                El código no contiene variables, funciones o estructuras detectables.
              </motion.p>
            </motion.div>
          )}
        </AnimatePresence>
      </motion.div>
    );
  };

  return (
    <AnimatePresence>
      <motion.div 
        className="fixed inset-0 bg-gradient-to-br from-black/60 via-gray-900/70 to-black/80 backdrop-blur-sm flex items-center justify-center z-50 p-4"
        variants={backdropVariants}
        initial="hidden"
        animate="visible"
        exit="exit"
      >
        <motion.div 
          className={`${isDarkMode ? 'bg-black/40 backdrop-blur-3xl border-emerald-500/20' : 'bg-white/95 backdrop-blur-3xl border-emerald-400/30'} rounded-2xl shadow-2xl max-w-6xl w-full max-h-[90vh] flex flex-col overflow-hidden border transition-all duration-500`}
          variants={modalVariants}
          initial="hidden"
          animate="visible"
          exit="exit"
          style={{ perspective: 1000 }}
        >
          {/* Header */}
          <motion.div 
            className={`bg-gradient-to-r ${config.gradient} text-white p-6 flex items-center justify-between transition-all duration-500`}
            variants={headerVariants}
          >
            <div className="flex items-center gap-4">
              <motion.div 
                className="bg-white/20 p-3 rounded-xl backdrop-blur-md"
                whileHover={{ scale: 1.1, rotate: 5 }}
                transition={{ type: "spring", stiffness: 300, damping: 20 }}
              >
                <motion.span 
                  className="text-3xl"
                  variants={floatingVariants}
                  initial="initial"
                  animate="animate"
                >
                  {config.emoji}
                </motion.span>
              </motion.div>
              <motion.div
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.3, duration: 0.5 }}
              >
                <h2 className="text-2xl font-bold">{config.title}</h2>
                <p className="text-white/80 text-sm">
                  Análisis generado el {new Date().toLocaleDateString('es-ES', {
                    day: '2-digit',
                    month: '2-digit',
                    year: 'numeric',
                    hour: '2-digit',
                    minute: '2-digit'
                  })}
                </p>
              </motion.div>
            </div>
            <motion.button
              onClick={onClose}
              className="bg-white/20 hover:bg-white/30 text-white p-3 rounded-xl transition-all backdrop-blur-md"
              title="Cerrar reporte"
              variants={buttonVariants}
              initial="initial"
              whileHover="hover"
              whileTap="tap"
            >
              <motion.div variants={iconVariants} whileHover="hover">
                <X className="w-6 h-6" />
              </motion.div>
            </motion.button>
          </motion.div>

          {/* Content */}
          <motion.div 
            className="flex-1 overflow-y-auto p-6"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.4, duration: 0.5 }}
          >
            {reportType === 'errors' ? renderErrorReport() : renderSymbolTable()}
          </motion.div>

          {/* Footer */}
          <motion.div 
            className={`bg-gradient-to-r ${config.bgGradient} px-6 py-4 border-t ${isDarkMode ? 'border-gray-700/50' : 'border-gray-200'} transition-all duration-500`}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.5, duration: 0.5 }}
          >
            <div className={`flex items-center justify-between text-sm ${themeClasses.textSecondary} transition-all duration-500`}>
              <motion.div 
                className="flex items-center gap-4"
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.6 }}
              >
                <span className="flex items-center gap-2">
                  <motion.div variants={iconVariants} whileHover="hover">
                    <Calendar className="w-4 h-4" />
                  </motion.div>
                  Compilador V→ARM64 v2.0
                </span>
              </motion.div>
              <motion.div 
                className="flex items-center gap-2"
                initial={{ opacity: 0, x: 20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.7 }}
              >
                <motion.div
                  initial={{ scale: 0 }}
                  animate={{ scale: 1 }}
                  transition={{ delay: 0.8, type: "spring", stiffness: 200 }}
                >
                  <CheckCircle className="w-4 h-4 text-green-500" />
                </motion.div>
                <span>Reporte generado exitosamente</span>
              </motion.div>
            </div>
          </motion.div>
        </motion.div>
      </motion.div>
    </AnimatePresence>
  );
};

export default ReportViewer;