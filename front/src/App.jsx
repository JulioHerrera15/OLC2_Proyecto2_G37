import { useState, useRef, useEffect } from 'react'
import { motion, AnimatePresence, useAnimation } from 'framer-motion'
import { 
  FileText, 
  FolderOpen, 
  Save, 
  Bug, 
  Hash, 
  Play, 
  Trash2, 
  Circle,
  Cpu,
  Code2,
  Terminal,
  Clock,
  CheckCircle,
  GitBranch,
  Moon,
  Sun,
  Loader2
} from 'lucide-react'

// Importar los nuevos servicios
import compilerService from './services/compilerService'
import { useReports } from './hooks/useReports'
import CSTViewer from './components/CSTViewer'
import ReportViewer from './components/ReportViewer'

// Variantes de animación
const containerVariants = {
  hidden: { opacity: 0 },
  visible: {
    opacity: 1,
    transition: {
      duration: 0.6,
      staggerChildren: 0.1
    }
  }
}

const headerVariants = {
  hidden: { y: -100, opacity: 0 },
  visible: {
    y: 0,
    opacity: 1,
    transition: {
      type: "spring",
      stiffness: 100,
      damping: 15,
      duration: 0.6
    }
  }
}

const panelVariants = {
  hidden: { scale: 0.8, opacity: 0, y: 50 },
  visible: {
    scale: 1,
    opacity: 1,
    y: 0,
    transition: {
      type: "spring",
      stiffness: 120,
      damping: 20,
      duration: 0.8
    }
  }
}

const footerVariants = {
  hidden: { y: 100, opacity: 0 },
  visible: {
    y: 0,
    opacity: 1,
    transition: {
      type: "spring",
      stiffness: 100,
      damping: 15,
      delay: 0.3,
      duration: 0.6
    }
  }
}

const menuVariants = {
  hidden: { opacity: 0, scale: 0.9, y: -20 },
  visible: {
    opacity: 1,
    scale: 1,
    y: 0,
    transition: {
      type: "spring",
      stiffness: 200,
      damping: 20,
      duration: 0.4
    }
  },
  exit: {
    opacity: 0,
    scale: 0.9,
    y: -20,
    transition: {
      duration: 0.2
    }
  }
}

const menuItemVariants = {
  hidden: { opacity: 0, x: -20 },
  visible: {
    opacity: 1,
    x: 0,
    transition: {
      type: "spring",
      stiffness: 200,
      damping: 20
    }
  }
}

const buttonHoverVariants = {
  hover: {
    scale: 1.05,
    y: -2,
    transition: {
      type: "spring",
      stiffness: 300,
      damping: 10
    }
  },
  tap: {
    scale: 0.95
  }
}

const floatingVariants = {
  initial: { y: 0 },
  animate: {
    y: [-5, 5, -5],
    transition: {
      duration: 3,
      repeat: Infinity,
      ease: "easeInOut"
    }
  }
}

