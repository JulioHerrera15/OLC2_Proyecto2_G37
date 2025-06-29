import React, { useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { TransformWrapper, TransformComponent } from 'react-zoom-pan-pinch';
import { X, Download, Image, ZoomIn, ZoomOut, Home, Target, TreePine } from 'lucide-react';

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

const controlsVariants = {
  hidden: { x: -50, opacity: 0 },
  visible: {
    x: 0,
    opacity: 1,
    transition: {
      type: "spring",
      stiffness: 200,
      damping: 20,
      delay: 0.3,
      staggerChildren: 0.1
    }
  }
};

const controlButtonVariants = {
  hidden: { scale: 0, opacity: 0 },
  visible: {
    scale: 1,
    opacity: 1,
    transition: {
      type: "spring",
      stiffness: 300,
      damping: 15
    }
  },
  hover: {
    scale: 1.1,
    rotate: 5,
    transition: {
      type: "spring",
      stiffness: 400,
      damping: 10
    }
  },
  tap: {
    scale: 0.9,
    rotate: -5
  }
};

const svgContainerVariants = {
  hidden: { 
    opacity: 0, 
    scale: 0.8,
    filter: "blur(10px)"
  },
  visible: {
    opacity: 1,
    scale: 1,
    filter: "blur(0px)",
    transition: {
      duration: 0.8,
      delay: 0.4,
      ease: "easeOut"
    }
  }
};

const floatingVariants = {
  initial: { y: 0 },
  animate: {
    y: [-2, 2, -2],
    transition: {
      duration: 2,
      ease: "easeInOut"
    }
  }
};

const CSTViewer = ({ svgContent, isDarkMode, onClose }) => {
  // Clases dinámicas basadas en el tema
  const themeClasses = {
    textPrimary: isDarkMode ? 'text-white' : 'text-gray-900',
    textSecondary: isDarkMode ? 'text-gray-400' : 'text-gray-600',
    panel: isDarkMode 
      ? 'bg-black/40 backdrop-blur-3xl border-emerald-500/20' 
      : 'bg-white/95 backdrop-blur-3xl border-emerald-400/30',
  };

  const downloadSVG = () => {
    const blob = new Blob([svgContent], { type: 'image/svg+xml' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'cst-tree.svg';
    a.click();
    URL.revokeObjectURL(url);
  };

  const downloadPNG = async () => {
    const canvas = document.createElement('canvas');
    const ctx = canvas.getContext('2d');
    const img = new Image();
    
    const svgBlob = new Blob([svgContent], { type: 'image/svg+xml;charset=utf-8' });
    const url = URL.createObjectURL(svgBlob);
    
    img.onload = () => {
      canvas.width = img.width * 2;
      canvas.height = img.height * 2;
      ctx.scale(2, 2);
      ctx.drawImage(img, 0, 0);
      
      canvas.toBlob((blob) => {
        const pngUrl = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = pngUrl;
        a.download = 'cst-tree.png';
        a.click();
        URL.revokeObjectURL(pngUrl);
        URL.revokeObjectURL(url);
      });
    };
    
    img.src = url;
  };

  return (
    <AnimatePresence>
      <motion.div 
        className="fixed inset-0 bg-gradient-to-br from-black/60 via-gray-900/70 to-black/80 backdrop-blur-sm flex items-center justify-center z-50"
        variants={backdropVariants}
        initial="hidden"
        animate="visible"
        exit="exit"
      >
        <motion.div 
          className={`${themeClasses.panel} rounded-2xl shadow-2xl max-w-[95vw] max-h-[95vh] w-full mx-4 flex flex-col border transition-all duration-500`}
          variants={modalVariants}
          initial="hidden"
          animate="visible"
          exit="exit"
          style={{ perspective: 1000 }}
        >
          {/* Header */}
          <motion.div 
            className="flex items-center justify-between p-4 border-b bg-gradient-to-r from-green-500 to-emerald-600 text-white rounded-t-2xl"
            variants={headerVariants}
          >
            <div className="flex items-center space-x-3">
              <motion.div
                variants={floatingVariants}
                initial="initial"
                animate="animate"
                className="text-3xl"
              >
                🌳
              </motion.div>
              <motion.div
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.3, duration: 0.5 }}
              >
                <h2 className="text-xl font-bold">Árbol CST Interactivo</h2>
                <p className="text-green-100 text-sm">Generado con ANTLR Lab</p>
              </motion.div>
            </div>
            
            <motion.div 
              className="flex items-center space-x-2"
              initial={{ opacity: 0, x: 20 }}
              animate={{ opacity: 1, x: 0 }}
              transition={{ delay: 0.2, duration: 0.5 }}
            >
              <motion.button
                onClick={downloadSVG}
                className="bg-white/20 hover:bg-white/30 text-white px-3 py-2 rounded-lg flex items-center space-x-2 transition-all backdrop-blur-md border border-white/20 hover:border-white/40"
                title="Descargar SVG"
                variants={buttonVariants}
                initial="initial"
                whileHover="hover"
                whileTap="tap"
              >
                <Download className="w-4 h-4" />
                <span className="hidden sm:inline font-medium">SVG</span>
              </motion.button>
              
              <motion.button
                onClick={downloadPNG}
                className="bg-white/20 hover:bg-white/30 text-white px-3 py-2 rounded-lg flex items-center space-x-2 transition-all backdrop-blur-md border border-white/20 hover:border-white/40"
                title="Descargar PNG"
                variants={buttonVariants}
                initial="initial"
                whileHover="hover"
                whileTap="tap"
              >
                <Image className="w-4 h-4" />
                <span className="hidden sm:inline font-medium">PNG</span>
              </motion.button>
              
              <motion.button
                onClick={onClose}
                className="bg-red-500 hover:bg-red-600 text-white px-3 py-2 rounded-lg transition-all font-medium"
                title="Cerrar"
                variants={buttonVariants}
                initial="initial"
                whileHover="hover"
                whileTap="tap"
              >
                <X className="w-4 h-4" />
              </motion.button>
            </motion.div>
          </motion.div>

          {/* CST Container con zoom */}
          <motion.div 
            className={`flex-1 overflow-hidden relative transition-all duration-500 ${
              isDarkMode 
                ? 'bg-gradient-to-br from-gray-900/50 to-black/50' 
                : 'bg-gradient-to-br from-gray-50 to-white'
            }`}
            variants={contentVariants}
            initial="hidden"
            animate="visible"
          >
            <TransformWrapper
              initialScale={0.6}
              minScale={0.1}
              maxScale={5}
              centerOnInit={true}
              wheel={{ 
                step: 0.1,
                smoothStep: 0.005,
                activationKeys: []
              }}
              doubleClick={{ 
                mode: 'reset',
                step: 0.6,
                animationTime: 300
              }}
              panning={{
                velocityDisabled: false,
                lockAxisX: false,
                lockAxisY: false,
                activationKeys: []
              }}
              limitToBounds={false}
              centerZoomedOut={false}
              transformEnabled={true}
              pinchEnabled={true}
              alignmentAnimation={{
                disabled: false,
                sizeX: 0,
                sizeY: 0,
                animationTime: 200,
                velocityAligner: {
                  tolerance: 1.2,
                  time: 100
                }
              }}
            >
              {({ zoomIn, zoomOut, resetTransform, setTransform, instance }) => (
                <>
                  {/* Controls flotantes */}
                  <motion.div 
                    className={`absolute top-4 left-4 z-10 flex flex-col space-y-2 ${isDarkMode ? 'bg-black/90' : 'bg-white/90'} backdrop-blur-sm rounded-lg shadow-lg p-2 border ${isDarkMode ? 'border-gray-700/50' : 'border-gray-200'} transition-all duration-500`}
                    variants={controlsVariants}
                    initial="hidden"
                    animate="visible"
                  >
                    <motion.button
                      onClick={() => zoomIn(0.3)}
                      className="bg-blue-500 hover:bg-blue-600 text-white p-2 rounded text-sm font-bold transition-all"
                      title="Ampliar (Scroll ↑)"
                      variants={controlButtonVariants}
                      whileHover="hover"
                      whileTap="tap"
                    >
                      <ZoomIn className="w-4 h-4" />
                    </motion.button>
                    
                    <motion.button
                      onClick={() => zoomOut(0.3)}
                      className="bg-blue-500 hover:bg-blue-600 text-white p-2 rounded text-sm font-bold transition-all"
                      title="Reducir (Scroll ↓)"
                      variants={controlButtonVariants}
                      whileHover="hover"
                      whileTap="tap"
                    >
                      <ZoomOut className="w-4 h-4" />
                    </motion.button>
                    
                    <motion.button
                      onClick={() => {
                        resetTransform();
                        setTimeout(() => {
                          const wrapper = instance.wrapperComponent;
                          if (wrapper) {
                            const { clientWidth, clientHeight } = wrapper;
                            setTransform(
                              clientWidth / 2 - (clientWidth * 0.6) / 2,
                              clientHeight / 2 - (clientHeight * 0.6) / 2,
                              0.6
                            );
                          }
                        }, 100);
                      }}
                      className="bg-green-500 hover:bg-green-600 text-white p-2 rounded text-sm font-bold transition-all"
                      title="Restablecer (Doble click)"
                      variants={controlButtonVariants}
                      whileHover="hover"
                      whileTap="tap"
                    >
                      <Home className="w-4 h-4" />
                    </motion.button>
                    
                    <motion.button
                      onClick={() => {
                        const wrapper = instance.wrapperComponent;
                        const content = instance.contentComponent;
                        
                        if (wrapper && content) {
                          const { clientWidth: wrapperWidth, clientHeight: wrapperHeight } = wrapper;
                          const { clientWidth: contentWidth, clientHeight: contentHeight } = content;
                          const currentScale = instance.transformState.scale;
                          
                          const x = (wrapperWidth - contentWidth * currentScale) / 2;
                          const y = (wrapperHeight - contentHeight * currentScale) / 2;
                          
                          setTransform(x, y, currentScale, 300);
                        }
                      }}
                      className="bg-purple-500 hover:bg-purple-600 text-white p-2 rounded text-sm font-bold transition-all"
                      title="Centrar vista"
                      variants={controlButtonVariants}
                      whileHover="hover"
                      whileTap="tap"
                    >
                      <Target className="w-4 h-4" />
                    </motion.button>
                  </motion.div>

                  {/* SVG Container */}
                  <TransformComponent
                    wrapperStyle={{
                      width: '100%',
                      height: '70vh',
                      cursor: 'grab',
                    }}
                    contentStyle={{
                      width: '100%',
                      height: '100%',
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      transform: 'translateZ(0)',
                      backfaceVisibility: 'hidden',
                      perspective: '1000px',
                    }}
                  >
                    <motion.div
                      className={`select-none pointer-events-auto rounded-lg shadow-xl p-4 transition-all duration-300 ${
                        isDarkMode 
                          ? 'bg-white border-2 border-gray-300' 
                          : 'bg-white border border-gray-200'
                      }`}
                      style={{
                        willChange: 'transform',
                        transform: 'translateZ(0)',
                        minWidth: 'fit-content',
                        minHeight: 'fit-content',
                      }}
                      variants={svgContainerVariants}
                      initial="hidden"
                      animate="visible"
                      whileHover={{ 
                        boxShadow: "0 20px 50px rgba(0, 0, 0, 0.15)",
                        transition: { duration: 0.3 }
                      }}
                      dangerouslySetInnerHTML={{ __html: svgContent }}
                    />
                  </TransformComponent>
                </>
              )}
            </TransformWrapper>
          </motion.div>

          {/* Footer con instrucciones */}
          <motion.div 
            className={`p-4 border-t ${isDarkMode ? 'bg-gradient-to-r from-gray-800/50 to-gray-700/50 border-gray-700/50 text-gray-300' : 'bg-gradient-to-r from-gray-50 to-gray-100 border-gray-200 text-gray-600'} text-sm rounded-b-2xl transition-all duration-500`}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.6, duration: 0.5 }}
          >
            <div className="flex flex-wrap gap-4 justify-center text-center">
              <motion.span 
                className="flex items-center gap-1"
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.7 }}
              >
                <span>🖱️</span> <strong>Arrastrar:</strong> Mover árbol
              </motion.span>
              <motion.span 
                className="flex items-center gap-1"
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.8 }}
              >
                <span>🔄</span> <strong>Scroll:</strong> Zoom in/out
              </motion.span>
              <motion.span 
                className="flex items-center gap-1"
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.9 }}
              >
                <span>👆</span> <strong>Doble click:</strong> Restablecer
              </motion.span>
            </div>
          </motion.div>
        </motion.div>
      </motion.div>
    </AnimatePresence>
  );
};

export default CSTViewer;