function App() {
  // ...existing state...
  const [code, setCode] = useState(`// Escribe tu código Vlang aquí...`);
  const [output, setOutput] = useState('Salida...');
  const [activeMenu, setActiveMenu] = useState(null)
  const [errors, setErrors] = useState([])
  const [symbols, setSymbols] = useState([])
  const [isCompiling, setIsCompiling] = useState(false)
  const [isDarkMode, setIsDarkMode] = useState(true)
  const [compilationStats, setCompilationStats] = useState({
    time: '0ms',
    lines: 0,
    optimized: true,
    size: '0 KB'
  })

  const [currentFile, setCurrentFile] = useState('')
  const [hasUnsavedChanges, setHasUnsavedChanges] = useState(false)
  const fileInputRef = useRef(null)

  const [showCST, setShowCST] = useState(false)
  const [cstContent, setCstContent] = useState('')
  const [showReport, setShowReport] = useState(false)
  const [reportData, setReportData] = useState(null)
  const [reportType, setReportType] = useState('')

  const reports = useReports()
  const menuContainerRef = useRef(null)

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (activeMenu && menuContainerRef.current && !menuContainerRef.current.contains(event.target)) {
        setActiveMenu(null)
      }
    }

    const handleKeyDown = (event) => {
      if (event.key === 'Escape' && activeMenu) {
        setActiveMenu(null)
      }
    }

    if (activeMenu) {
      document.addEventListener('mousedown', handleClickOutside)
      document.addEventListener('keydown', handleKeyDown)
    }

    return () => {
      document.removeEventListener('mousedown', handleClickOutside)
      document.removeEventListener('keydown', handleKeyDown)
    }
  }, [activeMenu])

  // Función para manejar CST interactivo
  const handleCSTInteractive = async () => {
    if (!code.trim() || code === '// Escribe tu código Vlang aquí...') {
      alert('⚠️ No hay código para analizar');
      return;
    }

    try {
      setActiveMenu(null) // Cerrar menú
      const svgContent = await reports.generateCST(code);
      setCstContent(svgContent);
      setShowCST(true);
    } catch (error) {
      console.error('Error generando CST:', error);
      alert(`❌ Error generando el árbol CST: ${error.message}`);
    }
  };

  // Invalidar cache cuando cambie el código
  useEffect(() => {
    if (code !== `// Escribe tu código Vlang aquí...`) {
      setHasUnsavedChanges(true)
      updateStats()
      
      // Invalidar cache del servicio
      const hash = compilerService.generateCodeHash(code)
      compilerService.invalidateCache(hash)
    }
  }, [code])

  // Función para alternar tema
  const toggleTheme = () => {
    setIsDarkMode(!isDarkMode)
  }

  // Clases dinámicas basadas en el tema
  const themeClasses = {
    // Fondo principal
    background: isDarkMode 
      ? 'bg-gradient-to-br from-slate-900 via-gray-900 to-slate-800' 
      : 'bg-gradient-to-br from-gray-50 via-blue-50 to-indigo-50',
    
    // Header
    header: isDarkMode 
      ? 'bg-black/30 backdrop-blur-3xl border-emerald-500/20' 
      : 'bg-white/60 backdrop-blur-3xl border-emerald-400/30',
    
    // Texto principal
    textPrimary: isDarkMode ? 'text-white' : 'text-gray-900',
    textSecondary: isDarkMode ? 'text-gray-400' : 'text-gray-600',
    textAccent: isDarkMode ? 'text-emerald-400' : 'text-emerald-600',
    
    // Paneles principales
    panel: isDarkMode 
      ? 'bg-black/40 backdrop-blur-3xl border-emerald-500/20' 
      : 'bg-white/40 backdrop-blur-3xl border-emerald-400/30',
    
    // Panel headers
    panelHeader: isDarkMode 
      ? 'bg-gradient-to-r from-black/30 to-emerald-500/5 border-emerald-500/20'
      : 'bg-gradient-to-r from-white/50 to-emerald-400/10 border-emerald-400/30',
    
    // Editor de código
    editor: isDarkMode 
      ? 'text-gray-100 placeholder-gray-500'
      : 'text-gray-900 placeholder-gray-400',
    
    // Consola
    console: isDarkMode ? 'text-gray-200' : 'text-gray-800',
    
    // Footer
    footer: isDarkMode 
      ? 'bg-black/20 backdrop-blur-2xl border-gray-700/30'
      : 'bg-white/40 backdrop-blur-2xl border-gray-300/40'
  }

  // Función para crear nuevo proyecto
  function handleNew() {
    setCode(`// Escribe tu código Vlang aquí...`)
    setOutput('Salida...')
    setActiveMenu(null)
    setErrors([])
    setSymbols([])
    setCurrentFile('')
    setHasUnsavedChanges(false)
    compilerService.clearCache()
    updateStats()
  }

  // Función para abrir archivo
  function handleOpenFile() {
    fileInputRef.current?.click()
    setActiveMenu(null)
  }

  // Función para manejar la lectura de archivo
  function handleFileRead(event) {
    const file = event.target.files[0]
    if (!file) return

    const reader = new FileReader()
    reader.onload = (e) => {
      setCode(e.target.result)
      setCurrentFile(file.name)
      setHasUnsavedChanges(false)
      setOutput(`📁 Archivo "${file.name}" cargado exitosamente\n✅ Listo para compilación`)
      setActiveMenu(null)
    }
    reader.onerror = () => {
      setOutput(`❌ Error leyendo archivo: ${reader.error}`)
    }
    reader.readAsText(file)
  }

  // Función para guardar archivo
  function handleSaveFile() {
    const filename = currentFile || 'programa.vch'
    const blob = new Blob([code], { type: 'text/plain' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = filename
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    URL.revokeObjectURL(url)
    setHasUnsavedChanges(false)
    setActiveMenu(null)
  }

  // Función de compilación actualizada
  async function handleCompile() {
    if (!code.trim() || code === `// Escribe tu código Vlang aquí...`) {
      setOutput('❌ No hay código para compilar')
      return
    }

    setIsCompiling(true)
    setActiveMenu(null)
    
    const startTime = Date.now()
    
    try {
      setOutput('🔄 Compilando código...\n⏳ Por favor espera...')
      
      const result = await compilerService.executeCode(code)
      const endTime = Date.now()
      const compileTime = endTime - startTime
      
      if (result.success) {
        setOutput(result.output || '✅ Compilación exitosa')
      } else {
        setOutput(`❌ Error de compilación:\n${result.output || 'Error desconocido'}`)
      }
      
      setErrors(result.errors || [])
      setSymbols(result.symbols || [])
      
      setCompilationStats({
        time: `${compileTime}ms`,
        lines: code.split('\n').length,
        optimized: result.optimized || true,
        size: `${(code.length / 1024).toFixed(1)} KB`
      })
      
    } catch (error) {
      const endTime = Date.now()
      const compileTime = endTime - startTime
      
      setOutput(`❌ Error de conexión:\n${error.message}\n\n💡 Asegúrate de que el backend esté ejecutándose en puerto 8080`)
      setErrors([{
        type: 'ConnectionError',
        message: error.message,
        line: 1,
        column: 1,
        severity: 'high'
      }])
      
      setCompilationStats(prev => ({
        ...prev,
        time: `${compileTime}ms`
      }))
    } finally {
      setIsCompiling(false)
      updateStats()
    }
  }

  const handleErrorReport = async () => {
    if (!code.trim() || code === '// Escribe tu código Vlang aquí...') {
      alert('⚠️ No hay código para analizar');
      return;
    }

    try {
      setActiveMenu(null);
      const data = await reports.generateErrorReport(code);
      setReportData(data);
      setReportType('errors');
      setShowReport(true);
    } catch (error) {
      console.error('Error generando reporte:', error);
      alert(`❌ ${error.message}`);
    }
  };

  const handleSymbolReport = async () => {
    if (!code.trim() || code === '// Escribe tu código Vlang aquí...') {
      alert('⚠️ No hay código para analizar');
      return;
    }

    try {
      setActiveMenu(null);
      const data = await reports.generateSymbolReport(code);
      setReportData(data);
      setReportType('symbols');
      setShowReport(true);
    } catch (error) {
      console.error('Error generando reporte:', error);
      alert(`❌ ${error.message}`);
    }
  };

  // Función para actualizar estadísticas
  const updateStats = () => {
    const lines = code.split('\n').length
    setCompilationStats(prev => ({
      ...prev,
      lines,
      size: `${(code.length / 1024).toFixed(1)} KB`
    }))
  }

  const menuItems = {
    archivo: [
      { label: 'Nuevo Proyecto', icon: FileText, action: handleNew },
      { label: 'Abrir Archivo', icon: FolderOpen, action: handleOpenFile },
      { label: 'Guardar', icon: Save, action: handleSaveFile }
    ],
    herramientas: [
      { label: 'Compilar', icon: Cpu, action: handleCompile },
    ],
    reportes: [
      { 
        label: 'Tabla de Errores', 
        icon: Bug, 
        action: handleErrorReport,
        loading: reports.isGeneratingReport
      },
      { 
        label: 'Tabla de Símbolos', 
        icon: Hash, 
        action: handleSymbolReport,
        loading: reports.isGeneratingReport
      },
      { 
        label: 'Árbol CST', 
        icon: GitBranch, 
        action: handleCSTInteractive,
        loading: reports.isGeneratingReport
      }
    ]
  }

  return (
    <motion.div 
      className={`w-screen h-screen ${themeClasses.background} flex flex-col overflow-hidden relative transition-all duration-500`}
      variants={containerVariants}
      initial="hidden"
      animate="visible"
    >
      {/* Efectos de fondo */}
      <input
        type="file"
        ref={fileInputRef}
        onChange={handleFileRead}
        accept=".vch"
        className="hidden"
      />
      
      {/* Header con animación */}
      <motion.header 
        className={`${themeClasses.header} border-b mx-3 mt-3 mb-2 px-6 py-3 relative z-50 shadow-2xl shrink-0 rounded-2xl transition-all duration-500`}
        variants={headerVariants}
      >
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4">
            <div className="flex items-center gap-3">
              <motion.div 
                className="relative"
                initial={{ opacity: 0, scale: 0.3, rotate: -720 }}
                animate={{ opacity: 1, scale: 1, rotate: 0 }}
                transition={{ 
                  delay: 0.2, 
                  duration: 1.2, 
                  type: "spring", 
                  stiffness: 120, 
                  damping: 20 
                }}
              >
                <div className="w-10 h-10 bg-gradient-to-br from-emerald-400 via-cyan-500 to-blue-500 rounded-2xl flex items-center justify-center shadow-lg shadow-emerald-500/20">
                  <Code2 className="w-5 h-5 text-white drop-shadow-lg" />
                </div>
                <div className={`absolute -top-0.5 -right-0.5 w-3 h-3 bg-green-400 rounded-full border-2 ${isDarkMode ? 'border-gray-900' : 'border-white'} transition-all duration-500`}></div>
              </motion.div>
              <motion.div
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.2, duration: 0.6 }}
              >
                <h1 className={`${themeClasses.textPrimary} text-base font-bold text-transparent bg-gradient-to-r from-emerald-400 to-cyan-400 bg-clip-text transition-all duration-500`}>
                  VLang→ARM64{hasUnsavedChanges ? '*' : ''}
                </h1>
                <p className={`${themeClasses.textAccent} text-xs transition-all duration-500`}>
                  {currentFile ? `${currentFile}${hasUnsavedChanges ? ' (modificado)' : ''}` : 'Compilador v2.0'}
                </p>
              </motion.div>
            </div>
            
            <motion.div 
              className="flex gap-1 ml-6" 
              ref={menuContainerRef}
              initial={{ opacity: 0, y: -20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: 0.3, duration: 0.6 }}
            >
              {Object.keys(menuItems).map((menu, index) => (
                <div key={menu} className="relative">
                  <motion.button
                    className={`
                      relative px-4 py-2 ${themeClasses.textPrimary} font-medium text-sm rounded-xl
                      transition-all duration-300 ease-out overflow-hidden group
                      backdrop-blur-md border shadow-lg
                      ${isDarkMode 
                        ? 'bg-emerald-500/10 border-emerald-500/20 hover:bg-emerald-500/20' 
                        : 'bg-emerald-400/10 border-emerald-400/20 hover:bg-emerald-400/20'
                      }
                      before:absolute before:inset-0 before:bg-gradient-to-r 
                      before:from-transparent before:via-white/10 before:to-transparent
                      before:translate-x-[-100%] hover:before:translate-x-[100%]
                      before:transition-transform before:duration-1000 before:rounded-xl
                      backdrop-saturate-150 backdrop-brightness-110
                      hover:backdrop-blur-lg hover:backdrop-saturate-200
                      border-2 hover:border-emerald-400/50
                      ${activeMenu === menu ? 
                        (isDarkMode ? 'bg-emerald-500/20 border-emerald-400/50 shadow-md' : 'bg-emerald-400/20 border-emerald-400/50 shadow-md') 
                        : ''}
                    `}
                    onClick={() => setActiveMenu(activeMenu === menu ? null : menu)}
                    variants={buttonHoverVariants}
                    whileHover="hover"
                    whileTap="tap"
                    initial={{ opacity: 0, y: -20 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ delay: 0.4 + index * 0.1, duration: 0.6 }}
                  >
                    <span className="relative z-10">
                      {menu.charAt(0).toUpperCase() + menu.slice(1)}
                    </span>
                  </motion.button>
                  
                  <AnimatePresence>
                    {activeMenu === menu && (
                      <motion.div 
                        className={`absolute top-full left-0 mt-2 ${isDarkMode ? 'bg-black/95' : 'bg-white/95'} backdrop-blur-3xl rounded-2xl shadow-2xl min-w-56 p-3 border border-emerald-500/30`}
                        variants={menuVariants}
                        initial="hidden"
                        animate="visible"
                        exit="exit"
                        style={{ zIndex: 1000 }}
                      >
                        {menuItems[menu].map((item, index) => {
                          const IconComponent = item.icon
                          return (
                            <motion.button
                              key={index}
                              className={`
                                w-full text-left px-4 py-3 text-sm rounded-xl flex items-center justify-between group
                                transition-all duration-200 hover:scale-[1.02] hover:brightness-105
                                backdrop-blur-md border shadow-lg
                                ${isDarkMode 
                                  ? 'text-gray-200 hover:text-emerald-400 bg-emerald-500/10 border-emerald-500/20 hover:bg-emerald-500/15' 
                                  : 'text-gray-800 hover:text-emerald-600 bg-emerald-400/10 border-emerald-400/20 hover:bg-emerald-400/15'
                                }
                                before:absolute before:inset-0 before:bg-gradient-to-r 
                                before:from-transparent before:via-white/10 before:to-transparent
                                before:translate-x-[-100%] hover:before:translate-x-[100%]
                                before:transition-transform before:duration-1000
                                relative overflow-hidden before:rounded-xl
                                backdrop-saturate-150 backdrop-brightness-110
                                hover:backdrop-blur-lg hover:backdrop-saturate-200
                                border-2 hover:border-emerald-400/30
                              `}
                              onClick={item.action}
                              variants={menuItemVariants}
                              initial="hidden"
                              animate="visible"
                              transition={{ delay: index * 0.05 }}
                              whileHover={{ scale: 1.02 }}
                              whileTap={{ scale: 0.98 }}
                            >
                              <div className="flex items-center gap-3 relative z-10">
                                <IconComponent className={`w-4 h-4 ${isDarkMode ? 'group-hover:text-emerald-400' : 'group-hover:text-emerald-600'} transition-colors`} />
                                <span className="font-medium">{item.label}</span>
                              </div>
                            </motion.button>
                          )
                        })}
                      </motion.div>
                    )}
                  </AnimatePresence>
                </div>
              ))}
            </motion.div>
          </div>
          
          <motion.div 
            className="flex items-center gap-3"
            initial={{ opacity: 0, x: 20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 0.2, duration: 0.6 }}
          >
            {/* Switch de tema con fondo dinámico */}
            <motion.button
              onClick={toggleTheme}
              className={`
                flex items-center gap-2 px-3 py-2 rounded-xl border transition-all duration-300
                backdrop-blur-md shadow-lg
                ${isDarkMode 
                  ? 'bg-yellow-500/10 border-yellow-500/20 text-yellow-400 hover:bg-yellow-500/20 hover:shadow-yellow-500/20' 
                  : 'bg-purple-500/10 border-purple-500/20 text-purple-600 hover:bg-purple-500/20 hover:shadow-purple-500/20'
                }
                before:absolute before:inset-0 before:bg-gradient-to-r 
                before:from-transparent before:via-white/10 before:to-transparent
                before:translate-x-[-100%] hover:before:translate-x-[100%]
                before:transition-transform before:duration-1000
                relative overflow-hidden before:rounded-xl
                backdrop-saturate-150 backdrop-brightness-110
                hover:backdrop-blur-lg hover:backdrop-saturate-200
                border-2 hover:border-yellow-400/30
              `}
              variants={buttonHoverVariants}
              whileHover="hover"
              whileTap="tap"
            >
              {isDarkMode ? (
                <Sun className="w-4 h-4 relative z-10" />
              ) : (
                <Moon className="w-4 h-4 relative z-10" />
              )}
              <span className="text-xs font-medium relative z-10">
                {isDarkMode ? 'Claro' : 'Oscuro'}
              </span>
            </motion.button>
            
            {/* Stats con fondos dinámicos */}
            <motion.div 
              className="flex items-center gap-2"
              initial={{ opacity: 0, scale: 0.8 }}
              animate={{ opacity: 1, scale: 1 }}
              transition={{ delay: 0.4, duration: 0.6 }}
            >
              <div className={`flex items-center gap-2 px-2.5 py-1.5 rounded-lg border transition-all duration-500 ${
                isDarkMode ? 'bg-blue-500/10 border-blue-500/20' : 'bg-blue-400/10 border-blue-400/20'
              }`}>
                <Clock className="w-3.5 h-3.5 text-blue-400" />
                <span className="text-blue-400 text-xs font-medium">{compilationStats.time}</span>
              </div>
            </motion.div>
            
            <motion.div 
              className={`flex items-center gap-2 px-3 py-1.5 rounded-lg border transition-all duration-500 ${
                isDarkMode ? 'bg-green-500/10 border-green-500/20' : 'bg-green-400/10 border-green-400/20'
              }`}
              initial={{ opacity: 0, scale: 0.8 }}
              animate={{ opacity: 1, scale: 1 }}
              transition={{ delay: 0.5, duration: 0.6 }}
            >
              <Circle className={`w-2.5 h-2.5 fill-current transition-all duration-500 ${isCompiling ? 'text-yellow-400 ' : 'text-green-400'}`} />
              <span className={`text-xs font-semibold transition-all duration-500 ${isCompiling ? 'text-yellow-400' : 'text-green-400'}`}>
                {isCompiling ? 'COMPILANDO' : 'LISTO'}
              </span>
            </motion.div>
          </motion.div>
        </div>
      </motion.header>

      {/* Contenido principal con animaciones */}
      <main className="flex-1 grid grid-cols-2 gap-3 px-3 pb-3 min-h-0 w-full">
        {/* Editor de código */}
        <motion.section 
          className={`${themeClasses.panel} rounded-3xl border flex flex-col overflow-hidden shadow-2xl transition-all duration-500`}
          variants={panelVariants}
        >
          <motion.div 
            className={`${themeClasses.panelHeader} px-6 py-4 border-b flex justify-between items-center transition-all duration-500`}
            initial={{ opacity: 0, y: -20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.6, duration: 0.6 }}
          >
            <div className="flex items-center gap-4">
              <div className="flex items-center gap-3">
                <motion.div 
                  className="w-8 h-8 bg-gradient-to-br from-emerald-400 to-green-500 rounded-xl flex items-center justify-center shadow-lg"
                  whileHover={{ rotate: 360 }}
                  transition={{ duration: 0.6 }}
                >
                  <Code2 className="w-5 h-5 text-white" />
                </motion.div>
                <div>
                  <h2 className={`${themeClasses.textPrimary} text-lg font-bold transition-all duration-500`}>Editor Vlang</h2>
                  <p className={`${themeClasses.textAccent} text-xs transition-all duration-500`}>Editor de código</p>
                </div>
              </div>
              <motion.div 
                className="flex items-center gap-2 bg-emerald-500/10 px-3 py-1.5 rounded-lg border border-emerald-500/20"
                initial={{ opacity: 0, scale: 0.8 }}
                animate={{ opacity: 1, scale: 1 }}
                transition={{ delay: 0.7, duration: 0.6 }}
              >
                <Circle className="w-2 h-2 text-emerald-400 fill-current" />
                <span className="text-emerald-400 text-xs font-medium">.vch</span>
              </motion.div>
            </div>
            <div className="flex items-center gap-4">
              <motion.button 
                onClick={handleCompile}
                disabled={isCompiling}
                className={`
                  flex items-center gap-2 px-6 py-2.5 rounded-xl font-semibold transition-all duration-300 disabled:opacity-50 disabled:cursor-not-allowed
                  backdrop-blur-md shadow-lg
                  bg-gradient-to-r from-emerald-500/80 to-cyan-500/80 text-white border-emerald-400/30
                  hover:from-emerald-500 hover:to-cyan-500 hover:shadow-xl hover:shadow-emerald-500/20
                  before:absolute before:inset-0 before:bg-gradient-to-r 
                  before:from-transparent before:via-white/10 before:to-transparent
                  before:translate-x-[-100%] hover:before:translate-x-[100%]
                  before:transition-transform before:duration-1000
                  relative overflow-hidden before:rounded-xl
                  backdrop-saturate-150 backdrop-brightness-110
                  hover:backdrop-blur-lg hover:backdrop-saturate-200
                  border-2 hover:border-emerald-400/50
                `}
                variants={buttonHoverVariants}
                whileHover="hover"
                whileTap="tap"
              >
                {isCompiling ? (
                  <>
                    <div className="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin relative z-10"></div>
                    <span className="relative z-10">Compilando...</span>
                  </>
                ) : (
                  <>
                    <Play className="w-4 h-4 relative z-10" />
                    <span className="relative z-10">Compilar</span>
                  </>
                )}
              </motion.button>
            </div>
          </motion.div>
          <div className="flex-1 relative">
            <div className={`absolute inset-0 ${isDarkMode ? 'bg-gradient-to-br from-gray-900/50 to-black/50' : 'bg-gradient-to-br from-white/30 to-gray-100/30'} transition-all duration-500`}></div>
            <motion.div 
              className="relative h-full p-6"
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ delay: 0.8, duration: 0.6 }}
            >
              <textarea
                className={`w-full h-full bg-transparent border-none ${themeClasses.editor} font-mono text-sm leading-7 resize-none outline-none scrollbar-thin scrollbar-thumb-emerald-500/30 scrollbar-track-gray-800/30 scrollbar-thumb-rounded-full transition-all duration-500`}
                value={code}
                onChange={(e) => {
                  setCode(e.target.value)
                  updateStats()
                }}
                onKeyDown={(e) => {
                  // Manejar tecla Tab para indentación
                  if (e.key === 'Tab') {
                    e.preventDefault() // Evitar comportamiento por defecto
                    
                    const target = e.target
                    const start = target.selectionStart
                    const end = target.selectionEnd
                    
                    if (e.shiftKey) {
                      // Shift + Tab: Remover indentación
                      const lines = code.split('\n')
                      const startLineIndex = code.substring(0, start).split('\n').length - 1
                      const endLineIndex = code.substring(0, end).split('\n').length - 1
                      
                      let newCode = ''
                      let newStart = start
                      let newEnd = end
                      
                      for (let i = 0; i < lines.length; i++) {
                        if (i >= startLineIndex && i <= endLineIndex) {
                          if (lines[i].startsWith('  ')) {
                            lines[i] = lines[i].substring(2)
                            if (i === startLineIndex) newStart = Math.max(0, start - 2)
                            if (i <= endLineIndex) newEnd = Math.max(0, end - 2)
                          } else if (lines[i].startsWith('\t')) {
                            lines[i] = lines[i].substring(1)
                            if (i === startLineIndex) newStart = Math.max(0, start - 1)
                            if (i <= endLineIndex) newEnd = Math.max(0, end - 1)
                          }
                        }
                        newCode += lines[i] + (i < lines.length - 1 ? '\n' : '')
                      }
                      
                      setCode(newCode)
                      updateStats()
                      
                      // Restaurar selección
                      setTimeout(() => {
                        target.setSelectionRange(newStart, newEnd)
                      }, 0)
                      
                    } else {
                      // Tab normal: Agregar indentación
                      if (start === end) {
                        // No hay selección - insertar tab en posición actual
                        const newCode = code.substring(0, start) + '  ' + code.substring(end)
                        setCode(newCode)
                        updateStats()
                        
                        // Mover cursor después del tab
                        setTimeout(() => {
                          target.setSelectionRange(start + 2, start + 2)
                        }, 0)
                      } else {
                        // Hay selección - indentar líneas seleccionadas
                        const lines = code.split('\n')
                        const startLineIndex = code.substring(0, start).split('\n').length - 1
                        const endLineIndex = code.substring(0, end).split('\n').length - 1
                        
                        let newCode = ''
                        let addedChars = 0
                        
                        for (let i = 0; i < lines.length; i++) {
                          if (i >= startLineIndex && i <= endLineIndex) {
                            lines[i] = '  ' + lines[i]
                            addedChars += 2
                          }
                          newCode += lines[i] + (i < lines.length - 1 ? '\n' : '')
                        }
                        
                        setCode(newCode)
                        updateStats()
                        
                        // Ajustar selección
                        setTimeout(() => {
                          target.setSelectionRange(start + 2, end + addedChars)
                        }, 0)
                      }
                    }
                  }
                  
                  // Manejar Enter para auto-indentación
                  else if (e.key === 'Enter') {
                    const target = e.target
                    const start = target.selectionStart
                    const lines = code.substring(0, start).split('\n')
                    const currentLine = lines[lines.length - 1]
                    
                    // Detectar indentación actual
                    const indent = currentLine.match(/^(\s*)/)[1]
                    
                    // Auto-indentar después de llaves o dos puntos
                    const shouldIndentMore = /[{:]$/.test(currentLine.trim())
                    const extraIndent = shouldIndentMore ? '  ' : ''
                    
                    setTimeout(() => {
                      const newStart = start + 1 + indent.length + extraIndent.length
                      const newCode = code.substring(0, start) + '\n' + indent + extraIndent + code.substring(start)
                      setCode(newCode)
                      updateStats()
                      
                      // Posicionar cursor
                      setTimeout(() => {
                        target.setSelectionRange(newStart, newStart)
                      }, 0)
                    }, 0)
                  }
                }}
                placeholder="// Escribe tu código Vlang aquí..."
                spellCheck={false}
              />
              <motion.div 
                className={`absolute right-4 bottom-4 ${isDarkMode ? 'bg-black/60' : 'bg-white/60'} backdrop-blur-md px-3 py-2 rounded-lg border border-emerald-500/20 text-emerald-400 text-xs font-mono transition-all duration-500`}
                initial={{ opacity: 0, scale: 0.5 }}
                animate={{ opacity: 1, scale: 1 }}
                transition={{ delay: 1, duration: 0.6 }}
              >
                Ln {code.split('\n').length}
              </motion.div>
            </motion.div>
          </div>
        </motion.section>

        {/* Consola de salida */}
        <motion.section 
          className={`${themeClasses.panel} rounded-3xl border border-cyan-500/20 flex flex-col overflow-hidden shadow-2xl transition-all duration-500`}
          variants={panelVariants}
          transition={{ delay: 0.2 }}
        >
          <motion.div 
            className={`${themeClasses.panelHeader.replace('emerald', 'cyan')} px-6 py-4 border-b border-cyan-500/20 flex justify-between items-center transition-all duration-500`}
            initial={{ opacity: 0, y: -20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.8, duration: 0.6 }}
          >
            <div className="flex items-center gap-4">
              <div className="flex items-center gap-3">
                <motion.div 
                  className="w-8 h-8 bg-gradient-to-br from-cyan-400 to-blue-500 rounded-xl flex items-center justify-center shadow-lg"
                  whileHover={{ rotate: 360 }}
                  transition={{ duration: 0.6 }}
                >
                  <Terminal className="w-5 h-5 text-white" />
                </motion.div>
                <div>
                  <h2 className={`${themeClasses.textPrimary} text-lg font-bold transition-all duration-500`}>Consola ARM64</h2>
                  <p className="text-cyan-400 text-xs transition-all duration-500">Salida de Compilación</p>
                </div>
              </div>
              <motion.div 
                className="flex items-center gap-2 bg-cyan-500/10 px-3 py-1.5 rounded-lg border border-cyan-500/20"
                initial={{ opacity: 0, scale: 0.8 }}
                animate={{ opacity: 1, scale: 1 }}
                transition={{ delay: 0.9, duration: 0.6 }}
              >
                <Circle className="w-2 h-2 text-cyan-400 fill-current " />
                <span className="text-cyan-400 text-xs font-medium">ASM</span>
              </motion.div>
            </div>
            <div className="flex items-center gap-3">
              <motion.div 
                className={`flex items-center gap-2 px-3 py-2 rounded-xl border transition-all duration-500 ${
                  isDarkMode ? 'bg-gray-800/50 border-gray-700/50' : 'bg-gray-200/50 border-gray-300/50'
                }`}
                initial={{ opacity: 0, scale: 0.8 }}
                animate={{ opacity: 1, scale: 1 }}
                transition={{ delay: 1, duration: 0.6 }}
              >
                <Cpu className={`w-4 h-4 ${themeClasses.textSecondary}`} />
                <span className={`${themeClasses.textSecondary} text-xs font-medium`}>ARM64</span>
              </motion.div>
              <motion.button 
                className={`
                  flex items-center gap-2 px-4 py-2 rounded-xl text-xs font-medium 
                  transition-all duration-300
                  backdrop-blur-md border shadow-lg
                  ${isDarkMode 
                    ? 'bg-red-500/10 border-red-500/20 text-red-400 hover:bg-red-500/20 hover:shadow-red-500/20' 
                    : 'bg-red-400/10 border-red-400/20 text-red-600 hover:bg-red-400/20 hover:shadow-red-400/20'
                  }
                  before:absolute before:inset-0 before:bg-gradient-to-r 
                  before:from-transparent before:via-white/10 before:to-transparent
                  before:translate-x-[-100%] hover:before:translate-x-[100%]
                  before:transition-transform before:duration-1000
                  relative overflow-hidden
                  backdrop-saturate-150 backdrop-brightness-110
                  hover:backdrop-blur-lg hover:backdrop-saturate-200
                  shadow-lg shadow-red-500/10
                  hover:shadow-xl hover:shadow-red-500/20
                  border-2 hover:border-red-400/30
                `}
                onClick={() => setOutput('🔧 Consola limpiada')}
                variants={buttonHoverVariants}
                whileHover="hover"
                whileTap="tap"
              >
                <Trash2 className="w-3 h-3 relative z-10" />
                <span className="relative z-10">Limpiar</span>
              </motion.button>
            </div>
          </motion.div>
          <div className="flex-1 relative overflow-hidden">
            <div className={`absolute inset-0 ${isDarkMode ? 'bg-gradient-to-br from-gray-900/50 to-black/50' : 'bg-gradient-to-br from-white/30 to-gray-100/30'} transition-all duration-500`}></div>
            <motion.div 
              className="relative h-full"
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ delay: 1, duration: 0.6 }}
            >
              <pre className={`
                w-full h-full p-6 m-0
                ${themeClasses.console} 
                font-mono text-sm leading-7 
                whitespace-pre-wrap 
                overflow-y-auto overflow-x-hidden
                scrollbar-thin 
                scrollbar-thumb-cyan-500/40 
                scrollbar-track-transparent
                scrollbar-thumb-rounded-full
                hover:scrollbar-thumb-cyan-500/60
                transition-all duration-500
                resize-none
                border-none
                outline-none
                bg-transparent
              `}>
                {output}
              </pre>
              <motion.div 
                className={`absolute right-4 bottom-4 flex items-center gap-2 ${isDarkMode ? 'bg-black/60' : 'bg-white/60'} backdrop-blur-xl px-3 py-2 rounded-lg border border-cyan-500/20 transition-all duration-500 pointer-events-none`}
                initial={{ opacity: 0, scale: 0.5 }}
                animate={{ opacity: 1, scale: 1 }}
                transition={{ delay: 1.2, duration: 0.6 }}
              >
                {isCompiling ? (
                  <>
                    <div className="w-2 h-2 bg-yellow-400 rounded-full"></div>
                    <span className="text-yellow-400 text-xs font-medium">Compilando</span>
                  </>
                ) : (
                  <>
                    <div className="w-2 h-2 bg-green-400 rounded-full"></div>
                    <span className="text-green-400 text-xs font-medium">Listo</span>
                  </>
                )}
              </motion.div>
            </motion.div>
          </div>
        </motion.section>
      </main>

      {/* Footer */}
      <motion.footer 
        className={`${themeClasses.footer} border-t mx-3 mb-3 px-6 py-3 rounded-2xl flex justify-between items-center text-sm transition-all duration-500`}
        variants={footerVariants}
      >
        <div className="flex items-center gap-6">
          <motion.div 
            className="flex items-center gap-2"
            initial={{ opacity: 0, x: -20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 0.6, duration: 0.6 }}
          >
            <div className="w-2 h-2 bg-emerald-400 rounded-full "></div>
            <span className={`${themeClasses.textAccent} font-medium transition-all duration-500`}>Compilador Vlang v2.0</span>
          </motion.div>
          <motion.div 
            className={`${themeClasses.textSecondary} transition-all duration-500`}
            initial={{ opacity: 0, x: -20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 0.7, duration: 0.6 }}
          >
            ARM64 AArch64 Target
          </motion.div>
        </div>
        
        <div className="flex items-center gap-6">
          <motion.div 
            className="flex items-center gap-4"
            initial={{ opacity: 0, x: 20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 0.6, duration: 0.6 }}
          >
            <span className={`${themeClasses.textSecondary} transition-all duration-500`}>
              <span className={`${themeClasses.textAccent} font-medium`}>{(code.match(/fn\s+\w+/g) || []).length}</span> funciones
            </span>
            <span className={`${themeClasses.textSecondary} transition-all duration-500`}>
              <span className="text-cyan-400 font-medium">{(code.match(/mut\s+\w+|:\s*=/g) || []).length}</span> variables
            </span>
            <span className={`${themeClasses.textSecondary} transition-all duration-500`}>
              Tamaño: <span className="text-purple-400 font-medium">{compilationStats.size}</span>
            </span>
          </motion.div>
          
          <motion.div 
            className={`flex items-center gap-2 px-3 py-1.5 rounded-lg border transition-all duration-500 ${
              errors.length > 0 
                ? (isDarkMode ? 'bg-red-500/10 border-red-500/20' : 'bg-red-400/10 border-red-400/20')
                : (isDarkMode ? 'bg-green-500/10 border-green-500/20' : 'bg-green-400/10 border-green-400/20')
            }`}
            initial={{ opacity: 0, scale: 0.8 }}
            animate={{ opacity: 1, scale: 1 }}
            transition={{ delay: 0.8, duration: 0.6 }}
          >
            {errors.length > 0 ? (
              <>
                <Bug className="w-4 h-4 text-red-400" />
                <span className="text-red-400 font-medium text-xs">
                  {errors.length} Error{errors.length !== 1 ? 'es' : ''}
                </span>
              </>
            ) : (
              <>
                <CheckCircle className="w-4 h-4 text-green-400" />
                <span className="text-green-400 font-medium text-xs">Sin Errores</span>
              </>
            )}
          </motion.div>
        </div>
      </motion.footer>

      {/* Overlays con animaciones */}
      <AnimatePresence>
        {isCompiling && (
          <motion.div 
            className="absolute inset-0 bg-black/20 backdrop-blur-sm flex items-center justify-center z-50"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.3 }}
          >
            <motion.div 
              className={`${isDarkMode ? 'bg-black/90' : 'bg-white/90'} backdrop-blur-2xl rounded-3xl border border-emerald-500/30 p-8 shadow-2xl transition-all duration-500`}
              initial={{ scale: 0.8, opacity: 0 }}
              animate={{ scale: 1, opacity: 1 }}
              exit={{ scale: 0.8, opacity: 0 }}
              transition={{ type: "spring", stiffness: 200, damping: 20 }}
            >
              <div className="flex items-center gap-4">
                <div className="w-8 h-8 border-4 border-emerald-500/30 border-t-emerald-400 rounded-full animate-spin"></div>
                <div>
                  <h3 className={`${themeClasses.textPrimary} font-bold text-lg transition-all duration-500`}>Compilando a ARM64...</h3>
                  <p className={`${themeClasses.textAccent} text-sm transition-all duration-500`}>Generando código optimizado</p>
                </div>
              </div>
            </motion.div>
          </motion.div>
        )}
      </AnimatePresence>

      <AnimatePresence>
        {reports.isGeneratingReport && (
          <motion.div 
            className="absolute inset-0 bg-black/20 backdrop-blur-sm flex items-center justify-center z-50"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.3 }}
          >
            <motion.div 
              className={`${isDarkMode ? 'bg-black/90' : 'bg-white/90'} backdrop-blur-2xl rounded-3xl border border-emerald-500/30 p-8 shadow-2xl transition-all duration-500`}
              initial={{ scale: 0.8, opacity: 0 }}
              animate={{ scale: 1, opacity: 1 }}
              exit={{ scale: 0.8, opacity: 0 }}
              transition={{ type: "spring", stiffness: 200, damping: 20 }}
            >
              <div className="flex items-center gap-4">
                <Loader2 className="w-8 h-8 text-emerald-400 animate-spin" />
                <div>
                  <h3 className={`${themeClasses.textPrimary} font-bold text-lg transition-all duration-500`}>
                    {reports.reportProgress || 'Generando reporte...'}
                  </h3>
                  <p className={`${themeClasses.textAccent} text-sm transition-all duration-500`}>
                    Procesando código y símbolos
                  </p>
                </div>
              </div>
            </motion.div>
          </motion.div>
        )}
      </AnimatePresence>

      <AnimatePresence>
        {showReport && (
          <ReportViewer
            reportData={reportData}
            reportType={reportType}
            isDarkMode={isDarkMode}
            onClose={() => setShowReport(false)}
          />
        )}
      </AnimatePresence>

      <AnimatePresence>
        {showCST && (
          <CSTViewer
            svgContent={cstContent}
            isDarkMode={isDarkMode}
            onClose={() => setShowCST(false)}
          />
        )}
      </AnimatePresence>
    </motion.div>
  )
}

export default App