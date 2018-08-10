package gl

// #cgo darwin CFLAGS:	-DTAG_DARWIN
// #cgo darwin LDFLAGS:	-framework OpenGL
// #cgo linux freebsd CFLAGS:	-DTAG_POSIX
// #cgo linux freebsd LDFLAGS:	-lGL
// #cgo windows CFLAGS:		-DTAG_WINDOWS
// #cgo windows LDFLAGS:	-lopengl32
// #cgo egl CFLAGS:		-DTAG_EGL
// #cgo egl LDFLAGS:	-lEGL
// #if defined(_WIN32) && !defined(APIENTRY) && !defined(__CYGWIN__) && !defined(__SCITECH_SNAP__)
// #ifndef WIN32_LEAN_AND_MEAN
// #define WIN32_LEAN_AND_MEAN 1
// #endif
// #include <windows.h>
// #endif
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// #include <stddef.h>
// #ifndef GLEXT_64_TYPES_DEFINED
// /* This code block is duplicated in glxext.h, so must be protected */
// #define GLEXT_64_TYPES_DEFINED
// /* Define int32_t, int64_t, and uint64_t types for UST/MSC */
// /* (as used in the GL_EXT_timer_query extension). */
// #if defined(__STDC_VERSION__) && __STDC_VERSION__ >= 199901L
// #include <inttypes.h>
// #elif defined(__sun__) || defined(__digital__)
// #include <inttypes.h>
// #if defined(__STDC__)
// #if defined(__arch64__) || defined(_LP64)
// typedef long int int64_t;
// typedef unsigned long int uint64_t;
// #else
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #endif /* __arch64__ */
// #endif /* __STDC__ */
// #elif defined( __VMS ) || defined(__sgi)
// #include <inttypes.h>
// #elif defined(__SCO__) || defined(__USLC__)
// #include <stdint.h>
// #elif defined(__UNIXOS2__) || defined(__SOL64__)
// typedef long int int32_t;
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #elif defined(_WIN32) && defined(__GNUC__)
// #include <stdint.h>
// #elif defined(_WIN32)
// typedef __int32 int32_t;
// typedef __int64 int64_t;
// typedef unsigned __int64 uint64_t;
// #else
// /* Fallback if nothing above works */
// #include <inttypes.h>
// #endif
// #endif
// #if defined(TAG_EGL)
// 	#include <stdlib.h>
// 	#include <EGL/egl.h>
// 	void* ProcAddress(const char* name) {
// 		return eglGetProcAddress(name);
// 	}
// #elif defined(TAG_WINDOWS)
// 	#define WIN32_LEAN_AND_MEAN 1
// 	#include <windows.h>
// 	#include <stdlib.h>
// 	static HMODULE ogl32dll = NULL;
// 	void* ProcAddress(const char* name) {
// 		void* pf = wglGetProcAddress((LPCSTR) name);
// 		if (pf) {
// 			return pf;
// 		}
// 		if (ogl32dll == NULL) {
// 			ogl32dll = LoadLibraryA("opengl32.dll");
// 		}
// 		return GetProcAddress(ogl32dll, (LPCSTR) name);
// 	}
// #elif defined(TAG_DARWIN)
// 	#include <stdlib.h>
// 	#include <dlfcn.h>
// 	void* ProcAddress(const char* name) {
// 		return dlsym(RTLD_DEFAULT, name);
// 	}
// #elif defined(TAG_POSIX)
// 	#include <stdlib.h>
// 	#include <GL/glx.h>
// 	void* ProcAddress(const char* name) {
// 		return glXGetProcAddress(name);
// 	}
// #endif
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef signed char GLbyte;
// typedef short GLshort;
// typedef int GLint;
// typedef unsigned char GLubyte;
// typedef unsigned short GLushort;
// typedef unsigned int GLuint;
// typedef int GLsizei;
// typedef float GLfloat;
// typedef double GLdouble;
// typedef char GLchar;
// typedef ptrdiff_t GLintptr;
// typedef ptrdiff_t GLsizeiptr;
// typedef int64_t GLint64;
// typedef uint64_t GLuint64;
// typedef uint64_t GLuint64EXT;
// typedef struct __GLsync *GLsync;
// struct _cl_context;
// struct _cl_event;
// typedef void (APIENTRY *GLDEBUGPROC)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCARB)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCKHR)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// extern void debugCallback(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar* message, const void* userParam);
// static void APIENTRY cDebugCallback(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar* message, const void* userParam) {debugCallback(source, type, id, severity, length, message, userParam);}
// typedef void  (APIENTRYP GPACTIVESHADERPROGRAM)(GLuint  pipeline, GLuint  program);
// typedef void  (APIENTRYP GPACTIVETEXTURE)(GLenum  texture);
// typedef void  (APIENTRYP GPATTACHSHADER)(GLuint  program, GLuint  shader);
// typedef void  (APIENTRYP GPBEGINCONDITIONALRENDER)(GLuint  id, GLenum  mode);
// typedef void  (APIENTRYP GPBEGINQUERY)(GLenum  target, GLuint  id);
// typedef void  (APIENTRYP GPBEGINQUERYINDEXED)(GLenum  target, GLuint  index, GLuint  id);
// typedef void  (APIENTRYP GPBEGINTRANSFORMFEEDBACK)(GLenum  primitiveMode);
// typedef void  (APIENTRYP GPBINDATTRIBLOCATION)(GLuint  program, GLuint  index, const GLchar * name);
// typedef void  (APIENTRYP GPBINDBUFFER)(GLenum  target, GLuint  buffer);
// typedef void  (APIENTRYP GPBINDBUFFERBASE)(GLenum  target, GLuint  index, GLuint  buffer);
// typedef void  (APIENTRYP GPBINDBUFFERRANGE)(GLenum  target, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPBINDBUFFERSBASE)(GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers);
// typedef void  (APIENTRYP GPBINDBUFFERSRANGE)(GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizeiptr * sizes);
// typedef void  (APIENTRYP GPBINDFRAGDATALOCATION)(GLuint  program, GLuint  color, const GLchar * name);
// typedef void  (APIENTRYP GPBINDFRAGDATALOCATIONINDEXED)(GLuint  program, GLuint  colorNumber, GLuint  index, const GLchar * name);
// typedef void  (APIENTRYP GPBINDFRAMEBUFFER)(GLenum  target, GLuint  framebuffer);
// typedef void  (APIENTRYP GPBINDIMAGETEXTURE)(GLuint  unit, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  access, GLenum  format);
// typedef void  (APIENTRYP GPBINDIMAGETEXTURES)(GLuint  first, GLsizei  count, const GLuint * textures);
// typedef void  (APIENTRYP GPBINDPROGRAMPIPELINE)(GLuint  pipeline);
// typedef void  (APIENTRYP GPBINDRENDERBUFFER)(GLenum  target, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPBINDSAMPLER)(GLuint  unit, GLuint  sampler);
// typedef void  (APIENTRYP GPBINDSAMPLERS)(GLuint  first, GLsizei  count, const GLuint * samplers);
// typedef void  (APIENTRYP GPBINDTEXTURE)(GLenum  target, GLuint  texture);
// typedef void  (APIENTRYP GPBINDTEXTUREUNIT)(GLuint  unit, GLuint  texture);
// typedef void  (APIENTRYP GPBINDTEXTURES)(GLuint  first, GLsizei  count, const GLuint * textures);
// typedef void  (APIENTRYP GPBINDTRANSFORMFEEDBACK)(GLenum  target, GLuint  id);
// typedef void  (APIENTRYP GPBINDVERTEXARRAY)(GLuint  array);
// typedef void  (APIENTRYP GPBINDVERTEXBUFFER)(GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride);
// typedef void  (APIENTRYP GPBINDVERTEXBUFFERS)(GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides);
// typedef void  (APIENTRYP GPBLENDCOLOR)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPBLENDEQUATION)(GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATE)(GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATEI)(GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATEIARB)(GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONI)(GLuint  buf, GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONIARB)(GLuint  buf, GLenum  mode);
// typedef void  (APIENTRYP GPBLENDFUNC)(GLenum  sfactor, GLenum  dfactor);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATE)(GLenum  sfactorRGB, GLenum  dfactorRGB, GLenum  sfactorAlpha, GLenum  dfactorAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATEI)(GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATEIARB)(GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCI)(GLuint  buf, GLenum  src, GLenum  dst);
// typedef void  (APIENTRYP GPBLENDFUNCIARB)(GLuint  buf, GLenum  src, GLenum  dst);
// typedef void  (APIENTRYP GPBLITFRAMEBUFFER)(GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef void  (APIENTRYP GPBLITNAMEDFRAMEBUFFER)(GLuint  readFramebuffer, GLuint  drawFramebuffer, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef void  (APIENTRYP GPBUFFERDATA)(GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage);
// typedef void  (APIENTRYP GPBUFFERPAGECOMMITMENTARB)(GLenum  target, GLintptr  offset, GLsizei  size, GLboolean  commit);
// typedef void  (APIENTRYP GPBUFFERSTORAGE)(GLenum  target, GLsizeiptr  size, const void * data, GLbitfield  flags);
// typedef void  (APIENTRYP GPBUFFERSUBDATA)(GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data);
// typedef GLenum  (APIENTRYP GPCHECKFRAMEBUFFERSTATUS)(GLenum  target);
// typedef GLenum  (APIENTRYP GPCHECKNAMEDFRAMEBUFFERSTATUS)(GLuint  framebuffer, GLenum  target);
// typedef void  (APIENTRYP GPCLAMPCOLOR)(GLenum  target, GLenum  clamp);
// typedef void  (APIENTRYP GPCLEAR)(GLbitfield  mask);
// typedef void  (APIENTRYP GPCLEARBUFFERDATA)(GLenum  target, GLenum  internalformat, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARBUFFERSUBDATA)(GLenum  target, GLenum  internalformat, GLintptr  offset, GLsizeiptr  size, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARBUFFERFI)(GLenum  buffer, GLint  drawbuffer, GLfloat  depth, GLint  stencil);
// typedef void  (APIENTRYP GPCLEARBUFFERFV)(GLenum  buffer, GLint  drawbuffer, const GLfloat * value);
// typedef void  (APIENTRYP GPCLEARBUFFERIV)(GLenum  buffer, GLint  drawbuffer, const GLint * value);
// typedef void  (APIENTRYP GPCLEARBUFFERUIV)(GLenum  buffer, GLint  drawbuffer, const GLuint * value);
// typedef void  (APIENTRYP GPCLEARCOLOR)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPCLEARDEPTH)(GLdouble  depth);
// typedef void  (APIENTRYP GPCLEARDEPTHF)(GLfloat  d);
// typedef void  (APIENTRYP GPCLEARNAMEDBUFFERDATA)(GLuint  buffer, GLenum  internalformat, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARNAMEDBUFFERSUBDATA)(GLuint  buffer, GLenum  internalformat, GLintptr  offset, GLsizei  size, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERFI)(GLuint  framebuffer, GLenum  buffer, const GLfloat  depth, GLint  stencil);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERFV)(GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLfloat * value);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERIV)(GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLint * value);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERUIV)(GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLuint * value);
// typedef void  (APIENTRYP GPCLEARSTENCIL)(GLint  s);
// typedef void  (APIENTRYP GPCLEARTEXIMAGE)(GLuint  texture, GLint  level, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARTEXSUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * data);
// typedef GLenum  (APIENTRYP GPCLIENTWAITSYNC)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// typedef void  (APIENTRYP GPCLIPCONTROL)(GLenum  origin, GLenum  depth);
// typedef void  (APIENTRYP GPCOLORMASK)(GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha);
// typedef void  (APIENTRYP GPCOLORMASKI)(GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a);
// typedef void  (APIENTRYP GPCOMPILESHADER)(GLuint  shader);
// typedef void  (APIENTRYP GPCOMPILESHADERINCLUDEARB)(GLuint  shader, GLsizei  count, const GLchar *const* path, const GLint * length);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE1D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE2D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE3D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE1D)(GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXTURESUBIMAGE1D)(GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXTURESUBIMAGE2D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXTURESUBIMAGE3D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOPYBUFFERSUBDATA)(GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPCOPYIMAGESUBDATA)(GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth);
// typedef void  (APIENTRYP GPCOPYNAMEDBUFFERSUBDATA)(GLuint  readBuffer, GLuint  writeBuffer, GLintptr  readOffset, GLintptr  writeOffset, GLsizei  size);
// typedef void  (APIENTRYP GPCOPYTEXIMAGE1D)(GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLint  border);
// typedef void  (APIENTRYP GPCOPYTEXIMAGE2D)(GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLint  border);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE1D)(GLenum  target, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXTURESUBIMAGE1D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width);
// typedef void  (APIENTRYP GPCOPYTEXTURESUBIMAGE2D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXTURESUBIMAGE3D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCREATEBUFFERS)(GLsizei  n, GLuint * buffers);
// typedef void  (APIENTRYP GPCREATEFRAMEBUFFERS)(GLsizei  n, GLuint * framebuffers);
// typedef GLuint  (APIENTRYP GPCREATEPROGRAM)();
// typedef void  (APIENTRYP GPCREATEPROGRAMPIPELINES)(GLsizei  n, GLuint * pipelines);
// typedef void  (APIENTRYP GPCREATEQUERIES)(GLenum  target, GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPCREATERENDERBUFFERS)(GLsizei  n, GLuint * renderbuffers);
// typedef void  (APIENTRYP GPCREATESAMPLERS)(GLsizei  n, GLuint * samplers);
// typedef GLuint  (APIENTRYP GPCREATESHADER)(GLenum  type);
// typedef GLuint  (APIENTRYP GPCREATESHADERPROGRAMV)(GLenum  type, GLsizei  count, const GLchar *const* strings);
// typedef GLsync  (APIENTRYP GPCREATESYNCFROMCLEVENTARB)(struct _cl_context * context, struct _cl_event * event, GLbitfield  flags);
// typedef void  (APIENTRYP GPCREATETEXTURES)(GLenum  target, GLsizei  n, GLuint * textures);
// typedef void  (APIENTRYP GPCREATETRANSFORMFEEDBACKS)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPCREATEVERTEXARRAYS)(GLsizei  n, GLuint * arrays);
// typedef void  (APIENTRYP GPCULLFACE)(GLenum  mode);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACK)(GLDEBUGPROC  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACKARB)(GLDEBUGPROCARB  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACKKHR)(GLDEBUGPROCKHR  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROL)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROLARB)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROLKHR)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERT)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERTARB)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERTKHR)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDELETEBUFFERS)(GLsizei  n, const GLuint * buffers);
// typedef void  (APIENTRYP GPDELETEFRAMEBUFFERS)(GLsizei  n, const GLuint * framebuffers);
// typedef void  (APIENTRYP GPDELETENAMEDSTRINGARB)(GLint  namelen, const GLchar * name);
// typedef void  (APIENTRYP GPDELETEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPDELETEPROGRAMPIPELINES)(GLsizei  n, const GLuint * pipelines);
// typedef void  (APIENTRYP GPDELETEQUERIES)(GLsizei  n, const GLuint * ids);
// typedef void  (APIENTRYP GPDELETERENDERBUFFERS)(GLsizei  n, const GLuint * renderbuffers);
// typedef void  (APIENTRYP GPDELETESAMPLERS)(GLsizei  count, const GLuint * samplers);
// typedef void  (APIENTRYP GPDELETESHADER)(GLuint  shader);
// typedef void  (APIENTRYP GPDELETESYNC)(GLsync  sync);
// typedef void  (APIENTRYP GPDELETETEXTURES)(GLsizei  n, const GLuint * textures);
// typedef void  (APIENTRYP GPDELETETRANSFORMFEEDBACKS)(GLsizei  n, const GLuint * ids);
// typedef void  (APIENTRYP GPDELETEVERTEXARRAYS)(GLsizei  n, const GLuint * arrays);
// typedef void  (APIENTRYP GPDEPTHFUNC)(GLenum  func);
// typedef void  (APIENTRYP GPDEPTHMASK)(GLboolean  flag);
// typedef void  (APIENTRYP GPDEPTHRANGE)(GLdouble  xnear, GLdouble  xfar);
// typedef void  (APIENTRYP GPDEPTHRANGEARRAYV)(GLuint  first, GLsizei  count, const GLdouble * v);
// typedef void  (APIENTRYP GPDEPTHRANGEINDEXED)(GLuint  index, GLdouble  n, GLdouble  f);
// typedef void  (APIENTRYP GPDEPTHRANGEF)(GLfloat  n, GLfloat  f);
// typedef void  (APIENTRYP GPDETACHSHADER)(GLuint  program, GLuint  shader);
// typedef void  (APIENTRYP GPDISABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPDISABLEVERTEXARRAYATTRIB)(GLuint  vaobj, GLuint  index);
// typedef void  (APIENTRYP GPDISABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPDISABLEI)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTE)(GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTEGROUPSIZEARB)(GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z, GLuint  group_size_x, GLuint  group_size_y, GLuint  group_size_z);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTEINDIRECT)(GLintptr  indirect);
// typedef void  (APIENTRYP GPDRAWARRAYS)(GLenum  mode, GLint  first, GLsizei  count);
// typedef void  (APIENTRYP GPDRAWARRAYSINDIRECT)(GLenum  mode, const void * indirect);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCED)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCEDBASEINSTANCE)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWBUFFER)(GLenum  buf);
// typedef void  (APIENTRYP GPDRAWBUFFERS)(GLsizei  n, const GLenum * bufs);
// typedef void  (APIENTRYP GPDRAWELEMENTS)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices);
// typedef void  (APIENTRYP GPDRAWELEMENTSBASEVERTEX)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSINDIRECT)(GLenum  mode, GLenum  type, const void * indirect);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCED)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEINSTANCE)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEVERTEX)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWRANGEELEMENTS)(GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices);
// typedef void  (APIENTRYP GPDRAWRANGEELEMENTSBASEVERTEX)(GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACK)(GLenum  mode, GLuint  id);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKINSTANCED)(GLenum  mode, GLuint  id, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKSTREAM)(GLenum  mode, GLuint  id, GLuint  stream);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED)(GLenum  mode, GLuint  id, GLuint  stream, GLsizei  instancecount);
// typedef void  (APIENTRYP GPENABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPENABLEVERTEXARRAYATTRIB)(GLuint  vaobj, GLuint  index);
// typedef void  (APIENTRYP GPENABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPENABLEI)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPENDCONDITIONALRENDER)();
// typedef void  (APIENTRYP GPENDQUERY)(GLenum  target);
// typedef void  (APIENTRYP GPENDQUERYINDEXED)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPENDTRANSFORMFEEDBACK)();
// typedef GLsync  (APIENTRYP GPFENCESYNC)(GLenum  condition, GLbitfield  flags);
// typedef void  (APIENTRYP GPFINISH)();
// typedef void  (APIENTRYP GPFLUSH)();
// typedef void  (APIENTRYP GPFLUSHMAPPEDBUFFERRANGE)(GLenum  target, GLintptr  offset, GLsizeiptr  length);
// typedef void  (APIENTRYP GPFLUSHMAPPEDNAMEDBUFFERRANGE)(GLuint  buffer, GLintptr  offset, GLsizei  length);
// typedef void  (APIENTRYP GPFRAMEBUFFERPARAMETERI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPFRAMEBUFFERRENDERBUFFER)(GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE1D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE2D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE3D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  zoffset);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURELAYER)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer);
// typedef void  (APIENTRYP GPFRONTFACE)(GLenum  mode);
// typedef void  (APIENTRYP GPGENBUFFERS)(GLsizei  n, GLuint * buffers);
// typedef void  (APIENTRYP GPGENFRAMEBUFFERS)(GLsizei  n, GLuint * framebuffers);
// typedef void  (APIENTRYP GPGENPROGRAMPIPELINES)(GLsizei  n, GLuint * pipelines);
// typedef void  (APIENTRYP GPGENQUERIES)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPGENRENDERBUFFERS)(GLsizei  n, GLuint * renderbuffers);
// typedef void  (APIENTRYP GPGENSAMPLERS)(GLsizei  count, GLuint * samplers);
// typedef void  (APIENTRYP GPGENTEXTURES)(GLsizei  n, GLuint * textures);
// typedef void  (APIENTRYP GPGENTRANSFORMFEEDBACKS)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPGENVERTEXARRAYS)(GLsizei  n, GLuint * arrays);
// typedef void  (APIENTRYP GPGENERATEMIPMAP)(GLenum  target);
// typedef void  (APIENTRYP GPGENERATETEXTUREMIPMAP)(GLuint  texture);
// typedef void  (APIENTRYP GPGETACTIVEATOMICCOUNTERBUFFERIV)(GLuint  program, GLuint  bufferIndex, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETACTIVEATTRIB)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVESUBROUTINENAME)(GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVESUBROUTINEUNIFORMNAME)(GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVESUBROUTINEUNIFORMIV)(GLuint  program, GLenum  shadertype, GLuint  index, GLenum  pname, GLint * values);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORM)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMBLOCKNAME)(GLuint  program, GLuint  uniformBlockIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformBlockName);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMBLOCKIV)(GLuint  program, GLuint  uniformBlockIndex, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMNAME)(GLuint  program, GLuint  uniformIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformName);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMSIV)(GLuint  program, GLsizei  uniformCount, const GLuint * uniformIndices, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETATTACHEDSHADERS)(GLuint  program, GLsizei  maxCount, GLsizei * count, GLuint * shaders);
// typedef GLint  (APIENTRYP GPGETATTRIBLOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETBOOLEANI_V)(GLenum  target, GLuint  index, GLboolean * data);
// typedef void  (APIENTRYP GPGETBOOLEANV)(GLenum  pname, GLboolean * data);
// typedef void  (APIENTRYP GPGETBUFFERPARAMETERI64V)(GLenum  target, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETBUFFERPOINTERV)(GLenum  target, GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETBUFFERSUBDATA)(GLenum  target, GLintptr  offset, GLsizeiptr  size, void * data);
// typedef void  (APIENTRYP GPGETCOMPRESSEDTEXIMAGE)(GLenum  target, GLint  level, void * img);
// typedef void  (APIENTRYP GPGETCOMPRESSEDTEXTUREIMAGE)(GLuint  texture, GLint  level, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETCOMPRESSEDTEXTURESUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLsizei  bufSize, void * pixels);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOG)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOGARB)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOGKHR)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef void  (APIENTRYP GPGETDOUBLEI_V)(GLenum  target, GLuint  index, GLdouble * data);
// typedef void  (APIENTRYP GPGETDOUBLEV)(GLenum  pname, GLdouble * data);
// typedef GLenum  (APIENTRYP GPGETERROR)();
// typedef void  (APIENTRYP GPGETFLOATI_V)(GLenum  target, GLuint  index, GLfloat * data);
// typedef void  (APIENTRYP GPGETFLOATV)(GLenum  pname, GLfloat * data);
// typedef GLint  (APIENTRYP GPGETFRAGDATAINDEX)(GLuint  program, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETFRAGDATALOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETFRAMEBUFFERATTACHMENTPARAMETERIV)(GLenum  target, GLenum  attachment, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETFRAMEBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUS)();
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUSARB)();
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUSKHR)();
// typedef GLuint64  (APIENTRYP GPGETIMAGEHANDLEARB)(GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  format);
// typedef void  (APIENTRYP GPGETINTEGER64I_V)(GLenum  target, GLuint  index, GLint64 * data);
// typedef void  (APIENTRYP GPGETINTEGER64V)(GLenum  pname, GLint64 * data);
// typedef void  (APIENTRYP GPGETINTEGERI_V)(GLenum  target, GLuint  index, GLint * data);
// typedef void  (APIENTRYP GPGETINTEGERV)(GLenum  pname, GLint * data);
// typedef void  (APIENTRYP GPGETINTERNALFORMATI64V)(GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint64 * params);
// typedef void  (APIENTRYP GPGETINTERNALFORMATIV)(GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETMULTISAMPLEFV)(GLenum  pname, GLuint  index, GLfloat * val);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERPARAMETERI64V)(GLuint  buffer, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERPARAMETERIV)(GLuint  buffer, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERPOINTERV)(GLuint  buffer, GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERSUBDATA)(GLuint  buffer, GLintptr  offset, GLsizei  size, void * data);
// typedef void  (APIENTRYP GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV)(GLuint  framebuffer, GLenum  attachment, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNAMEDFRAMEBUFFERPARAMETERIV)(GLuint  framebuffer, GLenum  pname, GLint * param);
// typedef void  (APIENTRYP GPGETNAMEDRENDERBUFFERPARAMETERIV)(GLuint  renderbuffer, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNAMEDSTRINGARB)(GLint  namelen, const GLchar * name, GLsizei  bufSize, GLint * stringlen, GLchar * string);
// typedef void  (APIENTRYP GPGETNAMEDSTRINGIVARB)(GLint  namelen, const GLchar * name, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETOBJECTLABEL)(GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTLABELKHR)(GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTPTRLABEL)(const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTPTRLABELKHR)(const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETPOINTERV)(GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETPOINTERVKHR)(GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETPROGRAMBINARY)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary);
// typedef void  (APIENTRYP GPGETPROGRAMINFOLOG)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMINTERFACEIV)(GLuint  program, GLenum  programInterface, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEINFOLOG)(GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEIV)(GLuint  pipeline, GLenum  pname, GLint * params);
// typedef GLuint  (APIENTRYP GPGETPROGRAMRESOURCEINDEX)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETPROGRAMRESOURCELOCATION)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETPROGRAMRESOURCELOCATIONINDEX)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef void  (APIENTRYP GPGETPROGRAMRESOURCENAME)(GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  bufSize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETPROGRAMRESOURCEIV)(GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMSTAGEIV)(GLuint  program, GLenum  shadertype, GLenum  pname, GLint * values);
// typedef void  (APIENTRYP GPGETPROGRAMIV)(GLuint  program, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYINDEXEDIV)(GLenum  target, GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTI64V)(GLuint  id, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTIV)(GLuint  id, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTUI64V)(GLuint  id, GLenum  pname, GLuint64 * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTUIV)(GLuint  id, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETQUERYIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETRENDERBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIIV)(GLuint  sampler, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIUIV)(GLuint  sampler, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERFV)(GLuint  sampler, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIV)(GLuint  sampler, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSHADERINFOLOG)(GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETSHADERPRECISIONFORMAT)(GLenum  shadertype, GLenum  precisiontype, GLint * range, GLint * precision);
// typedef void  (APIENTRYP GPGETSHADERSOURCE)(GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * source);
// typedef void  (APIENTRYP GPGETSHADERIV)(GLuint  shader, GLenum  pname, GLint * params);
// typedef const GLubyte * (APIENTRYP GPGETSTRING)(GLenum  name);
// typedef const GLubyte * (APIENTRYP GPGETSTRINGI)(GLenum  name, GLuint  index);
// typedef GLuint  (APIENTRYP GPGETSUBROUTINEINDEX)(GLuint  program, GLenum  shadertype, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETSUBROUTINEUNIFORMLOCATION)(GLuint  program, GLenum  shadertype, const GLchar * name);
// typedef void  (APIENTRYP GPGETSYNCIV)(GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values);
// typedef void  (APIENTRYP GPGETTEXIMAGE)(GLenum  target, GLint  level, GLenum  format, GLenum  type, void * pixels);
// typedef void  (APIENTRYP GPGETTEXLEVELPARAMETERFV)(GLenum  target, GLint  level, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXLEVELPARAMETERIV)(GLenum  target, GLint  level, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIUIV)(GLenum  target, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERFV)(GLenum  target, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef GLuint64  (APIENTRYP GPGETTEXTUREHANDLEARB)(GLuint  texture);
// typedef void  (APIENTRYP GPGETTEXTUREIMAGE)(GLuint  texture, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETTEXTURELEVELPARAMETERFV)(GLuint  texture, GLint  level, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXTURELEVELPARAMETERIV)(GLuint  texture, GLint  level, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERIIV)(GLuint  texture, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERIUIV)(GLuint  texture, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERFV)(GLuint  texture, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERIV)(GLuint  texture, GLenum  pname, GLint * params);
// typedef GLuint64  (APIENTRYP GPGETTEXTURESAMPLERHANDLEARB)(GLuint  texture, GLuint  sampler);
// typedef void  (APIENTRYP GPGETTEXTURESUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKVARYING)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLsizei * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKI64_V)(GLuint  xfb, GLenum  pname, GLuint  index, GLint64 * param);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKI_V)(GLuint  xfb, GLenum  pname, GLuint  index, GLint * param);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKIV)(GLuint  xfb, GLenum  pname, GLint * param);
// typedef GLuint  (APIENTRYP GPGETUNIFORMBLOCKINDEX)(GLuint  program, const GLchar * uniformBlockName);
// typedef void  (APIENTRYP GPGETUNIFORMINDICES)(GLuint  program, GLsizei  uniformCount, const GLchar *const* uniformNames, GLuint * uniformIndices);
// typedef GLint  (APIENTRYP GPGETUNIFORMLOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETUNIFORMSUBROUTINEUIV)(GLenum  shadertype, GLint  location, GLuint * params);
// typedef void  (APIENTRYP GPGETUNIFORMDV)(GLuint  program, GLint  location, GLdouble * params);
// typedef void  (APIENTRYP GPGETUNIFORMFV)(GLuint  program, GLint  location, GLfloat * params);
// typedef void  (APIENTRYP GPGETUNIFORMIV)(GLuint  program, GLint  location, GLint * params);
// typedef void  (APIENTRYP GPGETUNIFORMUIV)(GLuint  program, GLint  location, GLuint * params);
// typedef void  (APIENTRYP GPGETVERTEXARRAYINDEXED64IV)(GLuint  vaobj, GLuint  index, GLenum  pname, GLint64 * param);
// typedef void  (APIENTRYP GPGETVERTEXARRAYINDEXEDIV)(GLuint  vaobj, GLuint  index, GLenum  pname, GLint * param);
// typedef void  (APIENTRYP GPGETVERTEXARRAYIV)(GLuint  vaobj, GLenum  pname, GLint * param);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIIV)(GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIUIV)(GLuint  index, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBLDV)(GLuint  index, GLenum  pname, GLdouble * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBLUI64VARB)(GLuint  index, GLenum  pname, GLuint64EXT * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBPOINTERV)(GLuint  index, GLenum  pname, void ** pointer);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBDV)(GLuint  index, GLenum  pname, GLdouble * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBFV)(GLuint  index, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIV)(GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNCOMPRESSEDTEXIMAGE)(GLenum  target, GLint  lod, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETNCOMPRESSEDTEXIMAGEARB)(GLenum  target, GLint  lod, GLsizei  bufSize, void * img);
// typedef void  (APIENTRYP GPGETNTEXIMAGE)(GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETNTEXIMAGEARB)(GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * img);
// typedef void  (APIENTRYP GPGETNUNIFORMDV)(GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params);
// typedef void  (APIENTRYP GPGETNUNIFORMDVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFV)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIV)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIV)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPHINT)(GLenum  target, GLenum  mode);
// typedef void  (APIENTRYP GPINVALIDATEBUFFERDATA)(GLuint  buffer);
// typedef void  (APIENTRYP GPINVALIDATEBUFFERSUBDATA)(GLuint  buffer, GLintptr  offset, GLsizeiptr  length);
// typedef void  (APIENTRYP GPINVALIDATEFRAMEBUFFER)(GLenum  target, GLsizei  numAttachments, const GLenum * attachments);
// typedef void  (APIENTRYP GPINVALIDATENAMEDFRAMEBUFFERDATA)(GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments);
// typedef void  (APIENTRYP GPINVALIDATENAMEDFRAMEBUFFERSUBDATA)(GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPINVALIDATESUBFRAMEBUFFER)(GLenum  target, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPINVALIDATETEXIMAGE)(GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPINVALIDATETEXSUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef GLboolean  (APIENTRYP GPISBUFFER)(GLuint  buffer);
// typedef GLboolean  (APIENTRYP GPISENABLED)(GLenum  cap);
// typedef GLboolean  (APIENTRYP GPISENABLEDI)(GLenum  target, GLuint  index);
// typedef GLboolean  (APIENTRYP GPISFRAMEBUFFER)(GLuint  framebuffer);
// typedef GLboolean  (APIENTRYP GPISIMAGEHANDLERESIDENTARB)(GLuint64  handle);
// typedef GLboolean  (APIENTRYP GPISNAMEDSTRINGARB)(GLint  namelen, const GLchar * name);
// typedef GLboolean  (APIENTRYP GPISPROGRAM)(GLuint  program);
// typedef GLboolean  (APIENTRYP GPISPROGRAMPIPELINE)(GLuint  pipeline);
// typedef GLboolean  (APIENTRYP GPISQUERY)(GLuint  id);
// typedef GLboolean  (APIENTRYP GPISRENDERBUFFER)(GLuint  renderbuffer);
// typedef GLboolean  (APIENTRYP GPISSAMPLER)(GLuint  sampler);
// typedef GLboolean  (APIENTRYP GPISSHADER)(GLuint  shader);
// typedef GLboolean  (APIENTRYP GPISSYNC)(GLsync  sync);
// typedef GLboolean  (APIENTRYP GPISTEXTURE)(GLuint  texture);
// typedef GLboolean  (APIENTRYP GPISTEXTUREHANDLERESIDENTARB)(GLuint64  handle);
// typedef GLboolean  (APIENTRYP GPISTRANSFORMFEEDBACK)(GLuint  id);
// typedef GLboolean  (APIENTRYP GPISVERTEXARRAY)(GLuint  array);
// typedef void  (APIENTRYP GPLINEWIDTH)(GLfloat  width);
// typedef void  (APIENTRYP GPLINKPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPLOGICOP)(GLenum  opcode);
// typedef void  (APIENTRYP GPMAKEIMAGEHANDLENONRESIDENTARB)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAKEIMAGEHANDLERESIDENTARB)(GLuint64  handle, GLenum  access);
// typedef void  (APIENTRYP GPMAKETEXTUREHANDLENONRESIDENTARB)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAKETEXTUREHANDLERESIDENTARB)(GLuint64  handle);
// typedef void * (APIENTRYP GPMAPBUFFER)(GLenum  target, GLenum  access);
// typedef void * (APIENTRYP GPMAPBUFFERRANGE)(GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access);
// typedef void * (APIENTRYP GPMAPNAMEDBUFFER)(GLuint  buffer, GLenum  access);
// typedef void * (APIENTRYP GPMAPNAMEDBUFFERRANGE)(GLuint  buffer, GLintptr  offset, GLsizei  length, GLbitfield  access);
// typedef void  (APIENTRYP GPMEMORYBARRIER)(GLbitfield  barriers);
// typedef void  (APIENTRYP GPMEMORYBARRIERBYREGION)(GLbitfield  barriers);
// typedef void  (APIENTRYP GPMINSAMPLESHADING)(GLfloat  value);
// typedef void  (APIENTRYP GPMINSAMPLESHADINGARB)(GLfloat  value);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYS)(GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  drawcount);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYSINDIRECT)(GLenum  mode, const void * indirect, GLsizei  drawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYSINDIRECTCOUNTARB)(GLenum  mode, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTS)(GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSBASEVERTEX)(GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount, const GLint * basevertex);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSINDIRECT)(GLenum  mode, GLenum  type, const void * indirect, GLsizei  drawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSINDIRECTCOUNTARB)(GLenum  mode, GLenum  type, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPNAMEDBUFFERDATA)(GLuint  buffer, GLsizei  size, const void * data, GLenum  usage);
// typedef void  (APIENTRYP GPNAMEDBUFFERPAGECOMMITMENTARB)(GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit);
// typedef void  (APIENTRYP GPNAMEDBUFFERPAGECOMMITMENTEXT)(GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit);
// typedef void  (APIENTRYP GPNAMEDBUFFERSTORAGE)(GLuint  buffer, GLsizei  size, const void * data, GLbitfield  flags);
// typedef void  (APIENTRYP GPNAMEDBUFFERSUBDATA)(GLuint  buffer, GLintptr  offset, GLsizei  size, const void * data);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERDRAWBUFFER)(GLuint  framebuffer, GLenum  buf);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERDRAWBUFFERS)(GLuint  framebuffer, GLsizei  n, const GLenum * bufs);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERPARAMETERI)(GLuint  framebuffer, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERREADBUFFER)(GLuint  framebuffer, GLenum  src);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERRENDERBUFFER)(GLuint  framebuffer, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERTEXTURE)(GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERTEXTURELAYER)(GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer);
// typedef void  (APIENTRYP GPNAMEDRENDERBUFFERSTORAGE)(GLuint  renderbuffer, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE)(GLuint  renderbuffer, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPNAMEDSTRINGARB)(GLenum  type, GLint  namelen, const GLchar * name, GLint  stringlen, const GLchar * string);
// typedef void  (APIENTRYP GPOBJECTLABEL)(GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTLABELKHR)(GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTPTRLABEL)(const void * ptr, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTPTRLABELKHR)(const void * ptr, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPPATCHPARAMETERFV)(GLenum  pname, const GLfloat * values);
// typedef void  (APIENTRYP GPPATCHPARAMETERI)(GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPAUSETRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPPIXELSTOREF)(GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPPIXELSTOREI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPPOINTPARAMETERF)(GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPPOINTPARAMETERFV)(GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPPOINTPARAMETERI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPPOINTPARAMETERIV)(GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPPOINTSIZE)(GLfloat  size);
// typedef void  (APIENTRYP GPPOLYGONMODE)(GLenum  face, GLenum  mode);
// typedef void  (APIENTRYP GPPOLYGONOFFSET)(GLfloat  factor, GLfloat  units);
// typedef void  (APIENTRYP GPPOPDEBUGGROUP)();
// typedef void  (APIENTRYP GPPOPDEBUGGROUPKHR)();
// typedef void  (APIENTRYP GPPRIMITIVERESTARTINDEX)(GLuint  index);
// typedef void  (APIENTRYP GPPROGRAMBINARY)(GLuint  program, GLenum  binaryFormat, const void * binary, GLsizei  length);
// typedef void  (APIENTRYP GPPROGRAMPARAMETERI)(GLuint  program, GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1D)(GLuint  program, GLint  location, GLdouble  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1F)(GLuint  program, GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1I)(GLuint  program, GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UI)(GLuint  program, GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2D)(GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2I)(GLuint  program, GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3D)(GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3I)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4D)(GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2, GLdouble  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4I)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64ARB)(GLuint  program, GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64VARB)(GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X3DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X4DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X2DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X4DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X2DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X3DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROVOKINGVERTEX)(GLenum  mode);
// typedef void  (APIENTRYP GPPUSHDEBUGGROUP)(GLenum  source, GLuint  id, GLsizei  length, const GLchar * message);
// typedef void  (APIENTRYP GPPUSHDEBUGGROUPKHR)(GLenum  source, GLuint  id, GLsizei  length, const GLchar * message);
// typedef void  (APIENTRYP GPQUERYCOUNTER)(GLuint  id, GLenum  target);
// typedef void  (APIENTRYP GPREADBUFFER)(GLenum  src);
// typedef void  (APIENTRYP GPREADPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels);
// typedef void  (APIENTRYP GPREADNPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPREADNPIXELSARB)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPREADNPIXELSKHR)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPRELEASESHADERCOMPILER)();
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGE)(GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRESUMETRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPSAMPLECOVERAGE)(GLfloat  value, GLboolean  invert);
// typedef void  (APIENTRYP GPSAMPLEMASKI)(GLuint  maskNumber, GLbitfield  mask);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIIV)(GLuint  sampler, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIUIV)(GLuint  sampler, GLenum  pname, const GLuint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERF)(GLuint  sampler, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERFV)(GLuint  sampler, GLenum  pname, const GLfloat * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERI)(GLuint  sampler, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIV)(GLuint  sampler, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPSCISSOR)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSORARRAYV)(GLuint  first, GLsizei  count, const GLint * v);
// typedef void  (APIENTRYP GPSCISSORINDEXED)(GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSORINDEXEDV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPSHADERBINARY)(GLsizei  count, const GLuint * shaders, GLenum  binaryformat, const void * binary, GLsizei  length);
// typedef void  (APIENTRYP GPSHADERSOURCE)(GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length);
// typedef void  (APIENTRYP GPSHADERSTORAGEBLOCKBINDING)(GLuint  program, GLuint  storageBlockIndex, GLuint  storageBlockBinding);
// typedef void  (APIENTRYP GPSTENCILFUNC)(GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILFUNCSEPARATE)(GLenum  face, GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILMASK)(GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILMASKSEPARATE)(GLenum  face, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILOP)(GLenum  fail, GLenum  zfail, GLenum  zpass);
// typedef void  (APIENTRYP GPSTENCILOPSEPARATE)(GLenum  face, GLenum  sfail, GLenum  dpfail, GLenum  dppass);
// typedef void  (APIENTRYP GPTEXBUFFER)(GLenum  target, GLenum  internalformat, GLuint  buffer);
// typedef void  (APIENTRYP GPTEXBUFFERRANGE)(GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPTEXIMAGE1D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE2D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE2DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXIMAGE3D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE3DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXPAGECOMMITMENTARB)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  resident);
// typedef void  (APIENTRYP GPTEXPARAMETERIIV)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERIUIV)(GLenum  target, GLenum  pname, const GLuint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERF)(GLenum  target, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPTEXPARAMETERFV)(GLenum  target, GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPTEXPARAMETERI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXPARAMETERIV)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXSTORAGE1D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width);
// typedef void  (APIENTRYP GPTEXSTORAGE2D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPTEXSTORAGE2DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXSTORAGE3D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef void  (APIENTRYP GPTEXSTORAGE3DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXSUBIMAGE1D)(GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTUREBARRIER)();
// typedef void  (APIENTRYP GPTEXTUREBUFFER)(GLuint  texture, GLenum  internalformat, GLuint  buffer);
// typedef void  (APIENTRYP GPTEXTUREBUFFERRANGE)(GLuint  texture, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizei  size);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERIIV)(GLuint  texture, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERIUIV)(GLuint  texture, GLenum  pname, const GLuint * params);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERF)(GLuint  texture, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERFV)(GLuint  texture, GLenum  pname, const GLfloat * param);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERI)(GLuint  texture, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERIV)(GLuint  texture, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPTEXTURESTORAGE1D)(GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width);
// typedef void  (APIENTRYP GPTEXTURESTORAGE2D)(GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPTEXTURESTORAGE2DMULTISAMPLE)(GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXTURESTORAGE3D)(GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef void  (APIENTRYP GPTEXTURESTORAGE3DMULTISAMPLE)(GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXTURESUBIMAGE1D)(GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTURESUBIMAGE2D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTURESUBIMAGE3D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTUREVIEW)(GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKBUFFERBASE)(GLuint  xfb, GLuint  index, GLuint  buffer);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKBUFFERRANGE)(GLuint  xfb, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizei  size);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKVARYINGS)(GLuint  program, GLsizei  count, const GLchar *const* varyings, GLenum  bufferMode);
// typedef void  (APIENTRYP GPUNIFORM1D)(GLint  location, GLdouble  x);
// typedef void  (APIENTRYP GPUNIFORM1DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM1F)(GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPUNIFORM1FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM1I)(GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPUNIFORM1IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM1UI)(GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPUNIFORM1UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM2D)(GLint  location, GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPUNIFORM2DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM2F)(GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPUNIFORM2FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM2I)(GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPUNIFORM2IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM2UI)(GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPUNIFORM2UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM3D)(GLint  location, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPUNIFORM3DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM3F)(GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPUNIFORM3FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM3I)(GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPUNIFORM3IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM3UI)(GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPUNIFORM3UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM4D)(GLint  location, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPUNIFORM4DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM4F)(GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPUNIFORM4FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM4I)(GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPUNIFORM4IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM4UI)(GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPUNIFORM4UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORMBLOCKBINDING)(GLuint  program, GLuint  uniformBlockIndex, GLuint  uniformBlockBinding);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64ARB)(GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64VARB)(GLint  location, GLsizei  count, const GLuint64 * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X3DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X4DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X2DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X4DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X2DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X3DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMSUBROUTINESUIV)(GLenum  shadertype, GLsizei  count, const GLuint * indices);
// typedef GLboolean  (APIENTRYP GPUNMAPBUFFER)(GLenum  target);
// typedef GLboolean  (APIENTRYP GPUNMAPNAMEDBUFFER)(GLuint  buffer);
// typedef void  (APIENTRYP GPUSEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPUSEPROGRAMSTAGES)(GLuint  pipeline, GLbitfield  stages, GLuint  program);
// typedef void  (APIENTRYP GPVALIDATEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPVALIDATEPROGRAMPIPELINE)(GLuint  pipeline);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBBINDING)(GLuint  vaobj, GLuint  attribindex, GLuint  bindingindex);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBFORMAT)(GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBIFORMAT)(GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBLFORMAT)(GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXARRAYBINDINGDIVISOR)(GLuint  vaobj, GLuint  bindingindex, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXARRAYELEMENTBUFFER)(GLuint  vaobj, GLuint  buffer);
// typedef void  (APIENTRYP GPVERTEXARRAYVERTEXBUFFER)(GLuint  vaobj, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride);
// typedef void  (APIENTRYP GPVERTEXARRAYVERTEXBUFFERS)(GLuint  vaobj, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides);
// typedef void  (APIENTRYP GPVERTEXATTRIB1D)(GLuint  index, GLdouble  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB1F)(GLuint  index, GLfloat  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB1S)(GLuint  index, GLshort  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2D)(GLuint  index, GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2F)(GLuint  index, GLfloat  x, GLfloat  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2S)(GLuint  index, GLshort  x, GLshort  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3F)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3S)(GLuint  index, GLshort  x, GLshort  y, GLshort  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NBV)(GLuint  index, const GLbyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NIV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NSV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUB)(GLuint  index, GLubyte  x, GLubyte  y, GLubyte  z, GLubyte  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUBV)(GLuint  index, const GLubyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUSV)(GLuint  index, const GLushort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4BV)(GLuint  index, const GLbyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4F)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4S)(GLuint  index, GLshort  x, GLshort  y, GLshort  z, GLshort  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4UBV)(GLuint  index, const GLubyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4USV)(GLuint  index, const GLushort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBBINDING)(GLuint  attribindex, GLuint  bindingindex);
// typedef void  (APIENTRYP GPVERTEXATTRIBDIVISOR)(GLuint  index, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXATTRIBFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1I)(GLuint  index, GLint  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1UI)(GLuint  index, GLuint  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2I)(GLuint  index, GLint  x, GLint  y);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2UI)(GLuint  index, GLuint  x, GLuint  y);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3I)(GLuint  index, GLint  x, GLint  y, GLint  z);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3UI)(GLuint  index, GLuint  x, GLuint  y, GLuint  z);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4BV)(GLuint  index, const GLbyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4I)(GLuint  index, GLint  x, GLint  y, GLint  z, GLint  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UBV)(GLuint  index, const GLubyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UI)(GLuint  index, GLuint  x, GLuint  y, GLuint  z, GLuint  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4USV)(GLuint  index, const GLushort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBIFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBIPOINTER)(GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1D)(GLuint  index, GLdouble  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1UI64ARB)(GLuint  index, GLuint64EXT  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1UI64VARB)(GLuint  index, const GLuint64EXT * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL2D)(GLuint  index, GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPVERTEXATTRIBL2DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL3D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPVERTEXATTRIBL3DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL4D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBL4DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBLFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBLPOINTER)(GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXATTRIBP1UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP1UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP2UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP2UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP3UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP3UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP4UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP4UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBPOINTER)(GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXBINDINGDIVISOR)(GLuint  bindingindex, GLuint  divisor);
// typedef void  (APIENTRYP GPVIEWPORT)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPVIEWPORTARRAYV)(GLuint  first, GLsizei  count, const GLfloat * v);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDF)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDFV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPWAITSYNC)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// static void  ptrActiveShaderProgram(GPACTIVESHADERPROGRAM fnptr, GLuint  pipeline, GLuint  program) {(*fnptr)(pipeline, program);}
// static void  ptrActiveTexture(GPACTIVETEXTURE fnptr, GLenum  texture) {(*fnptr)(texture);}
// static void  ptrAttachShader(GPATTACHSHADER fnptr, GLuint  program, GLuint  shader) {(*fnptr)(program, shader);}
// static void  ptrBeginConditionalRender(GPBEGINCONDITIONALRENDER fnptr, GLuint  id, GLenum  mode) {(*fnptr)(id, mode);}
// static void  ptrBeginQuery(GPBEGINQUERY fnptr, GLenum  target, GLuint  id) {(*fnptr)(target, id);}
// static void  ptrBeginQueryIndexed(GPBEGINQUERYINDEXED fnptr, GLenum  target, GLuint  index, GLuint  id) {(*fnptr)(target, index, id);}
// static void  ptrBeginTransformFeedback(GPBEGINTRANSFORMFEEDBACK fnptr, GLenum  primitiveMode) {(*fnptr)(primitiveMode);}
// static void  ptrBindAttribLocation(GPBINDATTRIBLOCATION fnptr, GLuint  program, GLuint  index, const GLchar * name) {(*fnptr)(program, index, name);}
// static void  ptrBindBuffer(GPBINDBUFFER fnptr, GLenum  target, GLuint  buffer) {(*fnptr)(target, buffer);}
// static void  ptrBindBufferBase(GPBINDBUFFERBASE fnptr, GLenum  target, GLuint  index, GLuint  buffer) {(*fnptr)(target, index, buffer);}
// static void  ptrBindBufferRange(GPBINDBUFFERRANGE fnptr, GLenum  target, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizeiptr  size) {(*fnptr)(target, index, buffer, offset, size);}
// static void  ptrBindBuffersBase(GPBINDBUFFERSBASE fnptr, GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers) {(*fnptr)(target, first, count, buffers);}
// static void  ptrBindBuffersRange(GPBINDBUFFERSRANGE fnptr, GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizeiptr * sizes) {(*fnptr)(target, first, count, buffers, offsets, sizes);}
// static void  ptrBindFragDataLocation(GPBINDFRAGDATALOCATION fnptr, GLuint  program, GLuint  color, const GLchar * name) {(*fnptr)(program, color, name);}
// static void  ptrBindFragDataLocationIndexed(GPBINDFRAGDATALOCATIONINDEXED fnptr, GLuint  program, GLuint  colorNumber, GLuint  index, const GLchar * name) {(*fnptr)(program, colorNumber, index, name);}
// static void  ptrBindFramebuffer(GPBINDFRAMEBUFFER fnptr, GLenum  target, GLuint  framebuffer) {(*fnptr)(target, framebuffer);}
// static void  ptrBindImageTexture(GPBINDIMAGETEXTURE fnptr, GLuint  unit, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  access, GLenum  format) {(*fnptr)(unit, texture, level, layered, layer, access, format);}
// static void  ptrBindImageTextures(GPBINDIMAGETEXTURES fnptr, GLuint  first, GLsizei  count, const GLuint * textures) {(*fnptr)(first, count, textures);}
// static void  ptrBindProgramPipeline(GPBINDPROGRAMPIPELINE fnptr, GLuint  pipeline) {(*fnptr)(pipeline);}
// static void  ptrBindRenderbuffer(GPBINDRENDERBUFFER fnptr, GLenum  target, GLuint  renderbuffer) {(*fnptr)(target, renderbuffer);}
// static void  ptrBindSampler(GPBINDSAMPLER fnptr, GLuint  unit, GLuint  sampler) {(*fnptr)(unit, sampler);}
// static void  ptrBindSamplers(GPBINDSAMPLERS fnptr, GLuint  first, GLsizei  count, const GLuint * samplers) {(*fnptr)(first, count, samplers);}
// static void  ptrBindTexture(GPBINDTEXTURE fnptr, GLenum  target, GLuint  texture) {(*fnptr)(target, texture);}
// static void  ptrBindTextureUnit(GPBINDTEXTUREUNIT fnptr, GLuint  unit, GLuint  texture) {(*fnptr)(unit, texture);}
// static void  ptrBindTextures(GPBINDTEXTURES fnptr, GLuint  first, GLsizei  count, const GLuint * textures) {(*fnptr)(first, count, textures);}
// static void  ptrBindTransformFeedback(GPBINDTRANSFORMFEEDBACK fnptr, GLenum  target, GLuint  id) {(*fnptr)(target, id);}
// static void  ptrBindVertexArray(GPBINDVERTEXARRAY fnptr, GLuint  array) {(*fnptr)(array);}
// static void  ptrBindVertexBuffer(GPBINDVERTEXBUFFER fnptr, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride) {(*fnptr)(bindingindex, buffer, offset, stride);}
// static void  ptrBindVertexBuffers(GPBINDVERTEXBUFFERS fnptr, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides) {(*fnptr)(first, count, buffers, offsets, strides);}
// static void  ptrBlendColor(GPBLENDCOLOR fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {(*fnptr)(red, green, blue, alpha);}
// static void  ptrBlendEquation(GPBLENDEQUATION fnptr, GLenum  mode) {(*fnptr)(mode);}
// static void  ptrBlendEquationSeparate(GPBLENDEQUATIONSEPARATE fnptr, GLenum  modeRGB, GLenum  modeAlpha) {(*fnptr)(modeRGB, modeAlpha);}
// static void  ptrBlendEquationSeparatei(GPBLENDEQUATIONSEPARATEI fnptr, GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha) {(*fnptr)(buf, modeRGB, modeAlpha);}
// static void  ptrBlendEquationSeparateiARB(GPBLENDEQUATIONSEPARATEIARB fnptr, GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha) {(*fnptr)(buf, modeRGB, modeAlpha);}
// static void  ptrBlendEquationi(GPBLENDEQUATIONI fnptr, GLuint  buf, GLenum  mode) {(*fnptr)(buf, mode);}
// static void  ptrBlendEquationiARB(GPBLENDEQUATIONIARB fnptr, GLuint  buf, GLenum  mode) {(*fnptr)(buf, mode);}
// static void  ptrBlendFunc(GPBLENDFUNC fnptr, GLenum  sfactor, GLenum  dfactor) {(*fnptr)(sfactor, dfactor);}
// static void  ptrBlendFuncSeparate(GPBLENDFUNCSEPARATE fnptr, GLenum  sfactorRGB, GLenum  dfactorRGB, GLenum  sfactorAlpha, GLenum  dfactorAlpha) {(*fnptr)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha);}
// static void  ptrBlendFuncSeparatei(GPBLENDFUNCSEPARATEI fnptr, GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha) {(*fnptr)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);}
// static void  ptrBlendFuncSeparateiARB(GPBLENDFUNCSEPARATEIARB fnptr, GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha) {(*fnptr)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);}
// static void  ptrBlendFunci(GPBLENDFUNCI fnptr, GLuint  buf, GLenum  src, GLenum  dst) {(*fnptr)(buf, src, dst);}
// static void  ptrBlendFunciARB(GPBLENDFUNCIARB fnptr, GLuint  buf, GLenum  src, GLenum  dst) {(*fnptr)(buf, src, dst);}
// static void  ptrBlitFramebuffer(GPBLITFRAMEBUFFER fnptr, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {(*fnptr)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);}
// static void  ptrBlitNamedFramebuffer(GPBLITNAMEDFRAMEBUFFER fnptr, GLuint  readFramebuffer, GLuint  drawFramebuffer, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {(*fnptr)(readFramebuffer, drawFramebuffer, srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);}
// static void  ptrBufferData(GPBUFFERDATA fnptr, GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage) {(*fnptr)(target, size, data, usage);}
// static void  ptrBufferPageCommitmentARB(GPBUFFERPAGECOMMITMENTARB fnptr, GLenum  target, GLintptr  offset, GLsizei  size, GLboolean  commit) {(*fnptr)(target, offset, size, commit);}
// static void  ptrBufferStorage(GPBUFFERSTORAGE fnptr, GLenum  target, GLsizeiptr  size, const void * data, GLbitfield  flags) {(*fnptr)(target, size, data, flags);}
// static void  ptrBufferSubData(GPBUFFERSUBDATA fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data) {(*fnptr)(target, offset, size, data);}
// static GLenum  ptrCheckFramebufferStatus(GPCHECKFRAMEBUFFERSTATUS fnptr, GLenum  target) {return (*fnptr)(target);}
// static GLenum  ptrCheckNamedFramebufferStatus(GPCHECKNAMEDFRAMEBUFFERSTATUS fnptr, GLuint  framebuffer, GLenum  target) {return (*fnptr)(framebuffer, target);}
// static void  ptrClampColor(GPCLAMPCOLOR fnptr, GLenum  target, GLenum  clamp) {(*fnptr)(target, clamp);}
// static void  ptrClear(GPCLEAR fnptr, GLbitfield  mask) {(*fnptr)(mask);}
// static void  ptrClearBufferData(GPCLEARBUFFERDATA fnptr, GLenum  target, GLenum  internalformat, GLenum  format, GLenum  type, const void * data) {(*fnptr)(target, internalformat, format, type, data);}
// static void  ptrClearBufferSubData(GPCLEARBUFFERSUBDATA fnptr, GLenum  target, GLenum  internalformat, GLintptr  offset, GLsizeiptr  size, GLenum  format, GLenum  type, const void * data) {(*fnptr)(target, internalformat, offset, size, format, type, data);}
// static void  ptrClearBufferfi(GPCLEARBUFFERFI fnptr, GLenum  buffer, GLint  drawbuffer, GLfloat  depth, GLint  stencil) {(*fnptr)(buffer, drawbuffer, depth, stencil);}
// static void  ptrClearBufferfv(GPCLEARBUFFERFV fnptr, GLenum  buffer, GLint  drawbuffer, const GLfloat * value) {(*fnptr)(buffer, drawbuffer, value);}
// static void  ptrClearBufferiv(GPCLEARBUFFERIV fnptr, GLenum  buffer, GLint  drawbuffer, const GLint * value) {(*fnptr)(buffer, drawbuffer, value);}
// static void  ptrClearBufferuiv(GPCLEARBUFFERUIV fnptr, GLenum  buffer, GLint  drawbuffer, const GLuint * value) {(*fnptr)(buffer, drawbuffer, value);}
// static void  ptrClearColor(GPCLEARCOLOR fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {(*fnptr)(red, green, blue, alpha);}
// static void  ptrClearDepth(GPCLEARDEPTH fnptr, GLdouble  depth) {(*fnptr)(depth);}
// static void  ptrClearDepthf(GPCLEARDEPTHF fnptr, GLfloat  d) {(*fnptr)(d);}
// static void  ptrClearNamedBufferData(GPCLEARNAMEDBUFFERDATA fnptr, GLuint  buffer, GLenum  internalformat, GLenum  format, GLenum  type, const void * data) {(*fnptr)(buffer, internalformat, format, type, data);}
// static void  ptrClearNamedBufferSubData(GPCLEARNAMEDBUFFERSUBDATA fnptr, GLuint  buffer, GLenum  internalformat, GLintptr  offset, GLsizei  size, GLenum  format, GLenum  type, const void * data) {(*fnptr)(buffer, internalformat, offset, size, format, type, data);}
// static void  ptrClearNamedFramebufferfi(GPCLEARNAMEDFRAMEBUFFERFI fnptr, GLuint  framebuffer, GLenum  buffer, const GLfloat  depth, GLint  stencil) {(*fnptr)(framebuffer, buffer, depth, stencil);}
// static void  ptrClearNamedFramebufferfv(GPCLEARNAMEDFRAMEBUFFERFV fnptr, GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLfloat * value) {(*fnptr)(framebuffer, buffer, drawbuffer, value);}
// static void  ptrClearNamedFramebufferiv(GPCLEARNAMEDFRAMEBUFFERIV fnptr, GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLint * value) {(*fnptr)(framebuffer, buffer, drawbuffer, value);}
// static void  ptrClearNamedFramebufferuiv(GPCLEARNAMEDFRAMEBUFFERUIV fnptr, GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLuint * value) {(*fnptr)(framebuffer, buffer, drawbuffer, value);}
// static void  ptrClearStencil(GPCLEARSTENCIL fnptr, GLint  s) {(*fnptr)(s);}
// static void  ptrClearTexImage(GPCLEARTEXIMAGE fnptr, GLuint  texture, GLint  level, GLenum  format, GLenum  type, const void * data) {(*fnptr)(texture, level, format, type, data);}
// static void  ptrClearTexSubImage(GPCLEARTEXSUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * data) {(*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, data);}
// static GLenum  ptrClientWaitSync(GPCLIENTWAITSYNC fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {return (*fnptr)(sync, flags, timeout);}
// static void  ptrClipControl(GPCLIPCONTROL fnptr, GLenum  origin, GLenum  depth) {(*fnptr)(origin, depth);}
// static void  ptrColorMask(GPCOLORMASK fnptr, GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha) {(*fnptr)(red, green, blue, alpha);}
// static void  ptrColorMaski(GPCOLORMASKI fnptr, GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a) {(*fnptr)(index, r, g, b, a);}
// static void  ptrCompileShader(GPCOMPILESHADER fnptr, GLuint  shader) {(*fnptr)(shader);}
// static void  ptrCompileShaderIncludeARB(GPCOMPILESHADERINCLUDEARB fnptr, GLuint  shader, GLsizei  count, const GLchar *const* path, const GLint * length) {(*fnptr)(shader, count, path, length);}
// static void  ptrCompressedTexImage1D(GPCOMPRESSEDTEXIMAGE1D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLint  border, GLsizei  imageSize, const void * data) {(*fnptr)(target, level, internalformat, width, border, imageSize, data);}
// static void  ptrCompressedTexImage2D(GPCOMPRESSEDTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLsizei  imageSize, const void * data) {(*fnptr)(target, level, internalformat, width, height, border, imageSize, data);}
// static void  ptrCompressedTexImage3D(GPCOMPRESSEDTEXIMAGE3D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data) {(*fnptr)(target, level, internalformat, width, height, depth, border, imageSize, data);}
// static void  ptrCompressedTexSubImage1D(GPCOMPRESSEDTEXSUBIMAGE1D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data) {(*fnptr)(target, level, xoffset, width, format, imageSize, data);}
// static void  ptrCompressedTexSubImage2D(GPCOMPRESSEDTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data) {(*fnptr)(target, level, xoffset, yoffset, width, height, format, imageSize, data);}
// static void  ptrCompressedTexSubImage3D(GPCOMPRESSEDTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data) {(*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);}
// static void  ptrCompressedTextureSubImage1D(GPCOMPRESSEDTEXTURESUBIMAGE1D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data) {(*fnptr)(texture, level, xoffset, width, format, imageSize, data);}
// static void  ptrCompressedTextureSubImage2D(GPCOMPRESSEDTEXTURESUBIMAGE2D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data) {(*fnptr)(texture, level, xoffset, yoffset, width, height, format, imageSize, data);}
// static void  ptrCompressedTextureSubImage3D(GPCOMPRESSEDTEXTURESUBIMAGE3D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data) {(*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);}
// static void  ptrCopyBufferSubData(GPCOPYBUFFERSUBDATA fnptr, GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size) {(*fnptr)(readTarget, writeTarget, readOffset, writeOffset, size);}
// static void  ptrCopyImageSubData(GPCOPYIMAGESUBDATA fnptr, GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth) {(*fnptr)(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);}
// static void  ptrCopyNamedBufferSubData(GPCOPYNAMEDBUFFERSUBDATA fnptr, GLuint  readBuffer, GLuint  writeBuffer, GLintptr  readOffset, GLintptr  writeOffset, GLsizei  size) {(*fnptr)(readBuffer, writeBuffer, readOffset, writeOffset, size);}
// static void  ptrCopyTexImage1D(GPCOPYTEXIMAGE1D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLint  border) {(*fnptr)(target, level, internalformat, x, y, width, border);}
// static void  ptrCopyTexImage2D(GPCOPYTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLint  border) {(*fnptr)(target, level, internalformat, x, y, width, height, border);}
// static void  ptrCopyTexSubImage1D(GPCOPYTEXSUBIMAGE1D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width) {(*fnptr)(target, level, xoffset, x, y, width);}
// static void  ptrCopyTexSubImage2D(GPCOPYTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {(*fnptr)(target, level, xoffset, yoffset, x, y, width, height);}
// static void  ptrCopyTexSubImage3D(GPCOPYTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {(*fnptr)(target, level, xoffset, yoffset, zoffset, x, y, width, height);}
// static void  ptrCopyTextureSubImage1D(GPCOPYTEXTURESUBIMAGE1D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width) {(*fnptr)(texture, level, xoffset, x, y, width);}
// static void  ptrCopyTextureSubImage2D(GPCOPYTEXTURESUBIMAGE2D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {(*fnptr)(texture, level, xoffset, yoffset, x, y, width, height);}
// static void  ptrCopyTextureSubImage3D(GPCOPYTEXTURESUBIMAGE3D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {(*fnptr)(texture, level, xoffset, yoffset, zoffset, x, y, width, height);}
// static void  ptrCreateBuffers(GPCREATEBUFFERS fnptr, GLsizei  n, GLuint * buffers) {(*fnptr)(n, buffers);}
// static void  ptrCreateFramebuffers(GPCREATEFRAMEBUFFERS fnptr, GLsizei  n, GLuint * framebuffers) {(*fnptr)(n, framebuffers);}
// static GLuint  ptrCreateProgram(GPCREATEPROGRAM fnptr) {return (*fnptr)();}
// static void  ptrCreateProgramPipelines(GPCREATEPROGRAMPIPELINES fnptr, GLsizei  n, GLuint * pipelines) {(*fnptr)(n, pipelines);}
// static void  ptrCreateQueries(GPCREATEQUERIES fnptr, GLenum  target, GLsizei  n, GLuint * ids) {(*fnptr)(target, n, ids);}
// static void  ptrCreateRenderbuffers(GPCREATERENDERBUFFERS fnptr, GLsizei  n, GLuint * renderbuffers) {(*fnptr)(n, renderbuffers);}
// static void  ptrCreateSamplers(GPCREATESAMPLERS fnptr, GLsizei  n, GLuint * samplers) {(*fnptr)(n, samplers);}
// static GLuint  ptrCreateShader(GPCREATESHADER fnptr, GLenum  type) {return (*fnptr)(type);}
// static GLuint  ptrCreateShaderProgramv(GPCREATESHADERPROGRAMV fnptr, GLenum  type, GLsizei  count, const GLchar *const* strings) {return (*fnptr)(type, count, strings);}
// static GLsync  ptrCreateSyncFromCLeventARB(GPCREATESYNCFROMCLEVENTARB fnptr, struct _cl_context * context, struct _cl_event * event, GLbitfield  flags) {return (*fnptr)(context, event, flags);}
// static void  ptrCreateTextures(GPCREATETEXTURES fnptr, GLenum  target, GLsizei  n, GLuint * textures) {(*fnptr)(target, n, textures);}
// static void  ptrCreateTransformFeedbacks(GPCREATETRANSFORMFEEDBACKS fnptr, GLsizei  n, GLuint * ids) {(*fnptr)(n, ids);}
// static void  ptrCreateVertexArrays(GPCREATEVERTEXARRAYS fnptr, GLsizei  n, GLuint * arrays) {(*fnptr)(n, arrays);}
// static void  ptrCullFace(GPCULLFACE fnptr, GLenum  mode) {(*fnptr)(mode);}
// static void  ptrDebugMessageCallback(GPDEBUGMESSAGECALLBACK fnptr, GLDEBUGPROC  callback, const void * userParam) {(*fnptr)(cDebugCallback, userParam);}
// static void  ptrDebugMessageCallbackARB(GPDEBUGMESSAGECALLBACKARB fnptr, GLDEBUGPROCARB  callback, const void * userParam) {(*fnptr)(cDebugCallback, userParam);}
// static void  ptrDebugMessageCallbackKHR(GPDEBUGMESSAGECALLBACKKHR fnptr, GLDEBUGPROCKHR  callback, const void * userParam) {(*fnptr)(cDebugCallback, userParam);}
// static void  ptrDebugMessageControl(GPDEBUGMESSAGECONTROL fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {(*fnptr)(source, type, severity, count, ids, enabled);}
// static void  ptrDebugMessageControlARB(GPDEBUGMESSAGECONTROLARB fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {(*fnptr)(source, type, severity, count, ids, enabled);}
// static void  ptrDebugMessageControlKHR(GPDEBUGMESSAGECONTROLKHR fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {(*fnptr)(source, type, severity, count, ids, enabled);}
// static void  ptrDebugMessageInsert(GPDEBUGMESSAGEINSERT fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {(*fnptr)(source, type, id, severity, length, buf);}
// static void  ptrDebugMessageInsertARB(GPDEBUGMESSAGEINSERTARB fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {(*fnptr)(source, type, id, severity, length, buf);}
// static void  ptrDebugMessageInsertKHR(GPDEBUGMESSAGEINSERTKHR fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {(*fnptr)(source, type, id, severity, length, buf);}
// static void  ptrDeleteBuffers(GPDELETEBUFFERS fnptr, GLsizei  n, const GLuint * buffers) {(*fnptr)(n, buffers);}
// static void  ptrDeleteFramebuffers(GPDELETEFRAMEBUFFERS fnptr, GLsizei  n, const GLuint * framebuffers) {(*fnptr)(n, framebuffers);}
// static void  ptrDeleteNamedStringARB(GPDELETENAMEDSTRINGARB fnptr, GLint  namelen, const GLchar * name) {(*fnptr)(namelen, name);}
// static void  ptrDeleteProgram(GPDELETEPROGRAM fnptr, GLuint  program) {(*fnptr)(program);}
// static void  ptrDeleteProgramPipelines(GPDELETEPROGRAMPIPELINES fnptr, GLsizei  n, const GLuint * pipelines) {(*fnptr)(n, pipelines);}
// static void  ptrDeleteQueries(GPDELETEQUERIES fnptr, GLsizei  n, const GLuint * ids) {(*fnptr)(n, ids);}
// static void  ptrDeleteRenderbuffers(GPDELETERENDERBUFFERS fnptr, GLsizei  n, const GLuint * renderbuffers) {(*fnptr)(n, renderbuffers);}
// static void  ptrDeleteSamplers(GPDELETESAMPLERS fnptr, GLsizei  count, const GLuint * samplers) {(*fnptr)(count, samplers);}
// static void  ptrDeleteShader(GPDELETESHADER fnptr, GLuint  shader) {(*fnptr)(shader);}
// static void  ptrDeleteSync(GPDELETESYNC fnptr, GLsync  sync) {(*fnptr)(sync);}
// static void  ptrDeleteTextures(GPDELETETEXTURES fnptr, GLsizei  n, const GLuint * textures) {(*fnptr)(n, textures);}
// static void  ptrDeleteTransformFeedbacks(GPDELETETRANSFORMFEEDBACKS fnptr, GLsizei  n, const GLuint * ids) {(*fnptr)(n, ids);}
// static void  ptrDeleteVertexArrays(GPDELETEVERTEXARRAYS fnptr, GLsizei  n, const GLuint * arrays) {(*fnptr)(n, arrays);}
// static void  ptrDepthFunc(GPDEPTHFUNC fnptr, GLenum  func) {(*fnptr)(func);}
// static void  ptrDepthMask(GPDEPTHMASK fnptr, GLboolean  flag) {(*fnptr)(flag);}
// static void  ptrDepthRange(GPDEPTHRANGE fnptr, GLdouble  xnear, GLdouble  xfar) {(*fnptr)(xnear, xfar);}
// static void  ptrDepthRangeArrayv(GPDEPTHRANGEARRAYV fnptr, GLuint  first, GLsizei  count, const GLdouble * v) {(*fnptr)(first, count, v);}
// static void  ptrDepthRangeIndexed(GPDEPTHRANGEINDEXED fnptr, GLuint  index, GLdouble  n, GLdouble  f) {(*fnptr)(index, n, f);}
// static void  ptrDepthRangef(GPDEPTHRANGEF fnptr, GLfloat  n, GLfloat  f) {(*fnptr)(n, f);}
// static void  ptrDetachShader(GPDETACHSHADER fnptr, GLuint  program, GLuint  shader) {(*fnptr)(program, shader);}
// static void  ptrDisable(GPDISABLE fnptr, GLenum  cap) {(*fnptr)(cap);}
// static void  ptrDisableVertexArrayAttrib(GPDISABLEVERTEXARRAYATTRIB fnptr, GLuint  vaobj, GLuint  index) {(*fnptr)(vaobj, index);}
// static void  ptrDisableVertexAttribArray(GPDISABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {(*fnptr)(index);}
// static void  ptrDisablei(GPDISABLEI fnptr, GLenum  target, GLuint  index) {(*fnptr)(target, index);}
// static void  ptrDispatchCompute(GPDISPATCHCOMPUTE fnptr, GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z) {(*fnptr)(num_groups_x, num_groups_y, num_groups_z);}
// static void  ptrDispatchComputeGroupSizeARB(GPDISPATCHCOMPUTEGROUPSIZEARB fnptr, GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z, GLuint  group_size_x, GLuint  group_size_y, GLuint  group_size_z) {(*fnptr)(num_groups_x, num_groups_y, num_groups_z, group_size_x, group_size_y, group_size_z);}
// static void  ptrDispatchComputeIndirect(GPDISPATCHCOMPUTEINDIRECT fnptr, GLintptr  indirect) {(*fnptr)(indirect);}
// static void  ptrDrawArrays(GPDRAWARRAYS fnptr, GLenum  mode, GLint  first, GLsizei  count) {(*fnptr)(mode, first, count);}
// static void  ptrDrawArraysIndirect(GPDRAWARRAYSINDIRECT fnptr, GLenum  mode, const void * indirect) {(*fnptr)(mode, indirect);}
// static void  ptrDrawArraysInstanced(GPDRAWARRAYSINSTANCED fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount) {(*fnptr)(mode, first, count, instancecount);}
// static void  ptrDrawArraysInstancedBaseInstance(GPDRAWARRAYSINSTANCEDBASEINSTANCE fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount, GLuint  baseinstance) {(*fnptr)(mode, first, count, instancecount, baseinstance);}
// static void  ptrDrawBuffer(GPDRAWBUFFER fnptr, GLenum  buf) {(*fnptr)(buf);}
// static void  ptrDrawBuffers(GPDRAWBUFFERS fnptr, GLsizei  n, const GLenum * bufs) {(*fnptr)(n, bufs);}
// static void  ptrDrawElements(GPDRAWELEMENTS fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices) {(*fnptr)(mode, count, type, indices);}
// static void  ptrDrawElementsBaseVertex(GPDRAWELEMENTSBASEVERTEX fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {(*fnptr)(mode, count, type, indices, basevertex);}
// static void  ptrDrawElementsIndirect(GPDRAWELEMENTSINDIRECT fnptr, GLenum  mode, GLenum  type, const void * indirect) {(*fnptr)(mode, type, indirect);}
// static void  ptrDrawElementsInstanced(GPDRAWELEMENTSINSTANCED fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount) {(*fnptr)(mode, count, type, indices, instancecount);}
// static void  ptrDrawElementsInstancedBaseInstance(GPDRAWELEMENTSINSTANCEDBASEINSTANCE fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLuint  baseinstance) {(*fnptr)(mode, count, type, indices, instancecount, baseinstance);}
// static void  ptrDrawElementsInstancedBaseVertex(GPDRAWELEMENTSINSTANCEDBASEVERTEX fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex) {(*fnptr)(mode, count, type, indices, instancecount, basevertex);}
// static void  ptrDrawElementsInstancedBaseVertexBaseInstance(GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex, GLuint  baseinstance) {(*fnptr)(mode, count, type, indices, instancecount, basevertex, baseinstance);}
// static void  ptrDrawRangeElements(GPDRAWRANGEELEMENTS fnptr, GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices) {(*fnptr)(mode, start, end, count, type, indices);}
// static void  ptrDrawRangeElementsBaseVertex(GPDRAWRANGEELEMENTSBASEVERTEX fnptr, GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {(*fnptr)(mode, start, end, count, type, indices, basevertex);}
// static void  ptrDrawTransformFeedback(GPDRAWTRANSFORMFEEDBACK fnptr, GLenum  mode, GLuint  id) {(*fnptr)(mode, id);}
// static void  ptrDrawTransformFeedbackInstanced(GPDRAWTRANSFORMFEEDBACKINSTANCED fnptr, GLenum  mode, GLuint  id, GLsizei  instancecount) {(*fnptr)(mode, id, instancecount);}
// static void  ptrDrawTransformFeedbackStream(GPDRAWTRANSFORMFEEDBACKSTREAM fnptr, GLenum  mode, GLuint  id, GLuint  stream) {(*fnptr)(mode, id, stream);}
// static void  ptrDrawTransformFeedbackStreamInstanced(GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED fnptr, GLenum  mode, GLuint  id, GLuint  stream, GLsizei  instancecount) {(*fnptr)(mode, id, stream, instancecount);}
// static void  ptrEnable(GPENABLE fnptr, GLenum  cap) {(*fnptr)(cap);}
// static void  ptrEnableVertexArrayAttrib(GPENABLEVERTEXARRAYATTRIB fnptr, GLuint  vaobj, GLuint  index) {(*fnptr)(vaobj, index);}
// static void  ptrEnableVertexAttribArray(GPENABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {(*fnptr)(index);}
// static void  ptrEnablei(GPENABLEI fnptr, GLenum  target, GLuint  index) {(*fnptr)(target, index);}
// static void  ptrEndConditionalRender(GPENDCONDITIONALRENDER fnptr) {(*fnptr)();}
// static void  ptrEndQuery(GPENDQUERY fnptr, GLenum  target) {(*fnptr)(target);}
// static void  ptrEndQueryIndexed(GPENDQUERYINDEXED fnptr, GLenum  target, GLuint  index) {(*fnptr)(target, index);}
// static void  ptrEndTransformFeedback(GPENDTRANSFORMFEEDBACK fnptr) {(*fnptr)();}
// static GLsync  ptrFenceSync(GPFENCESYNC fnptr, GLenum  condition, GLbitfield  flags) {return (*fnptr)(condition, flags);}
// static void  ptrFinish(GPFINISH fnptr) {(*fnptr)();}
// static void  ptrFlush(GPFLUSH fnptr) {(*fnptr)();}
// static void  ptrFlushMappedBufferRange(GPFLUSHMAPPEDBUFFERRANGE fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length) {(*fnptr)(target, offset, length);}
// static void  ptrFlushMappedNamedBufferRange(GPFLUSHMAPPEDNAMEDBUFFERRANGE fnptr, GLuint  buffer, GLintptr  offset, GLsizei  length) {(*fnptr)(buffer, offset, length);}
// static void  ptrFramebufferParameteri(GPFRAMEBUFFERPARAMETERI fnptr, GLenum  target, GLenum  pname, GLint  param) {(*fnptr)(target, pname, param);}
// static void  ptrFramebufferRenderbuffer(GPFRAMEBUFFERRENDERBUFFER fnptr, GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer) {(*fnptr)(target, attachment, renderbuffertarget, renderbuffer);}
// static void  ptrFramebufferTexture(GPFRAMEBUFFERTEXTURE fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level) {(*fnptr)(target, attachment, texture, level);}
// static void  ptrFramebufferTexture1D(GPFRAMEBUFFERTEXTURE1D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level) {(*fnptr)(target, attachment, textarget, texture, level);}
// static void  ptrFramebufferTexture2D(GPFRAMEBUFFERTEXTURE2D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level) {(*fnptr)(target, attachment, textarget, texture, level);}
// static void  ptrFramebufferTexture3D(GPFRAMEBUFFERTEXTURE3D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  zoffset) {(*fnptr)(target, attachment, textarget, texture, level, zoffset);}
// static void  ptrFramebufferTextureLayer(GPFRAMEBUFFERTEXTURELAYER fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer) {(*fnptr)(target, attachment, texture, level, layer);}
// static void  ptrFrontFace(GPFRONTFACE fnptr, GLenum  mode) {(*fnptr)(mode);}
// static void  ptrGenBuffers(GPGENBUFFERS fnptr, GLsizei  n, GLuint * buffers) {(*fnptr)(n, buffers);}
// static void  ptrGenFramebuffers(GPGENFRAMEBUFFERS fnptr, GLsizei  n, GLuint * framebuffers) {(*fnptr)(n, framebuffers);}
// static void  ptrGenProgramPipelines(GPGENPROGRAMPIPELINES fnptr, GLsizei  n, GLuint * pipelines) {(*fnptr)(n, pipelines);}
// static void  ptrGenQueries(GPGENQUERIES fnptr, GLsizei  n, GLuint * ids) {(*fnptr)(n, ids);}
// static void  ptrGenRenderbuffers(GPGENRENDERBUFFERS fnptr, GLsizei  n, GLuint * renderbuffers) {(*fnptr)(n, renderbuffers);}
// static void  ptrGenSamplers(GPGENSAMPLERS fnptr, GLsizei  count, GLuint * samplers) {(*fnptr)(count, samplers);}
// static void  ptrGenTextures(GPGENTEXTURES fnptr, GLsizei  n, GLuint * textures) {(*fnptr)(n, textures);}
// static void  ptrGenTransformFeedbacks(GPGENTRANSFORMFEEDBACKS fnptr, GLsizei  n, GLuint * ids) {(*fnptr)(n, ids);}
// static void  ptrGenVertexArrays(GPGENVERTEXARRAYS fnptr, GLsizei  n, GLuint * arrays) {(*fnptr)(n, arrays);}
// static void  ptrGenerateMipmap(GPGENERATEMIPMAP fnptr, GLenum  target) {(*fnptr)(target);}
// static void  ptrGenerateTextureMipmap(GPGENERATETEXTUREMIPMAP fnptr, GLuint  texture) {(*fnptr)(texture);}
// static void  ptrGetActiveAtomicCounterBufferiv(GPGETACTIVEATOMICCOUNTERBUFFERIV fnptr, GLuint  program, GLuint  bufferIndex, GLenum  pname, GLint * params) {(*fnptr)(program, bufferIndex, pname, params);}
// static void  ptrGetActiveAttrib(GPGETACTIVEATTRIB fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name) {(*fnptr)(program, index, bufSize, length, size, type, name);}
// static void  ptrGetActiveSubroutineName(GPGETACTIVESUBROUTINENAME fnptr, GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name) {(*fnptr)(program, shadertype, index, bufsize, length, name);}
// static void  ptrGetActiveSubroutineUniformName(GPGETACTIVESUBROUTINEUNIFORMNAME fnptr, GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name) {(*fnptr)(program, shadertype, index, bufsize, length, name);}
// static void  ptrGetActiveSubroutineUniformiv(GPGETACTIVESUBROUTINEUNIFORMIV fnptr, GLuint  program, GLenum  shadertype, GLuint  index, GLenum  pname, GLint * values) {(*fnptr)(program, shadertype, index, pname, values);}
// static void  ptrGetActiveUniform(GPGETACTIVEUNIFORM fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name) {(*fnptr)(program, index, bufSize, length, size, type, name);}
// static void  ptrGetActiveUniformBlockName(GPGETACTIVEUNIFORMBLOCKNAME fnptr, GLuint  program, GLuint  uniformBlockIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformBlockName) {(*fnptr)(program, uniformBlockIndex, bufSize, length, uniformBlockName);}
// static void  ptrGetActiveUniformBlockiv(GPGETACTIVEUNIFORMBLOCKIV fnptr, GLuint  program, GLuint  uniformBlockIndex, GLenum  pname, GLint * params) {(*fnptr)(program, uniformBlockIndex, pname, params);}
// static void  ptrGetActiveUniformName(GPGETACTIVEUNIFORMNAME fnptr, GLuint  program, GLuint  uniformIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformName) {(*fnptr)(program, uniformIndex, bufSize, length, uniformName);}
// static void  ptrGetActiveUniformsiv(GPGETACTIVEUNIFORMSIV fnptr, GLuint  program, GLsizei  uniformCount, const GLuint * uniformIndices, GLenum  pname, GLint * params) {(*fnptr)(program, uniformCount, uniformIndices, pname, params);}
// static void  ptrGetAttachedShaders(GPGETATTACHEDSHADERS fnptr, GLuint  program, GLsizei  maxCount, GLsizei * count, GLuint * shaders) {(*fnptr)(program, maxCount, count, shaders);}
// static GLint  ptrGetAttribLocation(GPGETATTRIBLOCATION fnptr, GLuint  program, const GLchar * name) {return (*fnptr)(program, name);}
// static void  ptrGetBooleani_v(GPGETBOOLEANI_V fnptr, GLenum  target, GLuint  index, GLboolean * data) {(*fnptr)(target, index, data);}
// static void  ptrGetBooleanv(GPGETBOOLEANV fnptr, GLenum  pname, GLboolean * data) {(*fnptr)(pname, data);}
// static void  ptrGetBufferParameteri64v(GPGETBUFFERPARAMETERI64V fnptr, GLenum  target, GLenum  pname, GLint64 * params) {(*fnptr)(target, pname, params);}
// static void  ptrGetBufferParameteriv(GPGETBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {(*fnptr)(target, pname, params);}
// static void  ptrGetBufferPointerv(GPGETBUFFERPOINTERV fnptr, GLenum  target, GLenum  pname, void ** params) {(*fnptr)(target, pname, params);}
// static void  ptrGetBufferSubData(GPGETBUFFERSUBDATA fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  size, void * data) {(*fnptr)(target, offset, size, data);}
// static void  ptrGetCompressedTexImage(GPGETCOMPRESSEDTEXIMAGE fnptr, GLenum  target, GLint  level, void * img) {(*fnptr)(target, level, img);}
// static void  ptrGetCompressedTextureImage(GPGETCOMPRESSEDTEXTUREIMAGE fnptr, GLuint  texture, GLint  level, GLsizei  bufSize, void * pixels) {(*fnptr)(texture, level, bufSize, pixels);}
// static void  ptrGetCompressedTextureSubImage(GPGETCOMPRESSEDTEXTURESUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLsizei  bufSize, void * pixels) {(*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, bufSize, pixels);}
// static GLuint  ptrGetDebugMessageLog(GPGETDEBUGMESSAGELOG fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);}
// static GLuint  ptrGetDebugMessageLogARB(GPGETDEBUGMESSAGELOGARB fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);}
// static GLuint  ptrGetDebugMessageLogKHR(GPGETDEBUGMESSAGELOGKHR fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);}
// static void  ptrGetDoublei_v(GPGETDOUBLEI_V fnptr, GLenum  target, GLuint  index, GLdouble * data) {(*fnptr)(target, index, data);}
// static void  ptrGetDoublev(GPGETDOUBLEV fnptr, GLenum  pname, GLdouble * data) {(*fnptr)(pname, data);}
// static GLenum  ptrGetError(GPGETERROR fnptr) {return (*fnptr)();}
// static void  ptrGetFloati_v(GPGETFLOATI_V fnptr, GLenum  target, GLuint  index, GLfloat * data) {(*fnptr)(target, index, data);}
// static void  ptrGetFloatv(GPGETFLOATV fnptr, GLenum  pname, GLfloat * data) {(*fnptr)(pname, data);}
// static GLint  ptrGetFragDataIndex(GPGETFRAGDATAINDEX fnptr, GLuint  program, const GLchar * name) {return (*fnptr)(program, name);}
// static GLint  ptrGetFragDataLocation(GPGETFRAGDATALOCATION fnptr, GLuint  program, const GLchar * name) {return (*fnptr)(program, name);}
// static void  ptrGetFramebufferAttachmentParameteriv(GPGETFRAMEBUFFERATTACHMENTPARAMETERIV fnptr, GLenum  target, GLenum  attachment, GLenum  pname, GLint * params) {(*fnptr)(target, attachment, pname, params);}
// static void  ptrGetFramebufferParameteriv(GPGETFRAMEBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {(*fnptr)(target, pname, params);}
// static GLenum  ptrGetGraphicsResetStatus(GPGETGRAPHICSRESETSTATUS fnptr) {return (*fnptr)();}
// static GLenum  ptrGetGraphicsResetStatusARB(GPGETGRAPHICSRESETSTATUSARB fnptr) {return (*fnptr)();}
// static GLenum  ptrGetGraphicsResetStatusKHR(GPGETGRAPHICSRESETSTATUSKHR fnptr) {return (*fnptr)();}
// static GLuint64  ptrGetImageHandleARB(GPGETIMAGEHANDLEARB fnptr, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  format) {return (*fnptr)(texture, level, layered, layer, format);}
// static void  ptrGetInteger64i_v(GPGETINTEGER64I_V fnptr, GLenum  target, GLuint  index, GLint64 * data) {(*fnptr)(target, index, data);}
// static void  ptrGetInteger64v(GPGETINTEGER64V fnptr, GLenum  pname, GLint64 * data) {(*fnptr)(pname, data);}
// static void  ptrGetIntegeri_v(GPGETINTEGERI_V fnptr, GLenum  target, GLuint  index, GLint * data) {(*fnptr)(target, index, data);}
// static void  ptrGetIntegerv(GPGETINTEGERV fnptr, GLenum  pname, GLint * data) {(*fnptr)(pname, data);}
// static void  ptrGetInternalformati64v(GPGETINTERNALFORMATI64V fnptr, GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint64 * params) {(*fnptr)(target, internalformat, pname, bufSize, params);}
// static void  ptrGetInternalformativ(GPGETINTERNALFORMATIV fnptr, GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint * params) {(*fnptr)(target, internalformat, pname, bufSize, params);}
// static void  ptrGetMultisamplefv(GPGETMULTISAMPLEFV fnptr, GLenum  pname, GLuint  index, GLfloat * val) {(*fnptr)(pname, index, val);}
// static void  ptrGetNamedBufferParameteri64v(GPGETNAMEDBUFFERPARAMETERI64V fnptr, GLuint  buffer, GLenum  pname, GLint64 * params) {(*fnptr)(buffer, pname, params);}
// static void  ptrGetNamedBufferParameteriv(GPGETNAMEDBUFFERPARAMETERIV fnptr, GLuint  buffer, GLenum  pname, GLint * params) {(*fnptr)(buffer, pname, params);}
// static void  ptrGetNamedBufferPointerv(GPGETNAMEDBUFFERPOINTERV fnptr, GLuint  buffer, GLenum  pname, void ** params) {(*fnptr)(buffer, pname, params);}
// static void  ptrGetNamedBufferSubData(GPGETNAMEDBUFFERSUBDATA fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, void * data) {(*fnptr)(buffer, offset, size, data);}
// static void  ptrGetNamedFramebufferAttachmentParameteriv(GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV fnptr, GLuint  framebuffer, GLenum  attachment, GLenum  pname, GLint * params) {(*fnptr)(framebuffer, attachment, pname, params);}
// static void  ptrGetNamedFramebufferParameteriv(GPGETNAMEDFRAMEBUFFERPARAMETERIV fnptr, GLuint  framebuffer, GLenum  pname, GLint * param) {(*fnptr)(framebuffer, pname, param);}
// static void  ptrGetNamedRenderbufferParameteriv(GPGETNAMEDRENDERBUFFERPARAMETERIV fnptr, GLuint  renderbuffer, GLenum  pname, GLint * params) {(*fnptr)(renderbuffer, pname, params);}
// static void  ptrGetNamedStringARB(GPGETNAMEDSTRINGARB fnptr, GLint  namelen, const GLchar * name, GLsizei  bufSize, GLint * stringlen, GLchar * string) {(*fnptr)(namelen, name, bufSize, stringlen, string);}
// static void  ptrGetNamedStringivARB(GPGETNAMEDSTRINGIVARB fnptr, GLint  namelen, const GLchar * name, GLenum  pname, GLint * params) {(*fnptr)(namelen, name, pname, params);}
// static void  ptrGetObjectLabel(GPGETOBJECTLABEL fnptr, GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label) {(*fnptr)(identifier, name, bufSize, length, label);}
// static void  ptrGetObjectLabelKHR(GPGETOBJECTLABELKHR fnptr, GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label) {(*fnptr)(identifier, name, bufSize, length, label);}
// static void  ptrGetObjectPtrLabel(GPGETOBJECTPTRLABEL fnptr, const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label) {(*fnptr)(ptr, bufSize, length, label);}
// static void  ptrGetObjectPtrLabelKHR(GPGETOBJECTPTRLABELKHR fnptr, const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label) {(*fnptr)(ptr, bufSize, length, label);}
// static void  ptrGetPointerv(GPGETPOINTERV fnptr, GLenum  pname, void ** params) {(*fnptr)(pname, params);}
// static void  ptrGetPointervKHR(GPGETPOINTERVKHR fnptr, GLenum  pname, void ** params) {(*fnptr)(pname, params);}
// static void  ptrGetProgramBinary(GPGETPROGRAMBINARY fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary) {(*fnptr)(program, bufSize, length, binaryFormat, binary);}
// static void  ptrGetProgramInfoLog(GPGETPROGRAMINFOLOG fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {(*fnptr)(program, bufSize, length, infoLog);}
// static void  ptrGetProgramInterfaceiv(GPGETPROGRAMINTERFACEIV fnptr, GLuint  program, GLenum  programInterface, GLenum  pname, GLint * params) {(*fnptr)(program, programInterface, pname, params);}
// static void  ptrGetProgramPipelineInfoLog(GPGETPROGRAMPIPELINEINFOLOG fnptr, GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {(*fnptr)(pipeline, bufSize, length, infoLog);}
// static void  ptrGetProgramPipelineiv(GPGETPROGRAMPIPELINEIV fnptr, GLuint  pipeline, GLenum  pname, GLint * params) {(*fnptr)(pipeline, pname, params);}
// static GLuint  ptrGetProgramResourceIndex(GPGETPROGRAMRESOURCEINDEX fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {return (*fnptr)(program, programInterface, name);}
// static GLint  ptrGetProgramResourceLocation(GPGETPROGRAMRESOURCELOCATION fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {return (*fnptr)(program, programInterface, name);}
// static GLint  ptrGetProgramResourceLocationIndex(GPGETPROGRAMRESOURCELOCATIONINDEX fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {return (*fnptr)(program, programInterface, name);}
// static void  ptrGetProgramResourceName(GPGETPROGRAMRESOURCENAME fnptr, GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  bufSize, GLsizei * length, GLchar * name) {(*fnptr)(program, programInterface, index, bufSize, length, name);}
// static void  ptrGetProgramResourceiv(GPGETPROGRAMRESOURCEIV fnptr, GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLint * params) {(*fnptr)(program, programInterface, index, propCount, props, bufSize, length, params);}
// static void  ptrGetProgramStageiv(GPGETPROGRAMSTAGEIV fnptr, GLuint  program, GLenum  shadertype, GLenum  pname, GLint * values) {(*fnptr)(program, shadertype, pname, values);}
// static void  ptrGetProgramiv(GPGETPROGRAMIV fnptr, GLuint  program, GLenum  pname, GLint * params) {(*fnptr)(program, pname, params);}
// static void  ptrGetQueryIndexediv(GPGETQUERYINDEXEDIV fnptr, GLenum  target, GLuint  index, GLenum  pname, GLint * params) {(*fnptr)(target, index, pname, params);}
// static void  ptrGetQueryObjecti64v(GPGETQUERYOBJECTI64V fnptr, GLuint  id, GLenum  pname, GLint64 * params) {(*fnptr)(id, pname, params);}
// static void  ptrGetQueryObjectiv(GPGETQUERYOBJECTIV fnptr, GLuint  id, GLenum  pname, GLint * params) {(*fnptr)(id, pname, params);}
// static void  ptrGetQueryObjectui64v(GPGETQUERYOBJECTUI64V fnptr, GLuint  id, GLenum  pname, GLuint64 * params) {(*fnptr)(id, pname, params);}
// static void  ptrGetQueryObjectuiv(GPGETQUERYOBJECTUIV fnptr, GLuint  id, GLenum  pname, GLuint * params) {(*fnptr)(id, pname, params);}
// static void  ptrGetQueryiv(GPGETQUERYIV fnptr, GLenum  target, GLenum  pname, GLint * params) {(*fnptr)(target, pname, params);}
// static void  ptrGetRenderbufferParameteriv(GPGETRENDERBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {(*fnptr)(target, pname, params);}
// static void  ptrGetSamplerParameterIiv(GPGETSAMPLERPARAMETERIIV fnptr, GLuint  sampler, GLenum  pname, GLint * params) {(*fnptr)(sampler, pname, params);}
// static void  ptrGetSamplerParameterIuiv(GPGETSAMPLERPARAMETERIUIV fnptr, GLuint  sampler, GLenum  pname, GLuint * params) {(*fnptr)(sampler, pname, params);}
// static void  ptrGetSamplerParameterfv(GPGETSAMPLERPARAMETERFV fnptr, GLuint  sampler, GLenum  pname, GLfloat * params) {(*fnptr)(sampler, pname, params);}
// static void  ptrGetSamplerParameteriv(GPGETSAMPLERPARAMETERIV fnptr, GLuint  sampler, GLenum  pname, GLint * params) {(*fnptr)(sampler, pname, params);}
// static void  ptrGetShaderInfoLog(GPGETSHADERINFOLOG fnptr, GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {(*fnptr)(shader, bufSize, length, infoLog);}
// static void  ptrGetShaderPrecisionFormat(GPGETSHADERPRECISIONFORMAT fnptr, GLenum  shadertype, GLenum  precisiontype, GLint * range, GLint * precision) {(*fnptr)(shadertype, precisiontype, range, precision);}
// static void  ptrGetShaderSource(GPGETSHADERSOURCE fnptr, GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * source) {(*fnptr)(shader, bufSize, length, source);}
// static void  ptrGetShaderiv(GPGETSHADERIV fnptr, GLuint  shader, GLenum  pname, GLint * params) {(*fnptr)(shader, pname, params);}
// static const GLubyte * ptrGetString(GPGETSTRING fnptr, GLenum  name) {return (*fnptr)(name);}
// static const GLubyte * ptrGetStringi(GPGETSTRINGI fnptr, GLenum  name, GLuint  index) {return (*fnptr)(name, index);}
// static GLuint  ptrGetSubroutineIndex(GPGETSUBROUTINEINDEX fnptr, GLuint  program, GLenum  shadertype, const GLchar * name) {return (*fnptr)(program, shadertype, name);}
// static GLint  ptrGetSubroutineUniformLocation(GPGETSUBROUTINEUNIFORMLOCATION fnptr, GLuint  program, GLenum  shadertype, const GLchar * name) {return (*fnptr)(program, shadertype, name);}
// static void  ptrGetSynciv(GPGETSYNCIV fnptr, GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values) {(*fnptr)(sync, pname, bufSize, length, values);}
// static void  ptrGetTexImage(GPGETTEXIMAGE fnptr, GLenum  target, GLint  level, GLenum  format, GLenum  type, void * pixels) {(*fnptr)(target, level, format, type, pixels);}
// static void  ptrGetTexLevelParameterfv(GPGETTEXLEVELPARAMETERFV fnptr, GLenum  target, GLint  level, GLenum  pname, GLfloat * params) {(*fnptr)(target, level, pname, params);}
// static void  ptrGetTexLevelParameteriv(GPGETTEXLEVELPARAMETERIV fnptr, GLenum  target, GLint  level, GLenum  pname, GLint * params) {(*fnptr)(target, level, pname, params);}
// static void  ptrGetTexParameterIiv(GPGETTEXPARAMETERIIV fnptr, GLenum  target, GLenum  pname, GLint * params) {(*fnptr)(target, pname, params);}
// static void  ptrGetTexParameterIuiv(GPGETTEXPARAMETERIUIV fnptr, GLenum  target, GLenum  pname, GLuint * params) {(*fnptr)(target, pname, params);}
// static void  ptrGetTexParameterfv(GPGETTEXPARAMETERFV fnptr, GLenum  target, GLenum  pname, GLfloat * params) {(*fnptr)(target, pname, params);}
// static void  ptrGetTexParameteriv(GPGETTEXPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {(*fnptr)(target, pname, params);}
// static GLuint64  ptrGetTextureHandleARB(GPGETTEXTUREHANDLEARB fnptr, GLuint  texture) {return (*fnptr)(texture);}
// static void  ptrGetTextureImage(GPGETTEXTUREIMAGE fnptr, GLuint  texture, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels) {(*fnptr)(texture, level, format, type, bufSize, pixels);}
// static void  ptrGetTextureLevelParameterfv(GPGETTEXTURELEVELPARAMETERFV fnptr, GLuint  texture, GLint  level, GLenum  pname, GLfloat * params) {(*fnptr)(texture, level, pname, params);}
// static void  ptrGetTextureLevelParameteriv(GPGETTEXTURELEVELPARAMETERIV fnptr, GLuint  texture, GLint  level, GLenum  pname, GLint * params) {(*fnptr)(texture, level, pname, params);}
// static void  ptrGetTextureParameterIiv(GPGETTEXTUREPARAMETERIIV fnptr, GLuint  texture, GLenum  pname, GLint * params) {(*fnptr)(texture, pname, params);}
// static void  ptrGetTextureParameterIuiv(GPGETTEXTUREPARAMETERIUIV fnptr, GLuint  texture, GLenum  pname, GLuint * params) {(*fnptr)(texture, pname, params);}
// static void  ptrGetTextureParameterfv(GPGETTEXTUREPARAMETERFV fnptr, GLuint  texture, GLenum  pname, GLfloat * params) {(*fnptr)(texture, pname, params);}
// static void  ptrGetTextureParameteriv(GPGETTEXTUREPARAMETERIV fnptr, GLuint  texture, GLenum  pname, GLint * params) {(*fnptr)(texture, pname, params);}
// static GLuint64  ptrGetTextureSamplerHandleARB(GPGETTEXTURESAMPLERHANDLEARB fnptr, GLuint  texture, GLuint  sampler) {return (*fnptr)(texture, sampler);}
// static void  ptrGetTextureSubImage(GPGETTEXTURESUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels) {(*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, bufSize, pixels);}
// static void  ptrGetTransformFeedbackVarying(GPGETTRANSFORMFEEDBACKVARYING fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLsizei * size, GLenum * type, GLchar * name) {(*fnptr)(program, index, bufSize, length, size, type, name);}
// static void  ptrGetTransformFeedbacki64_v(GPGETTRANSFORMFEEDBACKI64_V fnptr, GLuint  xfb, GLenum  pname, GLuint  index, GLint64 * param) {(*fnptr)(xfb, pname, index, param);}
// static void  ptrGetTransformFeedbacki_v(GPGETTRANSFORMFEEDBACKI_V fnptr, GLuint  xfb, GLenum  pname, GLuint  index, GLint * param) {(*fnptr)(xfb, pname, index, param);}
// static void  ptrGetTransformFeedbackiv(GPGETTRANSFORMFEEDBACKIV fnptr, GLuint  xfb, GLenum  pname, GLint * param) {(*fnptr)(xfb, pname, param);}
// static GLuint  ptrGetUniformBlockIndex(GPGETUNIFORMBLOCKINDEX fnptr, GLuint  program, const GLchar * uniformBlockName) {return (*fnptr)(program, uniformBlockName);}
// static void  ptrGetUniformIndices(GPGETUNIFORMINDICES fnptr, GLuint  program, GLsizei  uniformCount, const GLchar *const* uniformNames, GLuint * uniformIndices) {(*fnptr)(program, uniformCount, uniformNames, uniformIndices);}
// static GLint  ptrGetUniformLocation(GPGETUNIFORMLOCATION fnptr, GLuint  program, const GLchar * name) {return (*fnptr)(program, name);}
// static void  ptrGetUniformSubroutineuiv(GPGETUNIFORMSUBROUTINEUIV fnptr, GLenum  shadertype, GLint  location, GLuint * params) {(*fnptr)(shadertype, location, params);}
// static void  ptrGetUniformdv(GPGETUNIFORMDV fnptr, GLuint  program, GLint  location, GLdouble * params) {(*fnptr)(program, location, params);}
// static void  ptrGetUniformfv(GPGETUNIFORMFV fnptr, GLuint  program, GLint  location, GLfloat * params) {(*fnptr)(program, location, params);}
// static void  ptrGetUniformiv(GPGETUNIFORMIV fnptr, GLuint  program, GLint  location, GLint * params) {(*fnptr)(program, location, params);}
// static void  ptrGetUniformuiv(GPGETUNIFORMUIV fnptr, GLuint  program, GLint  location, GLuint * params) {(*fnptr)(program, location, params);}
// static void  ptrGetVertexArrayIndexed64iv(GPGETVERTEXARRAYINDEXED64IV fnptr, GLuint  vaobj, GLuint  index, GLenum  pname, GLint64 * param) {(*fnptr)(vaobj, index, pname, param);}
// static void  ptrGetVertexArrayIndexediv(GPGETVERTEXARRAYINDEXEDIV fnptr, GLuint  vaobj, GLuint  index, GLenum  pname, GLint * param) {(*fnptr)(vaobj, index, pname, param);}
// static void  ptrGetVertexArrayiv(GPGETVERTEXARRAYIV fnptr, GLuint  vaobj, GLenum  pname, GLint * param) {(*fnptr)(vaobj, pname, param);}
// static void  ptrGetVertexAttribIiv(GPGETVERTEXATTRIBIIV fnptr, GLuint  index, GLenum  pname, GLint * params) {(*fnptr)(index, pname, params);}
// static void  ptrGetVertexAttribIuiv(GPGETVERTEXATTRIBIUIV fnptr, GLuint  index, GLenum  pname, GLuint * params) {(*fnptr)(index, pname, params);}
// static void  ptrGetVertexAttribLdv(GPGETVERTEXATTRIBLDV fnptr, GLuint  index, GLenum  pname, GLdouble * params) {(*fnptr)(index, pname, params);}
// static void  ptrGetVertexAttribLui64vARB(GPGETVERTEXATTRIBLUI64VARB fnptr, GLuint  index, GLenum  pname, GLuint64EXT * params) {(*fnptr)(index, pname, params);}
// static void  ptrGetVertexAttribPointerv(GPGETVERTEXATTRIBPOINTERV fnptr, GLuint  index, GLenum  pname, void ** pointer) {(*fnptr)(index, pname, pointer);}
// static void  ptrGetVertexAttribdv(GPGETVERTEXATTRIBDV fnptr, GLuint  index, GLenum  pname, GLdouble * params) {(*fnptr)(index, pname, params);}
// static void  ptrGetVertexAttribfv(GPGETVERTEXATTRIBFV fnptr, GLuint  index, GLenum  pname, GLfloat * params) {(*fnptr)(index, pname, params);}
// static void  ptrGetVertexAttribiv(GPGETVERTEXATTRIBIV fnptr, GLuint  index, GLenum  pname, GLint * params) {(*fnptr)(index, pname, params);}
// static void  ptrGetnCompressedTexImage(GPGETNCOMPRESSEDTEXIMAGE fnptr, GLenum  target, GLint  lod, GLsizei  bufSize, void * pixels) {(*fnptr)(target, lod, bufSize, pixels);}
// static void  ptrGetnCompressedTexImageARB(GPGETNCOMPRESSEDTEXIMAGEARB fnptr, GLenum  target, GLint  lod, GLsizei  bufSize, void * img) {(*fnptr)(target, lod, bufSize, img);}
// static void  ptrGetnTexImage(GPGETNTEXIMAGE fnptr, GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels) {(*fnptr)(target, level, format, type, bufSize, pixels);}
// static void  ptrGetnTexImageARB(GPGETNTEXIMAGEARB fnptr, GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * img) {(*fnptr)(target, level, format, type, bufSize, img);}
// static void  ptrGetnUniformdv(GPGETNUNIFORMDV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformdvARB(GPGETNUNIFORMDVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformfv(GPGETNUNIFORMFV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformfvARB(GPGETNUNIFORMFVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformfvKHR(GPGETNUNIFORMFVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformiv(GPGETNUNIFORMIV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformivARB(GPGETNUNIFORMIVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformivKHR(GPGETNUNIFORMIVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformuiv(GPGETNUNIFORMUIV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformuivARB(GPGETNUNIFORMUIVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrGetnUniformuivKHR(GPGETNUNIFORMUIVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {(*fnptr)(program, location, bufSize, params);}
// static void  ptrHint(GPHINT fnptr, GLenum  target, GLenum  mode) {(*fnptr)(target, mode);}
// static void  ptrInvalidateBufferData(GPINVALIDATEBUFFERDATA fnptr, GLuint  buffer) {(*fnptr)(buffer);}
// static void  ptrInvalidateBufferSubData(GPINVALIDATEBUFFERSUBDATA fnptr, GLuint  buffer, GLintptr  offset, GLsizeiptr  length) {(*fnptr)(buffer, offset, length);}
// static void  ptrInvalidateFramebuffer(GPINVALIDATEFRAMEBUFFER fnptr, GLenum  target, GLsizei  numAttachments, const GLenum * attachments) {(*fnptr)(target, numAttachments, attachments);}
// static void  ptrInvalidateNamedFramebufferData(GPINVALIDATENAMEDFRAMEBUFFERDATA fnptr, GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments) {(*fnptr)(framebuffer, numAttachments, attachments);}
// static void  ptrInvalidateNamedFramebufferSubData(GPINVALIDATENAMEDFRAMEBUFFERSUBDATA fnptr, GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {(*fnptr)(framebuffer, numAttachments, attachments, x, y, width, height);}
// static void  ptrInvalidateSubFramebuffer(GPINVALIDATESUBFRAMEBUFFER fnptr, GLenum  target, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {(*fnptr)(target, numAttachments, attachments, x, y, width, height);}
// static void  ptrInvalidateTexImage(GPINVALIDATETEXIMAGE fnptr, GLuint  texture, GLint  level) {(*fnptr)(texture, level);}
// static void  ptrInvalidateTexSubImage(GPINVALIDATETEXSUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth) {(*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth);}
// static GLboolean  ptrIsBuffer(GPISBUFFER fnptr, GLuint  buffer) {return (*fnptr)(buffer);}
// static GLboolean  ptrIsEnabled(GPISENABLED fnptr, GLenum  cap) {return (*fnptr)(cap);}
// static GLboolean  ptrIsEnabledi(GPISENABLEDI fnptr, GLenum  target, GLuint  index) {return (*fnptr)(target, index);}
// static GLboolean  ptrIsFramebuffer(GPISFRAMEBUFFER fnptr, GLuint  framebuffer) {return (*fnptr)(framebuffer);}
// static GLboolean  ptrIsImageHandleResidentARB(GPISIMAGEHANDLERESIDENTARB fnptr, GLuint64  handle) {return (*fnptr)(handle);}
// static GLboolean  ptrIsNamedStringARB(GPISNAMEDSTRINGARB fnptr, GLint  namelen, const GLchar * name) {return (*fnptr)(namelen, name);}
// static GLboolean  ptrIsProgram(GPISPROGRAM fnptr, GLuint  program) {return (*fnptr)(program);}
// static GLboolean  ptrIsProgramPipeline(GPISPROGRAMPIPELINE fnptr, GLuint  pipeline) {return (*fnptr)(pipeline);}
// static GLboolean  ptrIsQuery(GPISQUERY fnptr, GLuint  id) {return (*fnptr)(id);}
// static GLboolean  ptrIsRenderbuffer(GPISRENDERBUFFER fnptr, GLuint  renderbuffer) {return (*fnptr)(renderbuffer);}
// static GLboolean  ptrIsSampler(GPISSAMPLER fnptr, GLuint  sampler) {return (*fnptr)(sampler);}
// static GLboolean  ptrIsShader(GPISSHADER fnptr, GLuint  shader) {return (*fnptr)(shader);}
// static GLboolean  ptrIsSync(GPISSYNC fnptr, GLsync  sync) {return (*fnptr)(sync);}
// static GLboolean  ptrIsTexture(GPISTEXTURE fnptr, GLuint  texture) {return (*fnptr)(texture);}
// static GLboolean  ptrIsTextureHandleResidentARB(GPISTEXTUREHANDLERESIDENTARB fnptr, GLuint64  handle) {return (*fnptr)(handle);}
// static GLboolean  ptrIsTransformFeedback(GPISTRANSFORMFEEDBACK fnptr, GLuint  id) {return (*fnptr)(id);}
// static GLboolean  ptrIsVertexArray(GPISVERTEXARRAY fnptr, GLuint  array) {return (*fnptr)(array);}
// static void  ptrLineWidth(GPLINEWIDTH fnptr, GLfloat  width) {(*fnptr)(width);}
// static void  ptrLinkProgram(GPLINKPROGRAM fnptr, GLuint  program) {(*fnptr)(program);}
// static void  ptrLogicOp(GPLOGICOP fnptr, GLenum  opcode) {(*fnptr)(opcode);}
// static void  ptrMakeImageHandleNonResidentARB(GPMAKEIMAGEHANDLENONRESIDENTARB fnptr, GLuint64  handle) {(*fnptr)(handle);}
// static void  ptrMakeImageHandleResidentARB(GPMAKEIMAGEHANDLERESIDENTARB fnptr, GLuint64  handle, GLenum  access) {(*fnptr)(handle, access);}
// static void  ptrMakeTextureHandleNonResidentARB(GPMAKETEXTUREHANDLENONRESIDENTARB fnptr, GLuint64  handle) {(*fnptr)(handle);}
// static void  ptrMakeTextureHandleResidentARB(GPMAKETEXTUREHANDLERESIDENTARB fnptr, GLuint64  handle) {(*fnptr)(handle);}
// static void * ptrMapBuffer(GPMAPBUFFER fnptr, GLenum  target, GLenum  access) {return (*fnptr)(target, access);}
// static void * ptrMapBufferRange(GPMAPBUFFERRANGE fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access) {return (*fnptr)(target, offset, length, access);}
// static void * ptrMapNamedBuffer(GPMAPNAMEDBUFFER fnptr, GLuint  buffer, GLenum  access) {return (*fnptr)(buffer, access);}
// static void * ptrMapNamedBufferRange(GPMAPNAMEDBUFFERRANGE fnptr, GLuint  buffer, GLintptr  offset, GLsizei  length, GLbitfield  access) {return (*fnptr)(buffer, offset, length, access);}
// static void  ptrMemoryBarrier(GPMEMORYBARRIER fnptr, GLbitfield  barriers) {(*fnptr)(barriers);}
// static void  ptrMemoryBarrierByRegion(GPMEMORYBARRIERBYREGION fnptr, GLbitfield  barriers) {(*fnptr)(barriers);}
// static void  ptrMinSampleShading(GPMINSAMPLESHADING fnptr, GLfloat  value) {(*fnptr)(value);}
// static void  ptrMinSampleShadingARB(GPMINSAMPLESHADINGARB fnptr, GLfloat  value) {(*fnptr)(value);}
// static void  ptrMultiDrawArrays(GPMULTIDRAWARRAYS fnptr, GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  drawcount) {(*fnptr)(mode, first, count, drawcount);}
// static void  ptrMultiDrawArraysIndirect(GPMULTIDRAWARRAYSINDIRECT fnptr, GLenum  mode, const void * indirect, GLsizei  drawcount, GLsizei  stride) {(*fnptr)(mode, indirect, drawcount, stride);}
// static void  ptrMultiDrawArraysIndirectCountARB(GPMULTIDRAWARRAYSINDIRECTCOUNTARB fnptr, GLenum  mode, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride) {(*fnptr)(mode, indirect, drawcount, maxdrawcount, stride);}
// static void  ptrMultiDrawElements(GPMULTIDRAWELEMENTS fnptr, GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount) {(*fnptr)(mode, count, type, indices, drawcount);}
// static void  ptrMultiDrawElementsBaseVertex(GPMULTIDRAWELEMENTSBASEVERTEX fnptr, GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount, const GLint * basevertex) {(*fnptr)(mode, count, type, indices, drawcount, basevertex);}
// static void  ptrMultiDrawElementsIndirect(GPMULTIDRAWELEMENTSINDIRECT fnptr, GLenum  mode, GLenum  type, const void * indirect, GLsizei  drawcount, GLsizei  stride) {(*fnptr)(mode, type, indirect, drawcount, stride);}
// static void  ptrMultiDrawElementsIndirectCountARB(GPMULTIDRAWELEMENTSINDIRECTCOUNTARB fnptr, GLenum  mode, GLenum  type, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride) {(*fnptr)(mode, type, indirect, drawcount, maxdrawcount, stride);}
// static void  ptrNamedBufferData(GPNAMEDBUFFERDATA fnptr, GLuint  buffer, GLsizei  size, const void * data, GLenum  usage) {(*fnptr)(buffer, size, data, usage);}
// static void  ptrNamedBufferPageCommitmentARB(GPNAMEDBUFFERPAGECOMMITMENTARB fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit) {(*fnptr)(buffer, offset, size, commit);}
// static void  ptrNamedBufferPageCommitmentEXT(GPNAMEDBUFFERPAGECOMMITMENTEXT fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit) {(*fnptr)(buffer, offset, size, commit);}
// static void  ptrNamedBufferStorage(GPNAMEDBUFFERSTORAGE fnptr, GLuint  buffer, GLsizei  size, const void * data, GLbitfield  flags) {(*fnptr)(buffer, size, data, flags);}
// static void  ptrNamedBufferSubData(GPNAMEDBUFFERSUBDATA fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, const void * data) {(*fnptr)(buffer, offset, size, data);}
// static void  ptrNamedFramebufferDrawBuffer(GPNAMEDFRAMEBUFFERDRAWBUFFER fnptr, GLuint  framebuffer, GLenum  buf) {(*fnptr)(framebuffer, buf);}
// static void  ptrNamedFramebufferDrawBuffers(GPNAMEDFRAMEBUFFERDRAWBUFFERS fnptr, GLuint  framebuffer, GLsizei  n, const GLenum * bufs) {(*fnptr)(framebuffer, n, bufs);}
// static void  ptrNamedFramebufferParameteri(GPNAMEDFRAMEBUFFERPARAMETERI fnptr, GLuint  framebuffer, GLenum  pname, GLint  param) {(*fnptr)(framebuffer, pname, param);}
// static void  ptrNamedFramebufferReadBuffer(GPNAMEDFRAMEBUFFERREADBUFFER fnptr, GLuint  framebuffer, GLenum  src) {(*fnptr)(framebuffer, src);}
// static void  ptrNamedFramebufferRenderbuffer(GPNAMEDFRAMEBUFFERRENDERBUFFER fnptr, GLuint  framebuffer, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer) {(*fnptr)(framebuffer, attachment, renderbuffertarget, renderbuffer);}
// static void  ptrNamedFramebufferTexture(GPNAMEDFRAMEBUFFERTEXTURE fnptr, GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level) {(*fnptr)(framebuffer, attachment, texture, level);}
// static void  ptrNamedFramebufferTextureLayer(GPNAMEDFRAMEBUFFERTEXTURELAYER fnptr, GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer) {(*fnptr)(framebuffer, attachment, texture, level, layer);}
// static void  ptrNamedRenderbufferStorage(GPNAMEDRENDERBUFFERSTORAGE fnptr, GLuint  renderbuffer, GLenum  internalformat, GLsizei  width, GLsizei  height) {(*fnptr)(renderbuffer, internalformat, width, height);}
// static void  ptrNamedRenderbufferStorageMultisample(GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE fnptr, GLuint  renderbuffer, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {(*fnptr)(renderbuffer, samples, internalformat, width, height);}
// static void  ptrNamedStringARB(GPNAMEDSTRINGARB fnptr, GLenum  type, GLint  namelen, const GLchar * name, GLint  stringlen, const GLchar * string) {(*fnptr)(type, namelen, name, stringlen, string);}
// static void  ptrObjectLabel(GPOBJECTLABEL fnptr, GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label) {(*fnptr)(identifier, name, length, label);}
// static void  ptrObjectLabelKHR(GPOBJECTLABELKHR fnptr, GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label) {(*fnptr)(identifier, name, length, label);}
// static void  ptrObjectPtrLabel(GPOBJECTPTRLABEL fnptr, const void * ptr, GLsizei  length, const GLchar * label) {(*fnptr)(ptr, length, label);}
// static void  ptrObjectPtrLabelKHR(GPOBJECTPTRLABELKHR fnptr, const void * ptr, GLsizei  length, const GLchar * label) {(*fnptr)(ptr, length, label);}
// static void  ptrPatchParameterfv(GPPATCHPARAMETERFV fnptr, GLenum  pname, const GLfloat * values) {(*fnptr)(pname, values);}
// static void  ptrPatchParameteri(GPPATCHPARAMETERI fnptr, GLenum  pname, GLint  value) {(*fnptr)(pname, value);}
// static void  ptrPauseTransformFeedback(GPPAUSETRANSFORMFEEDBACK fnptr) {(*fnptr)();}
// static void  ptrPixelStoref(GPPIXELSTOREF fnptr, GLenum  pname, GLfloat  param) {(*fnptr)(pname, param);}
// static void  ptrPixelStorei(GPPIXELSTOREI fnptr, GLenum  pname, GLint  param) {(*fnptr)(pname, param);}
// static void  ptrPointParameterf(GPPOINTPARAMETERF fnptr, GLenum  pname, GLfloat  param) {(*fnptr)(pname, param);}
// static void  ptrPointParameterfv(GPPOINTPARAMETERFV fnptr, GLenum  pname, const GLfloat * params) {(*fnptr)(pname, params);}
// static void  ptrPointParameteri(GPPOINTPARAMETERI fnptr, GLenum  pname, GLint  param) {(*fnptr)(pname, param);}
// static void  ptrPointParameteriv(GPPOINTPARAMETERIV fnptr, GLenum  pname, const GLint * params) {(*fnptr)(pname, params);}
// static void  ptrPointSize(GPPOINTSIZE fnptr, GLfloat  size) {(*fnptr)(size);}
// static void  ptrPolygonMode(GPPOLYGONMODE fnptr, GLenum  face, GLenum  mode) {(*fnptr)(face, mode);}
// static void  ptrPolygonOffset(GPPOLYGONOFFSET fnptr, GLfloat  factor, GLfloat  units) {(*fnptr)(factor, units);}
// static void  ptrPopDebugGroup(GPPOPDEBUGGROUP fnptr) {(*fnptr)();}
// static void  ptrPopDebugGroupKHR(GPPOPDEBUGGROUPKHR fnptr) {(*fnptr)();}
// static void  ptrPrimitiveRestartIndex(GPPRIMITIVERESTARTINDEX fnptr, GLuint  index) {(*fnptr)(index);}
// static void  ptrProgramBinary(GPPROGRAMBINARY fnptr, GLuint  program, GLenum  binaryFormat, const void * binary, GLsizei  length) {(*fnptr)(program, binaryFormat, binary, length);}
// static void  ptrProgramParameteri(GPPROGRAMPARAMETERI fnptr, GLuint  program, GLenum  pname, GLint  value) {(*fnptr)(program, pname, value);}
// static void  ptrProgramUniform1d(GPPROGRAMUNIFORM1D fnptr, GLuint  program, GLint  location, GLdouble  v0) {(*fnptr)(program, location, v0);}
// static void  ptrProgramUniform1dv(GPPROGRAMUNIFORM1DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform1f(GPPROGRAMUNIFORM1F fnptr, GLuint  program, GLint  location, GLfloat  v0) {(*fnptr)(program, location, v0);}
// static void  ptrProgramUniform1fv(GPPROGRAMUNIFORM1FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform1i(GPPROGRAMUNIFORM1I fnptr, GLuint  program, GLint  location, GLint  v0) {(*fnptr)(program, location, v0);}
// static void  ptrProgramUniform1iv(GPPROGRAMUNIFORM1IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform1ui(GPPROGRAMUNIFORM1UI fnptr, GLuint  program, GLint  location, GLuint  v0) {(*fnptr)(program, location, v0);}
// static void  ptrProgramUniform1uiv(GPPROGRAMUNIFORM1UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform2d(GPPROGRAMUNIFORM2D fnptr, GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1) {(*fnptr)(program, location, v0, v1);}
// static void  ptrProgramUniform2dv(GPPROGRAMUNIFORM2DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform2f(GPPROGRAMUNIFORM2F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1) {(*fnptr)(program, location, v0, v1);}
// static void  ptrProgramUniform2fv(GPPROGRAMUNIFORM2FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform2i(GPPROGRAMUNIFORM2I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1) {(*fnptr)(program, location, v0, v1);}
// static void  ptrProgramUniform2iv(GPPROGRAMUNIFORM2IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform2ui(GPPROGRAMUNIFORM2UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1) {(*fnptr)(program, location, v0, v1);}
// static void  ptrProgramUniform2uiv(GPPROGRAMUNIFORM2UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform3d(GPPROGRAMUNIFORM3D fnptr, GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2) {(*fnptr)(program, location, v0, v1, v2);}
// static void  ptrProgramUniform3dv(GPPROGRAMUNIFORM3DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform3f(GPPROGRAMUNIFORM3F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {(*fnptr)(program, location, v0, v1, v2);}
// static void  ptrProgramUniform3fv(GPPROGRAMUNIFORM3FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform3i(GPPROGRAMUNIFORM3I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2) {(*fnptr)(program, location, v0, v1, v2);}
// static void  ptrProgramUniform3iv(GPPROGRAMUNIFORM3IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform3ui(GPPROGRAMUNIFORM3UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {(*fnptr)(program, location, v0, v1, v2);}
// static void  ptrProgramUniform3uiv(GPPROGRAMUNIFORM3UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform4d(GPPROGRAMUNIFORM4D fnptr, GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2, GLdouble  v3) {(*fnptr)(program, location, v0, v1, v2, v3);}
// static void  ptrProgramUniform4dv(GPPROGRAMUNIFORM4DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform4f(GPPROGRAMUNIFORM4F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {(*fnptr)(program, location, v0, v1, v2, v3);}
// static void  ptrProgramUniform4fv(GPPROGRAMUNIFORM4FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform4i(GPPROGRAMUNIFORM4I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {(*fnptr)(program, location, v0, v1, v2, v3);}
// static void  ptrProgramUniform4iv(GPPROGRAMUNIFORM4IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniform4ui(GPPROGRAMUNIFORM4UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {(*fnptr)(program, location, v0, v1, v2, v3);}
// static void  ptrProgramUniform4uiv(GPPROGRAMUNIFORM4UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {(*fnptr)(program, location, count, value);}
// static void  ptrProgramUniformHandleui64ARB(GPPROGRAMUNIFORMHANDLEUI64ARB fnptr, GLuint  program, GLint  location, GLuint64  value) {(*fnptr)(program, location, value);}
// static void  ptrProgramUniformHandleui64vARB(GPPROGRAMUNIFORMHANDLEUI64VARB fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values) {(*fnptr)(program, location, count, values);}
// static void  ptrProgramUniformMatrix2dv(GPPROGRAMUNIFORMMATRIX2DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix2fv(GPPROGRAMUNIFORMMATRIX2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix2x3dv(GPPROGRAMUNIFORMMATRIX2X3DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix2x3fv(GPPROGRAMUNIFORMMATRIX2X3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix2x4dv(GPPROGRAMUNIFORMMATRIX2X4DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix2x4fv(GPPROGRAMUNIFORMMATRIX2X4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix3dv(GPPROGRAMUNIFORMMATRIX3DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix3fv(GPPROGRAMUNIFORMMATRIX3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix3x2dv(GPPROGRAMUNIFORMMATRIX3X2DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix3x2fv(GPPROGRAMUNIFORMMATRIX3X2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix3x4dv(GPPROGRAMUNIFORMMATRIX3X4DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix3x4fv(GPPROGRAMUNIFORMMATRIX3X4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix4dv(GPPROGRAMUNIFORMMATRIX4DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix4fv(GPPROGRAMUNIFORMMATRIX4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix4x2dv(GPPROGRAMUNIFORMMATRIX4X2DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix4x2fv(GPPROGRAMUNIFORMMATRIX4X2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix4x3dv(GPPROGRAMUNIFORMMATRIX4X3DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProgramUniformMatrix4x3fv(GPPROGRAMUNIFORMMATRIX4X3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(program, location, count, transpose, value);}
// static void  ptrProvokingVertex(GPPROVOKINGVERTEX fnptr, GLenum  mode) {(*fnptr)(mode);}
// static void  ptrPushDebugGroup(GPPUSHDEBUGGROUP fnptr, GLenum  source, GLuint  id, GLsizei  length, const GLchar * message) {(*fnptr)(source, id, length, message);}
// static void  ptrPushDebugGroupKHR(GPPUSHDEBUGGROUPKHR fnptr, GLenum  source, GLuint  id, GLsizei  length, const GLchar * message) {(*fnptr)(source, id, length, message);}
// static void  ptrQueryCounter(GPQUERYCOUNTER fnptr, GLuint  id, GLenum  target) {(*fnptr)(id, target);}
// static void  ptrReadBuffer(GPREADBUFFER fnptr, GLenum  src) {(*fnptr)(src);}
// static void  ptrReadPixels(GPREADPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels) {(*fnptr)(x, y, width, height, format, type, pixels);}
// static void  ptrReadnPixels(GPREADNPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {(*fnptr)(x, y, width, height, format, type, bufSize, data);}
// static void  ptrReadnPixelsARB(GPREADNPIXELSARB fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {(*fnptr)(x, y, width, height, format, type, bufSize, data);}
// static void  ptrReadnPixelsKHR(GPREADNPIXELSKHR fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {(*fnptr)(x, y, width, height, format, type, bufSize, data);}
// static void  ptrReleaseShaderCompiler(GPRELEASESHADERCOMPILER fnptr) {(*fnptr)();}
// static void  ptrRenderbufferStorage(GPRENDERBUFFERSTORAGE fnptr, GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height) {(*fnptr)(target, internalformat, width, height);}
// static void  ptrRenderbufferStorageMultisample(GPRENDERBUFFERSTORAGEMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {(*fnptr)(target, samples, internalformat, width, height);}
// static void  ptrResumeTransformFeedback(GPRESUMETRANSFORMFEEDBACK fnptr) {(*fnptr)();}
// static void  ptrSampleCoverage(GPSAMPLECOVERAGE fnptr, GLfloat  value, GLboolean  invert) {(*fnptr)(value, invert);}
// static void  ptrSampleMaski(GPSAMPLEMASKI fnptr, GLuint  maskNumber, GLbitfield  mask) {(*fnptr)(maskNumber, mask);}
// static void  ptrSamplerParameterIiv(GPSAMPLERPARAMETERIIV fnptr, GLuint  sampler, GLenum  pname, const GLint * param) {(*fnptr)(sampler, pname, param);}
// static void  ptrSamplerParameterIuiv(GPSAMPLERPARAMETERIUIV fnptr, GLuint  sampler, GLenum  pname, const GLuint * param) {(*fnptr)(sampler, pname, param);}
// static void  ptrSamplerParameterf(GPSAMPLERPARAMETERF fnptr, GLuint  sampler, GLenum  pname, GLfloat  param) {(*fnptr)(sampler, pname, param);}
// static void  ptrSamplerParameterfv(GPSAMPLERPARAMETERFV fnptr, GLuint  sampler, GLenum  pname, const GLfloat * param) {(*fnptr)(sampler, pname, param);}
// static void  ptrSamplerParameteri(GPSAMPLERPARAMETERI fnptr, GLuint  sampler, GLenum  pname, GLint  param) {(*fnptr)(sampler, pname, param);}
// static void  ptrSamplerParameteriv(GPSAMPLERPARAMETERIV fnptr, GLuint  sampler, GLenum  pname, const GLint * param) {(*fnptr)(sampler, pname, param);}
// static void  ptrScissor(GPSCISSOR fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {(*fnptr)(x, y, width, height);}
// static void  ptrScissorArrayv(GPSCISSORARRAYV fnptr, GLuint  first, GLsizei  count, const GLint * v) {(*fnptr)(first, count, v);}
// static void  ptrScissorIndexed(GPSCISSORINDEXED fnptr, GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height) {(*fnptr)(index, left, bottom, width, height);}
// static void  ptrScissorIndexedv(GPSCISSORINDEXEDV fnptr, GLuint  index, const GLint * v) {(*fnptr)(index, v);}
// static void  ptrShaderBinary(GPSHADERBINARY fnptr, GLsizei  count, const GLuint * shaders, GLenum  binaryformat, const void * binary, GLsizei  length) {(*fnptr)(count, shaders, binaryformat, binary, length);}
// static void  ptrShaderSource(GPSHADERSOURCE fnptr, GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length) {(*fnptr)(shader, count, string, length);}
// static void  ptrShaderStorageBlockBinding(GPSHADERSTORAGEBLOCKBINDING fnptr, GLuint  program, GLuint  storageBlockIndex, GLuint  storageBlockBinding) {(*fnptr)(program, storageBlockIndex, storageBlockBinding);}
// static void  ptrStencilFunc(GPSTENCILFUNC fnptr, GLenum  func, GLint  ref, GLuint  mask) {(*fnptr)(func, ref, mask);}
// static void  ptrStencilFuncSeparate(GPSTENCILFUNCSEPARATE fnptr, GLenum  face, GLenum  func, GLint  ref, GLuint  mask) {(*fnptr)(face, func, ref, mask);}
// static void  ptrStencilMask(GPSTENCILMASK fnptr, GLuint  mask) {(*fnptr)(mask);}
// static void  ptrStencilMaskSeparate(GPSTENCILMASKSEPARATE fnptr, GLenum  face, GLuint  mask) {(*fnptr)(face, mask);}
// static void  ptrStencilOp(GPSTENCILOP fnptr, GLenum  fail, GLenum  zfail, GLenum  zpass) {(*fnptr)(fail, zfail, zpass);}
// static void  ptrStencilOpSeparate(GPSTENCILOPSEPARATE fnptr, GLenum  face, GLenum  sfail, GLenum  dpfail, GLenum  dppass) {(*fnptr)(face, sfail, dpfail, dppass);}
// static void  ptrTexBuffer(GPTEXBUFFER fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer) {(*fnptr)(target, internalformat, buffer);}
// static void  ptrTexBufferRange(GPTEXBUFFERRANGE fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size) {(*fnptr)(target, internalformat, buffer, offset, size);}
// static void  ptrTexImage1D(GPTEXIMAGE1D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLint  border, GLenum  format, GLenum  type, const void * pixels) {(*fnptr)(target, level, internalformat, width, border, format, type, pixels);}
// static void  ptrTexImage2D(GPTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels) {(*fnptr)(target, level, internalformat, width, height, border, format, type, pixels);}
// static void  ptrTexImage2DMultisample(GPTEXIMAGE2DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {(*fnptr)(target, samples, internalformat, width, height, fixedsamplelocations);}
// static void  ptrTexImage3D(GPTEXIMAGE3D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels) {(*fnptr)(target, level, internalformat, width, height, depth, border, format, type, pixels);}
// static void  ptrTexImage3DMultisample(GPTEXIMAGE3DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {(*fnptr)(target, samples, internalformat, width, height, depth, fixedsamplelocations);}
// static void  ptrTexPageCommitmentARB(GPTEXPAGECOMMITMENTARB fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  resident) {(*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, resident);}
// static void  ptrTexParameterIiv(GPTEXPARAMETERIIV fnptr, GLenum  target, GLenum  pname, const GLint * params) {(*fnptr)(target, pname, params);}
// static void  ptrTexParameterIuiv(GPTEXPARAMETERIUIV fnptr, GLenum  target, GLenum  pname, const GLuint * params) {(*fnptr)(target, pname, params);}
// static void  ptrTexParameterf(GPTEXPARAMETERF fnptr, GLenum  target, GLenum  pname, GLfloat  param) {(*fnptr)(target, pname, param);}
// static void  ptrTexParameterfv(GPTEXPARAMETERFV fnptr, GLenum  target, GLenum  pname, const GLfloat * params) {(*fnptr)(target, pname, params);}
// static void  ptrTexParameteri(GPTEXPARAMETERI fnptr, GLenum  target, GLenum  pname, GLint  param) {(*fnptr)(target, pname, param);}
// static void  ptrTexParameteriv(GPTEXPARAMETERIV fnptr, GLenum  target, GLenum  pname, const GLint * params) {(*fnptr)(target, pname, params);}
// static void  ptrTexStorage1D(GPTEXSTORAGE1D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width) {(*fnptr)(target, levels, internalformat, width);}
// static void  ptrTexStorage2D(GPTEXSTORAGE2D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height) {(*fnptr)(target, levels, internalformat, width, height);}
// static void  ptrTexStorage2DMultisample(GPTEXSTORAGE2DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {(*fnptr)(target, samples, internalformat, width, height, fixedsamplelocations);}
// static void  ptrTexStorage3D(GPTEXSTORAGE3D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth) {(*fnptr)(target, levels, internalformat, width, height, depth);}
// static void  ptrTexStorage3DMultisample(GPTEXSTORAGE3DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {(*fnptr)(target, samples, internalformat, width, height, depth, fixedsamplelocations);}
// static void  ptrTexSubImage1D(GPTEXSUBIMAGE1D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels) {(*fnptr)(target, level, xoffset, width, format, type, pixels);}
// static void  ptrTexSubImage2D(GPTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels) {(*fnptr)(target, level, xoffset, yoffset, width, height, format, type, pixels);}
// static void  ptrTexSubImage3D(GPTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels) {(*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);}
// static void  ptrTextureBarrier(GPTEXTUREBARRIER fnptr) {(*fnptr)();}
// static void  ptrTextureBuffer(GPTEXTUREBUFFER fnptr, GLuint  texture, GLenum  internalformat, GLuint  buffer) {(*fnptr)(texture, internalformat, buffer);}
// static void  ptrTextureBufferRange(GPTEXTUREBUFFERRANGE fnptr, GLuint  texture, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizei  size) {(*fnptr)(texture, internalformat, buffer, offset, size);}
// static void  ptrTextureParameterIiv(GPTEXTUREPARAMETERIIV fnptr, GLuint  texture, GLenum  pname, const GLint * params) {(*fnptr)(texture, pname, params);}
// static void  ptrTextureParameterIuiv(GPTEXTUREPARAMETERIUIV fnptr, GLuint  texture, GLenum  pname, const GLuint * params) {(*fnptr)(texture, pname, params);}
// static void  ptrTextureParameterf(GPTEXTUREPARAMETERF fnptr, GLuint  texture, GLenum  pname, GLfloat  param) {(*fnptr)(texture, pname, param);}
// static void  ptrTextureParameterfv(GPTEXTUREPARAMETERFV fnptr, GLuint  texture, GLenum  pname, const GLfloat * param) {(*fnptr)(texture, pname, param);}
// static void  ptrTextureParameteri(GPTEXTUREPARAMETERI fnptr, GLuint  texture, GLenum  pname, GLint  param) {(*fnptr)(texture, pname, param);}
// static void  ptrTextureParameteriv(GPTEXTUREPARAMETERIV fnptr, GLuint  texture, GLenum  pname, const GLint * param) {(*fnptr)(texture, pname, param);}
// static void  ptrTextureStorage1D(GPTEXTURESTORAGE1D fnptr, GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width) {(*fnptr)(texture, levels, internalformat, width);}
// static void  ptrTextureStorage2D(GPTEXTURESTORAGE2D fnptr, GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height) {(*fnptr)(texture, levels, internalformat, width, height);}
// static void  ptrTextureStorage2DMultisample(GPTEXTURESTORAGE2DMULTISAMPLE fnptr, GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {(*fnptr)(texture, samples, internalformat, width, height, fixedsamplelocations);}
// static void  ptrTextureStorage3D(GPTEXTURESTORAGE3D fnptr, GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth) {(*fnptr)(texture, levels, internalformat, width, height, depth);}
// static void  ptrTextureStorage3DMultisample(GPTEXTURESTORAGE3DMULTISAMPLE fnptr, GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {(*fnptr)(texture, samples, internalformat, width, height, depth, fixedsamplelocations);}
// static void  ptrTextureSubImage1D(GPTEXTURESUBIMAGE1D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels) {(*fnptr)(texture, level, xoffset, width, format, type, pixels);}
// static void  ptrTextureSubImage2D(GPTEXTURESUBIMAGE2D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels) {(*fnptr)(texture, level, xoffset, yoffset, width, height, format, type, pixels);}
// static void  ptrTextureSubImage3D(GPTEXTURESUBIMAGE3D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels) {(*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);}
// static void  ptrTextureView(GPTEXTUREVIEW fnptr, GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers) {(*fnptr)(texture, target, origtexture, internalformat, minlevel, numlevels, minlayer, numlayers);}
// static void  ptrTransformFeedbackBufferBase(GPTRANSFORMFEEDBACKBUFFERBASE fnptr, GLuint  xfb, GLuint  index, GLuint  buffer) {(*fnptr)(xfb, index, buffer);}
// static void  ptrTransformFeedbackBufferRange(GPTRANSFORMFEEDBACKBUFFERRANGE fnptr, GLuint  xfb, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizei  size) {(*fnptr)(xfb, index, buffer, offset, size);}
// static void  ptrTransformFeedbackVaryings(GPTRANSFORMFEEDBACKVARYINGS fnptr, GLuint  program, GLsizei  count, const GLchar *const* varyings, GLenum  bufferMode) {(*fnptr)(program, count, varyings, bufferMode);}
// static void  ptrUniform1d(GPUNIFORM1D fnptr, GLint  location, GLdouble  x) {(*fnptr)(location, x);}
// static void  ptrUniform1dv(GPUNIFORM1DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform1f(GPUNIFORM1F fnptr, GLint  location, GLfloat  v0) {(*fnptr)(location, v0);}
// static void  ptrUniform1fv(GPUNIFORM1FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform1i(GPUNIFORM1I fnptr, GLint  location, GLint  v0) {(*fnptr)(location, v0);}
// static void  ptrUniform1iv(GPUNIFORM1IV fnptr, GLint  location, GLsizei  count, const GLint * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform1ui(GPUNIFORM1UI fnptr, GLint  location, GLuint  v0) {(*fnptr)(location, v0);}
// static void  ptrUniform1uiv(GPUNIFORM1UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform2d(GPUNIFORM2D fnptr, GLint  location, GLdouble  x, GLdouble  y) {(*fnptr)(location, x, y);}
// static void  ptrUniform2dv(GPUNIFORM2DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform2f(GPUNIFORM2F fnptr, GLint  location, GLfloat  v0, GLfloat  v1) {(*fnptr)(location, v0, v1);}
// static void  ptrUniform2fv(GPUNIFORM2FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform2i(GPUNIFORM2I fnptr, GLint  location, GLint  v0, GLint  v1) {(*fnptr)(location, v0, v1);}
// static void  ptrUniform2iv(GPUNIFORM2IV fnptr, GLint  location, GLsizei  count, const GLint * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform2ui(GPUNIFORM2UI fnptr, GLint  location, GLuint  v0, GLuint  v1) {(*fnptr)(location, v0, v1);}
// static void  ptrUniform2uiv(GPUNIFORM2UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform3d(GPUNIFORM3D fnptr, GLint  location, GLdouble  x, GLdouble  y, GLdouble  z) {(*fnptr)(location, x, y, z);}
// static void  ptrUniform3dv(GPUNIFORM3DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform3f(GPUNIFORM3F fnptr, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {(*fnptr)(location, v0, v1, v2);}
// static void  ptrUniform3fv(GPUNIFORM3FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform3i(GPUNIFORM3I fnptr, GLint  location, GLint  v0, GLint  v1, GLint  v2) {(*fnptr)(location, v0, v1, v2);}
// static void  ptrUniform3iv(GPUNIFORM3IV fnptr, GLint  location, GLsizei  count, const GLint * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform3ui(GPUNIFORM3UI fnptr, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {(*fnptr)(location, v0, v1, v2);}
// static void  ptrUniform3uiv(GPUNIFORM3UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform4d(GPUNIFORM4D fnptr, GLint  location, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {(*fnptr)(location, x, y, z, w);}
// static void  ptrUniform4dv(GPUNIFORM4DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform4f(GPUNIFORM4F fnptr, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {(*fnptr)(location, v0, v1, v2, v3);}
// static void  ptrUniform4fv(GPUNIFORM4FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform4i(GPUNIFORM4I fnptr, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {(*fnptr)(location, v0, v1, v2, v3);}
// static void  ptrUniform4iv(GPUNIFORM4IV fnptr, GLint  location, GLsizei  count, const GLint * value) {(*fnptr)(location, count, value);}
// static void  ptrUniform4ui(GPUNIFORM4UI fnptr, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {(*fnptr)(location, v0, v1, v2, v3);}
// static void  ptrUniform4uiv(GPUNIFORM4UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {(*fnptr)(location, count, value);}
// static void  ptrUniformBlockBinding(GPUNIFORMBLOCKBINDING fnptr, GLuint  program, GLuint  uniformBlockIndex, GLuint  uniformBlockBinding) {(*fnptr)(program, uniformBlockIndex, uniformBlockBinding);}
// static void  ptrUniformHandleui64ARB(GPUNIFORMHANDLEUI64ARB fnptr, GLint  location, GLuint64  value) {(*fnptr)(location, value);}
// static void  ptrUniformHandleui64vARB(GPUNIFORMHANDLEUI64VARB fnptr, GLint  location, GLsizei  count, const GLuint64 * value) {(*fnptr)(location, count, value);}
// static void  ptrUniformMatrix2dv(GPUNIFORMMATRIX2DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix2fv(GPUNIFORMMATRIX2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix2x3dv(GPUNIFORMMATRIX2X3DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix2x3fv(GPUNIFORMMATRIX2X3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix2x4dv(GPUNIFORMMATRIX2X4DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix2x4fv(GPUNIFORMMATRIX2X4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix3dv(GPUNIFORMMATRIX3DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix3fv(GPUNIFORMMATRIX3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix3x2dv(GPUNIFORMMATRIX3X2DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix3x2fv(GPUNIFORMMATRIX3X2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix3x4dv(GPUNIFORMMATRIX3X4DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix3x4fv(GPUNIFORMMATRIX3X4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix4dv(GPUNIFORMMATRIX4DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix4fv(GPUNIFORMMATRIX4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix4x2dv(GPUNIFORMMATRIX4X2DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix4x2fv(GPUNIFORMMATRIX4X2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix4x3dv(GPUNIFORMMATRIX4X3DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformMatrix4x3fv(GPUNIFORMMATRIX4X3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {(*fnptr)(location, count, transpose, value);}
// static void  ptrUniformSubroutinesuiv(GPUNIFORMSUBROUTINESUIV fnptr, GLenum  shadertype, GLsizei  count, const GLuint * indices) {(*fnptr)(shadertype, count, indices);}
// static GLboolean  ptrUnmapBuffer(GPUNMAPBUFFER fnptr, GLenum  target) {return (*fnptr)(target);}
// static GLboolean  ptrUnmapNamedBuffer(GPUNMAPNAMEDBUFFER fnptr, GLuint  buffer) {return (*fnptr)(buffer);}
// static void  ptrUseProgram(GPUSEPROGRAM fnptr, GLuint  program) {(*fnptr)(program);}
// static void  ptrUseProgramStages(GPUSEPROGRAMSTAGES fnptr, GLuint  pipeline, GLbitfield  stages, GLuint  program) {(*fnptr)(pipeline, stages, program);}
// static void  ptrValidateProgram(GPVALIDATEPROGRAM fnptr, GLuint  program) {(*fnptr)(program);}
// static void  ptrValidateProgramPipeline(GPVALIDATEPROGRAMPIPELINE fnptr, GLuint  pipeline) {(*fnptr)(pipeline);}
// static void  ptrVertexArrayAttribBinding(GPVERTEXARRAYATTRIBBINDING fnptr, GLuint  vaobj, GLuint  attribindex, GLuint  bindingindex) {(*fnptr)(vaobj, attribindex, bindingindex);}
// static void  ptrVertexArrayAttribFormat(GPVERTEXARRAYATTRIBFORMAT fnptr, GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset) {(*fnptr)(vaobj, attribindex, size, type, normalized, relativeoffset);}
// static void  ptrVertexArrayAttribIFormat(GPVERTEXARRAYATTRIBIFORMAT fnptr, GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {(*fnptr)(vaobj, attribindex, size, type, relativeoffset);}
// static void  ptrVertexArrayAttribLFormat(GPVERTEXARRAYATTRIBLFORMAT fnptr, GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {(*fnptr)(vaobj, attribindex, size, type, relativeoffset);}
// static void  ptrVertexArrayBindingDivisor(GPVERTEXARRAYBINDINGDIVISOR fnptr, GLuint  vaobj, GLuint  bindingindex, GLuint  divisor) {(*fnptr)(vaobj, bindingindex, divisor);}
// static void  ptrVertexArrayElementBuffer(GPVERTEXARRAYELEMENTBUFFER fnptr, GLuint  vaobj, GLuint  buffer) {(*fnptr)(vaobj, buffer);}
// static void  ptrVertexArrayVertexBuffer(GPVERTEXARRAYVERTEXBUFFER fnptr, GLuint  vaobj, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride) {(*fnptr)(vaobj, bindingindex, buffer, offset, stride);}
// static void  ptrVertexArrayVertexBuffers(GPVERTEXARRAYVERTEXBUFFERS fnptr, GLuint  vaobj, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides) {(*fnptr)(vaobj, first, count, buffers, offsets, strides);}
// static void  ptrVertexAttrib1d(GPVERTEXATTRIB1D fnptr, GLuint  index, GLdouble  x) {(*fnptr)(index, x);}
// static void  ptrVertexAttrib1dv(GPVERTEXATTRIB1DV fnptr, GLuint  index, const GLdouble * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib1f(GPVERTEXATTRIB1F fnptr, GLuint  index, GLfloat  x) {(*fnptr)(index, x);}
// static void  ptrVertexAttrib1fv(GPVERTEXATTRIB1FV fnptr, GLuint  index, const GLfloat * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib1s(GPVERTEXATTRIB1S fnptr, GLuint  index, GLshort  x) {(*fnptr)(index, x);}
// static void  ptrVertexAttrib1sv(GPVERTEXATTRIB1SV fnptr, GLuint  index, const GLshort * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib2d(GPVERTEXATTRIB2D fnptr, GLuint  index, GLdouble  x, GLdouble  y) {(*fnptr)(index, x, y);}
// static void  ptrVertexAttrib2dv(GPVERTEXATTRIB2DV fnptr, GLuint  index, const GLdouble * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib2f(GPVERTEXATTRIB2F fnptr, GLuint  index, GLfloat  x, GLfloat  y) {(*fnptr)(index, x, y);}
// static void  ptrVertexAttrib2fv(GPVERTEXATTRIB2FV fnptr, GLuint  index, const GLfloat * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib2s(GPVERTEXATTRIB2S fnptr, GLuint  index, GLshort  x, GLshort  y) {(*fnptr)(index, x, y);}
// static void  ptrVertexAttrib2sv(GPVERTEXATTRIB2SV fnptr, GLuint  index, const GLshort * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib3d(GPVERTEXATTRIB3D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z) {(*fnptr)(index, x, y, z);}
// static void  ptrVertexAttrib3dv(GPVERTEXATTRIB3DV fnptr, GLuint  index, const GLdouble * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib3f(GPVERTEXATTRIB3F fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z) {(*fnptr)(index, x, y, z);}
// static void  ptrVertexAttrib3fv(GPVERTEXATTRIB3FV fnptr, GLuint  index, const GLfloat * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib3s(GPVERTEXATTRIB3S fnptr, GLuint  index, GLshort  x, GLshort  y, GLshort  z) {(*fnptr)(index, x, y, z);}
// static void  ptrVertexAttrib3sv(GPVERTEXATTRIB3SV fnptr, GLuint  index, const GLshort * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4Nbv(GPVERTEXATTRIB4NBV fnptr, GLuint  index, const GLbyte * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4Niv(GPVERTEXATTRIB4NIV fnptr, GLuint  index, const GLint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4Nsv(GPVERTEXATTRIB4NSV fnptr, GLuint  index, const GLshort * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4Nub(GPVERTEXATTRIB4NUB fnptr, GLuint  index, GLubyte  x, GLubyte  y, GLubyte  z, GLubyte  w) {(*fnptr)(index, x, y, z, w);}
// static void  ptrVertexAttrib4Nubv(GPVERTEXATTRIB4NUBV fnptr, GLuint  index, const GLubyte * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4Nuiv(GPVERTEXATTRIB4NUIV fnptr, GLuint  index, const GLuint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4Nusv(GPVERTEXATTRIB4NUSV fnptr, GLuint  index, const GLushort * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4bv(GPVERTEXATTRIB4BV fnptr, GLuint  index, const GLbyte * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4d(GPVERTEXATTRIB4D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {(*fnptr)(index, x, y, z, w);}
// static void  ptrVertexAttrib4dv(GPVERTEXATTRIB4DV fnptr, GLuint  index, const GLdouble * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4f(GPVERTEXATTRIB4F fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w) {(*fnptr)(index, x, y, z, w);}
// static void  ptrVertexAttrib4fv(GPVERTEXATTRIB4FV fnptr, GLuint  index, const GLfloat * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4iv(GPVERTEXATTRIB4IV fnptr, GLuint  index, const GLint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4s(GPVERTEXATTRIB4S fnptr, GLuint  index, GLshort  x, GLshort  y, GLshort  z, GLshort  w) {(*fnptr)(index, x, y, z, w);}
// static void  ptrVertexAttrib4sv(GPVERTEXATTRIB4SV fnptr, GLuint  index, const GLshort * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4ubv(GPVERTEXATTRIB4UBV fnptr, GLuint  index, const GLubyte * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4uiv(GPVERTEXATTRIB4UIV fnptr, GLuint  index, const GLuint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttrib4usv(GPVERTEXATTRIB4USV fnptr, GLuint  index, const GLushort * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribBinding(GPVERTEXATTRIBBINDING fnptr, GLuint  attribindex, GLuint  bindingindex) {(*fnptr)(attribindex, bindingindex);}
// static void  ptrVertexAttribDivisor(GPVERTEXATTRIBDIVISOR fnptr, GLuint  index, GLuint  divisor) {(*fnptr)(index, divisor);}
// static void  ptrVertexAttribFormat(GPVERTEXATTRIBFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset) {(*fnptr)(attribindex, size, type, normalized, relativeoffset);}
// static void  ptrVertexAttribI1i(GPVERTEXATTRIBI1I fnptr, GLuint  index, GLint  x) {(*fnptr)(index, x);}
// static void  ptrVertexAttribI1iv(GPVERTEXATTRIBI1IV fnptr, GLuint  index, const GLint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI1ui(GPVERTEXATTRIBI1UI fnptr, GLuint  index, GLuint  x) {(*fnptr)(index, x);}
// static void  ptrVertexAttribI1uiv(GPVERTEXATTRIBI1UIV fnptr, GLuint  index, const GLuint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI2i(GPVERTEXATTRIBI2I fnptr, GLuint  index, GLint  x, GLint  y) {(*fnptr)(index, x, y);}
// static void  ptrVertexAttribI2iv(GPVERTEXATTRIBI2IV fnptr, GLuint  index, const GLint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI2ui(GPVERTEXATTRIBI2UI fnptr, GLuint  index, GLuint  x, GLuint  y) {(*fnptr)(index, x, y);}
// static void  ptrVertexAttribI2uiv(GPVERTEXATTRIBI2UIV fnptr, GLuint  index, const GLuint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI3i(GPVERTEXATTRIBI3I fnptr, GLuint  index, GLint  x, GLint  y, GLint  z) {(*fnptr)(index, x, y, z);}
// static void  ptrVertexAttribI3iv(GPVERTEXATTRIBI3IV fnptr, GLuint  index, const GLint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI3ui(GPVERTEXATTRIBI3UI fnptr, GLuint  index, GLuint  x, GLuint  y, GLuint  z) {(*fnptr)(index, x, y, z);}
// static void  ptrVertexAttribI3uiv(GPVERTEXATTRIBI3UIV fnptr, GLuint  index, const GLuint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI4bv(GPVERTEXATTRIBI4BV fnptr, GLuint  index, const GLbyte * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI4i(GPVERTEXATTRIBI4I fnptr, GLuint  index, GLint  x, GLint  y, GLint  z, GLint  w) {(*fnptr)(index, x, y, z, w);}
// static void  ptrVertexAttribI4iv(GPVERTEXATTRIBI4IV fnptr, GLuint  index, const GLint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI4sv(GPVERTEXATTRIBI4SV fnptr, GLuint  index, const GLshort * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI4ubv(GPVERTEXATTRIBI4UBV fnptr, GLuint  index, const GLubyte * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI4ui(GPVERTEXATTRIBI4UI fnptr, GLuint  index, GLuint  x, GLuint  y, GLuint  z, GLuint  w) {(*fnptr)(index, x, y, z, w);}
// static void  ptrVertexAttribI4uiv(GPVERTEXATTRIBI4UIV fnptr, GLuint  index, const GLuint * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribI4usv(GPVERTEXATTRIBI4USV fnptr, GLuint  index, const GLushort * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribIFormat(GPVERTEXATTRIBIFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {(*fnptr)(attribindex, size, type, relativeoffset);}
// static void  ptrVertexAttribIPointer(GPVERTEXATTRIBIPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {(*fnptr)(index, size, type, stride, pointer);}
// static void  ptrVertexAttribL1d(GPVERTEXATTRIBL1D fnptr, GLuint  index, GLdouble  x) {(*fnptr)(index, x);}
// static void  ptrVertexAttribL1dv(GPVERTEXATTRIBL1DV fnptr, GLuint  index, const GLdouble * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribL1ui64ARB(GPVERTEXATTRIBL1UI64ARB fnptr, GLuint  index, GLuint64EXT  x) {(*fnptr)(index, x);}
// static void  ptrVertexAttribL1ui64vARB(GPVERTEXATTRIBL1UI64VARB fnptr, GLuint  index, const GLuint64EXT * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribL2d(GPVERTEXATTRIBL2D fnptr, GLuint  index, GLdouble  x, GLdouble  y) {(*fnptr)(index, x, y);}
// static void  ptrVertexAttribL2dv(GPVERTEXATTRIBL2DV fnptr, GLuint  index, const GLdouble * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribL3d(GPVERTEXATTRIBL3D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z) {(*fnptr)(index, x, y, z);}
// static void  ptrVertexAttribL3dv(GPVERTEXATTRIBL3DV fnptr, GLuint  index, const GLdouble * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribL4d(GPVERTEXATTRIBL4D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {(*fnptr)(index, x, y, z, w);}
// static void  ptrVertexAttribL4dv(GPVERTEXATTRIBL4DV fnptr, GLuint  index, const GLdouble * v) {(*fnptr)(index, v);}
// static void  ptrVertexAttribLFormat(GPVERTEXATTRIBLFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {(*fnptr)(attribindex, size, type, relativeoffset);}
// static void  ptrVertexAttribLPointer(GPVERTEXATTRIBLPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {(*fnptr)(index, size, type, stride, pointer);}
// static void  ptrVertexAttribP1ui(GPVERTEXATTRIBP1UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {(*fnptr)(index, type, normalized, value);}
// static void  ptrVertexAttribP1uiv(GPVERTEXATTRIBP1UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {(*fnptr)(index, type, normalized, value);}
// static void  ptrVertexAttribP2ui(GPVERTEXATTRIBP2UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {(*fnptr)(index, type, normalized, value);}
// static void  ptrVertexAttribP2uiv(GPVERTEXATTRIBP2UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {(*fnptr)(index, type, normalized, value);}
// static void  ptrVertexAttribP3ui(GPVERTEXATTRIBP3UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {(*fnptr)(index, type, normalized, value);}
// static void  ptrVertexAttribP3uiv(GPVERTEXATTRIBP3UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {(*fnptr)(index, type, normalized, value);}
// static void  ptrVertexAttribP4ui(GPVERTEXATTRIBP4UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {(*fnptr)(index, type, normalized, value);}
// static void  ptrVertexAttribP4uiv(GPVERTEXATTRIBP4UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {(*fnptr)(index, type, normalized, value);}
// static void  ptrVertexAttribPointer(GPVERTEXATTRIBPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const void * pointer) {(*fnptr)(index, size, type, normalized, stride, pointer);}
// static void  ptrVertexBindingDivisor(GPVERTEXBINDINGDIVISOR fnptr, GLuint  bindingindex, GLuint  divisor) {(*fnptr)(bindingindex, divisor);}
// static void  ptrViewport(GPVIEWPORT fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {(*fnptr)(x, y, width, height);}
// static void  ptrViewportArrayv(GPVIEWPORTARRAYV fnptr, GLuint  first, GLsizei  count, const GLfloat * v) {(*fnptr)(first, count, v);}
// static void  ptrViewportIndexedf(GPVIEWPORTINDEXEDF fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h) {(*fnptr)(index, x, y, w, h);}
// static void  ptrViewportIndexedfv(GPVIEWPORTINDEXEDFV fnptr, GLuint  index, const GLfloat * v) {(*fnptr)(index, v);}
// static void  ptrWaitSync(GPWAITSYNC fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {(*fnptr)(sync, flags, timeout);}
import "C"
import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

var (
	gpActiveShaderProgram                         C.GPACTIVESHADERPROGRAM
	gpActiveTexture                               C.GPACTIVETEXTURE
	gpAttachShader                                C.GPATTACHSHADER
	gpBeginConditionalRender                      C.GPBEGINCONDITIONALRENDER
	gpBeginQuery                                  C.GPBEGINQUERY
	gpBeginQueryIndexed                           C.GPBEGINQUERYINDEXED
	gpBeginTransformFeedback                      C.GPBEGINTRANSFORMFEEDBACK
	gpBindAttribLocation                          C.GPBINDATTRIBLOCATION
	gpBindBuffer                                  C.GPBINDBUFFER
	gpBindBufferBase                              C.GPBINDBUFFERBASE
	gpBindBufferRange                             C.GPBINDBUFFERRANGE
	gpBindBuffersBase                             C.GPBINDBUFFERSBASE
	gpBindBuffersRange                            C.GPBINDBUFFERSRANGE
	gpBindFragDataLocation                        C.GPBINDFRAGDATALOCATION
	gpBindFragDataLocationIndexed                 C.GPBINDFRAGDATALOCATIONINDEXED
	gpBindFramebuffer                             C.GPBINDFRAMEBUFFER
	gpBindImageTexture                            C.GPBINDIMAGETEXTURE
	gpBindImageTextures                           C.GPBINDIMAGETEXTURES
	gpBindProgramPipeline                         C.GPBINDPROGRAMPIPELINE
	gpBindRenderbuffer                            C.GPBINDRENDERBUFFER
	gpBindSampler                                 C.GPBINDSAMPLER
	gpBindSamplers                                C.GPBINDSAMPLERS
	gpBindTexture                                 C.GPBINDTEXTURE
	gpBindTextureUnit                             C.GPBINDTEXTUREUNIT
	gpBindTextures                                C.GPBINDTEXTURES
	gpBindTransformFeedback                       C.GPBINDTRANSFORMFEEDBACK
	gpBindVertexArray                             C.GPBINDVERTEXARRAY
	gpBindVertexBuffer                            C.GPBINDVERTEXBUFFER
	gpBindVertexBuffers                           C.GPBINDVERTEXBUFFERS
	gpBlendColor                                  C.GPBLENDCOLOR
	gpBlendEquation                               C.GPBLENDEQUATION
	gpBlendEquationSeparate                       C.GPBLENDEQUATIONSEPARATE
	gpBlendEquationSeparatei                      C.GPBLENDEQUATIONSEPARATEI
	gpBlendEquationSeparateiARB                   C.GPBLENDEQUATIONSEPARATEIARB
	gpBlendEquationi                              C.GPBLENDEQUATIONI
	gpBlendEquationiARB                           C.GPBLENDEQUATIONIARB
	gpBlendFunc                                   C.GPBLENDFUNC
	gpBlendFuncSeparate                           C.GPBLENDFUNCSEPARATE
	gpBlendFuncSeparatei                          C.GPBLENDFUNCSEPARATEI
	gpBlendFuncSeparateiARB                       C.GPBLENDFUNCSEPARATEIARB
	gpBlendFunci                                  C.GPBLENDFUNCI
	gpBlendFunciARB                               C.GPBLENDFUNCIARB
	gpBlitFramebuffer                             C.GPBLITFRAMEBUFFER
	gpBlitNamedFramebuffer                        C.GPBLITNAMEDFRAMEBUFFER
	gpBufferData                                  C.GPBUFFERDATA
	gpBufferPageCommitmentARB                     C.GPBUFFERPAGECOMMITMENTARB
	gpBufferStorage                               C.GPBUFFERSTORAGE
	gpBufferSubData                               C.GPBUFFERSUBDATA
	gpCheckFramebufferStatus                      C.GPCHECKFRAMEBUFFERSTATUS
	gpCheckNamedFramebufferStatus                 C.GPCHECKNAMEDFRAMEBUFFERSTATUS
	gpClampColor                                  C.GPCLAMPCOLOR
	gpClear                                       C.GPCLEAR
	gpClearBufferData                             C.GPCLEARBUFFERDATA
	gpClearBufferSubData                          C.GPCLEARBUFFERSUBDATA
	gpClearBufferfi                               C.GPCLEARBUFFERFI
	gpClearBufferfv                               C.GPCLEARBUFFERFV
	gpClearBufferiv                               C.GPCLEARBUFFERIV
	gpClearBufferuiv                              C.GPCLEARBUFFERUIV
	gpClearColor                                  C.GPCLEARCOLOR
	gpClearDepth                                  C.GPCLEARDEPTH
	gpClearDepthf                                 C.GPCLEARDEPTHF
	gpClearNamedBufferData                        C.GPCLEARNAMEDBUFFERDATA
	gpClearNamedBufferSubData                     C.GPCLEARNAMEDBUFFERSUBDATA
	gpClearNamedFramebufferfi                     C.GPCLEARNAMEDFRAMEBUFFERFI
	gpClearNamedFramebufferfv                     C.GPCLEARNAMEDFRAMEBUFFERFV
	gpClearNamedFramebufferiv                     C.GPCLEARNAMEDFRAMEBUFFERIV
	gpClearNamedFramebufferuiv                    C.GPCLEARNAMEDFRAMEBUFFERUIV
	gpClearStencil                                C.GPCLEARSTENCIL
	gpClearTexImage                               C.GPCLEARTEXIMAGE
	gpClearTexSubImage                            C.GPCLEARTEXSUBIMAGE
	gpClientWaitSync                              C.GPCLIENTWAITSYNC
	gpClipControl                                 C.GPCLIPCONTROL
	gpColorMask                                   C.GPCOLORMASK
	gpColorMaski                                  C.GPCOLORMASKI
	gpCompileShader                               C.GPCOMPILESHADER
	gpCompileShaderIncludeARB                     C.GPCOMPILESHADERINCLUDEARB
	gpCompressedTexImage1D                        C.GPCOMPRESSEDTEXIMAGE1D
	gpCompressedTexImage2D                        C.GPCOMPRESSEDTEXIMAGE2D
	gpCompressedTexImage3D                        C.GPCOMPRESSEDTEXIMAGE3D
	gpCompressedTexSubImage1D                     C.GPCOMPRESSEDTEXSUBIMAGE1D
	gpCompressedTexSubImage2D                     C.GPCOMPRESSEDTEXSUBIMAGE2D
	gpCompressedTexSubImage3D                     C.GPCOMPRESSEDTEXSUBIMAGE3D
	gpCompressedTextureSubImage1D                 C.GPCOMPRESSEDTEXTURESUBIMAGE1D
	gpCompressedTextureSubImage2D                 C.GPCOMPRESSEDTEXTURESUBIMAGE2D
	gpCompressedTextureSubImage3D                 C.GPCOMPRESSEDTEXTURESUBIMAGE3D
	gpCopyBufferSubData                           C.GPCOPYBUFFERSUBDATA
	gpCopyImageSubData                            C.GPCOPYIMAGESUBDATA
	gpCopyNamedBufferSubData                      C.GPCOPYNAMEDBUFFERSUBDATA
	gpCopyTexImage1D                              C.GPCOPYTEXIMAGE1D
	gpCopyTexImage2D                              C.GPCOPYTEXIMAGE2D
	gpCopyTexSubImage1D                           C.GPCOPYTEXSUBIMAGE1D
	gpCopyTexSubImage2D                           C.GPCOPYTEXSUBIMAGE2D
	gpCopyTexSubImage3D                           C.GPCOPYTEXSUBIMAGE3D
	gpCopyTextureSubImage1D                       C.GPCOPYTEXTURESUBIMAGE1D
	gpCopyTextureSubImage2D                       C.GPCOPYTEXTURESUBIMAGE2D
	gpCopyTextureSubImage3D                       C.GPCOPYTEXTURESUBIMAGE3D
	gpCreateBuffers                               C.GPCREATEBUFFERS
	gpCreateFramebuffers                          C.GPCREATEFRAMEBUFFERS
	gpCreateProgram                               C.GPCREATEPROGRAM
	gpCreateProgramPipelines                      C.GPCREATEPROGRAMPIPELINES
	gpCreateQueries                               C.GPCREATEQUERIES
	gpCreateRenderbuffers                         C.GPCREATERENDERBUFFERS
	gpCreateSamplers                              C.GPCREATESAMPLERS
	gpCreateShader                                C.GPCREATESHADER
	gpCreateShaderProgramv                        C.GPCREATESHADERPROGRAMV
	gpCreateSyncFromCLeventARB                    C.GPCREATESYNCFROMCLEVENTARB
	gpCreateTextures                              C.GPCREATETEXTURES
	gpCreateTransformFeedbacks                    C.GPCREATETRANSFORMFEEDBACKS
	gpCreateVertexArrays                          C.GPCREATEVERTEXARRAYS
	gpCullFace                                    C.GPCULLFACE
	gpDebugMessageCallback                        C.GPDEBUGMESSAGECALLBACK
	gpDebugMessageCallbackARB                     C.GPDEBUGMESSAGECALLBACKARB
	gpDebugMessageCallbackKHR                     C.GPDEBUGMESSAGECALLBACKKHR
	gpDebugMessageControl                         C.GPDEBUGMESSAGECONTROL
	gpDebugMessageControlARB                      C.GPDEBUGMESSAGECONTROLARB
	gpDebugMessageControlKHR                      C.GPDEBUGMESSAGECONTROLKHR
	gpDebugMessageInsert                          C.GPDEBUGMESSAGEINSERT
	gpDebugMessageInsertARB                       C.GPDEBUGMESSAGEINSERTARB
	gpDebugMessageInsertKHR                       C.GPDEBUGMESSAGEINSERTKHR
	gpDeleteBuffers                               C.GPDELETEBUFFERS
	gpDeleteFramebuffers                          C.GPDELETEFRAMEBUFFERS
	gpDeleteNamedStringARB                        C.GPDELETENAMEDSTRINGARB
	gpDeleteProgram                               C.GPDELETEPROGRAM
	gpDeleteProgramPipelines                      C.GPDELETEPROGRAMPIPELINES
	gpDeleteQueries                               C.GPDELETEQUERIES
	gpDeleteRenderbuffers                         C.GPDELETERENDERBUFFERS
	gpDeleteSamplers                              C.GPDELETESAMPLERS
	gpDeleteShader                                C.GPDELETESHADER
	gpDeleteSync                                  C.GPDELETESYNC
	gpDeleteTextures                              C.GPDELETETEXTURES
	gpDeleteTransformFeedbacks                    C.GPDELETETRANSFORMFEEDBACKS
	gpDeleteVertexArrays                          C.GPDELETEVERTEXARRAYS
	gpDepthFunc                                   C.GPDEPTHFUNC
	gpDepthMask                                   C.GPDEPTHMASK
	gpDepthRange                                  C.GPDEPTHRANGE
	gpDepthRangeArrayv                            C.GPDEPTHRANGEARRAYV
	gpDepthRangeIndexed                           C.GPDEPTHRANGEINDEXED
	gpDepthRangef                                 C.GPDEPTHRANGEF
	gpDetachShader                                C.GPDETACHSHADER
	gpDisable                                     C.GPDISABLE
	gpDisableVertexArrayAttrib                    C.GPDISABLEVERTEXARRAYATTRIB
	gpDisableVertexAttribArray                    C.GPDISABLEVERTEXATTRIBARRAY
	gpDisablei                                    C.GPDISABLEI
	gpDispatchCompute                             C.GPDISPATCHCOMPUTE
	gpDispatchComputeGroupSizeARB                 C.GPDISPATCHCOMPUTEGROUPSIZEARB
	gpDispatchComputeIndirect                     C.GPDISPATCHCOMPUTEINDIRECT
	gpDrawArrays                                  C.GPDRAWARRAYS
	gpDrawArraysIndirect                          C.GPDRAWARRAYSINDIRECT
	gpDrawArraysInstanced                         C.GPDRAWARRAYSINSTANCED
	gpDrawArraysInstancedBaseInstance             C.GPDRAWARRAYSINSTANCEDBASEINSTANCE
	gpDrawBuffer                                  C.GPDRAWBUFFER
	gpDrawBuffers                                 C.GPDRAWBUFFERS
	gpDrawElements                                C.GPDRAWELEMENTS
	gpDrawElementsBaseVertex                      C.GPDRAWELEMENTSBASEVERTEX
	gpDrawElementsIndirect                        C.GPDRAWELEMENTSINDIRECT
	gpDrawElementsInstanced                       C.GPDRAWELEMENTSINSTANCED
	gpDrawElementsInstancedBaseInstance           C.GPDRAWELEMENTSINSTANCEDBASEINSTANCE
	gpDrawElementsInstancedBaseVertex             C.GPDRAWELEMENTSINSTANCEDBASEVERTEX
	gpDrawElementsInstancedBaseVertexBaseInstance C.GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE
	gpDrawRangeElements                           C.GPDRAWRANGEELEMENTS
	gpDrawRangeElementsBaseVertex                 C.GPDRAWRANGEELEMENTSBASEVERTEX
	gpDrawTransformFeedback                       C.GPDRAWTRANSFORMFEEDBACK
	gpDrawTransformFeedbackInstanced              C.GPDRAWTRANSFORMFEEDBACKINSTANCED
	gpDrawTransformFeedbackStream                 C.GPDRAWTRANSFORMFEEDBACKSTREAM
	gpDrawTransformFeedbackStreamInstanced        C.GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED
	gpEnable                                      C.GPENABLE
	gpEnableVertexArrayAttrib                     C.GPENABLEVERTEXARRAYATTRIB
	gpEnableVertexAttribArray                     C.GPENABLEVERTEXATTRIBARRAY
	gpEnablei                                     C.GPENABLEI
	gpEndConditionalRender                        C.GPENDCONDITIONALRENDER
	gpEndQuery                                    C.GPENDQUERY
	gpEndQueryIndexed                             C.GPENDQUERYINDEXED
	gpEndTransformFeedback                        C.GPENDTRANSFORMFEEDBACK
	gpFenceSync                                   C.GPFENCESYNC
	gpFinish                                      C.GPFINISH
	gpFlush                                       C.GPFLUSH
	gpFlushMappedBufferRange                      C.GPFLUSHMAPPEDBUFFERRANGE
	gpFlushMappedNamedBufferRange                 C.GPFLUSHMAPPEDNAMEDBUFFERRANGE
	gpFramebufferParameteri                       C.GPFRAMEBUFFERPARAMETERI
	gpFramebufferRenderbuffer                     C.GPFRAMEBUFFERRENDERBUFFER
	gpFramebufferTexture                          C.GPFRAMEBUFFERTEXTURE
	gpFramebufferTexture1D                        C.GPFRAMEBUFFERTEXTURE1D
	gpFramebufferTexture2D                        C.GPFRAMEBUFFERTEXTURE2D
	gpFramebufferTexture3D                        C.GPFRAMEBUFFERTEXTURE3D
	gpFramebufferTextureLayer                     C.GPFRAMEBUFFERTEXTURELAYER
	gpFrontFace                                   C.GPFRONTFACE
	gpGenBuffers                                  C.GPGENBUFFERS
	gpGenFramebuffers                             C.GPGENFRAMEBUFFERS
	gpGenProgramPipelines                         C.GPGENPROGRAMPIPELINES
	gpGenQueries                                  C.GPGENQUERIES
	gpGenRenderbuffers                            C.GPGENRENDERBUFFERS
	gpGenSamplers                                 C.GPGENSAMPLERS
	gpGenTextures                                 C.GPGENTEXTURES
	gpGenTransformFeedbacks                       C.GPGENTRANSFORMFEEDBACKS
	gpGenVertexArrays                             C.GPGENVERTEXARRAYS
	gpGenerateMipmap                              C.GPGENERATEMIPMAP
	gpGenerateTextureMipmap                       C.GPGENERATETEXTUREMIPMAP
	gpGetActiveAtomicCounterBufferiv              C.GPGETACTIVEATOMICCOUNTERBUFFERIV
	gpGetActiveAttrib                             C.GPGETACTIVEATTRIB
	gpGetActiveSubroutineName                     C.GPGETACTIVESUBROUTINENAME
	gpGetActiveSubroutineUniformName              C.GPGETACTIVESUBROUTINEUNIFORMNAME
	gpGetActiveSubroutineUniformiv                C.GPGETACTIVESUBROUTINEUNIFORMIV
	gpGetActiveUniform                            C.GPGETACTIVEUNIFORM
	gpGetActiveUniformBlockName                   C.GPGETACTIVEUNIFORMBLOCKNAME
	gpGetActiveUniformBlockiv                     C.GPGETACTIVEUNIFORMBLOCKIV
	gpGetActiveUniformName                        C.GPGETACTIVEUNIFORMNAME
	gpGetActiveUniformsiv                         C.GPGETACTIVEUNIFORMSIV
	gpGetAttachedShaders                          C.GPGETATTACHEDSHADERS
	gpGetAttribLocation                           C.GPGETATTRIBLOCATION
	gpGetBooleani_v                               C.GPGETBOOLEANI_V
	gpGetBooleanv                                 C.GPGETBOOLEANV
	gpGetBufferParameteri64v                      C.GPGETBUFFERPARAMETERI64V
	gpGetBufferParameteriv                        C.GPGETBUFFERPARAMETERIV
	gpGetBufferPointerv                           C.GPGETBUFFERPOINTERV
	gpGetBufferSubData                            C.GPGETBUFFERSUBDATA
	gpGetCompressedTexImage                       C.GPGETCOMPRESSEDTEXIMAGE
	gpGetCompressedTextureImage                   C.GPGETCOMPRESSEDTEXTUREIMAGE
	gpGetCompressedTextureSubImage                C.GPGETCOMPRESSEDTEXTURESUBIMAGE
	gpGetDebugMessageLog                          C.GPGETDEBUGMESSAGELOG
	gpGetDebugMessageLogARB                       C.GPGETDEBUGMESSAGELOGARB
	gpGetDebugMessageLogKHR                       C.GPGETDEBUGMESSAGELOGKHR
	gpGetDoublei_v                                C.GPGETDOUBLEI_V
	gpGetDoublev                                  C.GPGETDOUBLEV
	gpGetError                                    C.GPGETERROR
	gpGetFloati_v                                 C.GPGETFLOATI_V
	gpGetFloatv                                   C.GPGETFLOATV
	gpGetFragDataIndex                            C.GPGETFRAGDATAINDEX
	gpGetFragDataLocation                         C.GPGETFRAGDATALOCATION
	gpGetFramebufferAttachmentParameteriv         C.GPGETFRAMEBUFFERATTACHMENTPARAMETERIV
	gpGetFramebufferParameteriv                   C.GPGETFRAMEBUFFERPARAMETERIV
	gpGetGraphicsResetStatus                      C.GPGETGRAPHICSRESETSTATUS
	gpGetGraphicsResetStatusARB                   C.GPGETGRAPHICSRESETSTATUSARB
	gpGetGraphicsResetStatusKHR                   C.GPGETGRAPHICSRESETSTATUSKHR
	gpGetImageHandleARB                           C.GPGETIMAGEHANDLEARB
	gpGetInteger64i_v                             C.GPGETINTEGER64I_V
	gpGetInteger64v                               C.GPGETINTEGER64V
	gpGetIntegeri_v                               C.GPGETINTEGERI_V
	gpGetIntegerv                                 C.GPGETINTEGERV
	gpGetInternalformati64v                       C.GPGETINTERNALFORMATI64V
	gpGetInternalformativ                         C.GPGETINTERNALFORMATIV
	gpGetMultisamplefv                            C.GPGETMULTISAMPLEFV
	gpGetNamedBufferParameteri64v                 C.GPGETNAMEDBUFFERPARAMETERI64V
	gpGetNamedBufferParameteriv                   C.GPGETNAMEDBUFFERPARAMETERIV
	gpGetNamedBufferPointerv                      C.GPGETNAMEDBUFFERPOINTERV
	gpGetNamedBufferSubData                       C.GPGETNAMEDBUFFERSUBDATA
	gpGetNamedFramebufferAttachmentParameteriv    C.GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV
	gpGetNamedFramebufferParameteriv              C.GPGETNAMEDFRAMEBUFFERPARAMETERIV
	gpGetNamedRenderbufferParameteriv             C.GPGETNAMEDRENDERBUFFERPARAMETERIV
	gpGetNamedStringARB                           C.GPGETNAMEDSTRINGARB
	gpGetNamedStringivARB                         C.GPGETNAMEDSTRINGIVARB
	gpGetObjectLabel                              C.GPGETOBJECTLABEL
	gpGetObjectLabelKHR                           C.GPGETOBJECTLABELKHR
	gpGetObjectPtrLabel                           C.GPGETOBJECTPTRLABEL
	gpGetObjectPtrLabelKHR                        C.GPGETOBJECTPTRLABELKHR
	gpGetPointerv                                 C.GPGETPOINTERV
	gpGetPointervKHR                              C.GPGETPOINTERVKHR
	gpGetProgramBinary                            C.GPGETPROGRAMBINARY
	gpGetProgramInfoLog                           C.GPGETPROGRAMINFOLOG
	gpGetProgramInterfaceiv                       C.GPGETPROGRAMINTERFACEIV
	gpGetProgramPipelineInfoLog                   C.GPGETPROGRAMPIPELINEINFOLOG
	gpGetProgramPipelineiv                        C.GPGETPROGRAMPIPELINEIV
	gpGetProgramResourceIndex                     C.GPGETPROGRAMRESOURCEINDEX
	gpGetProgramResourceLocation                  C.GPGETPROGRAMRESOURCELOCATION
	gpGetProgramResourceLocationIndex             C.GPGETPROGRAMRESOURCELOCATIONINDEX
	gpGetProgramResourceName                      C.GPGETPROGRAMRESOURCENAME
	gpGetProgramResourceiv                        C.GPGETPROGRAMRESOURCEIV
	gpGetProgramStageiv                           C.GPGETPROGRAMSTAGEIV
	gpGetProgramiv                                C.GPGETPROGRAMIV
	gpGetQueryIndexediv                           C.GPGETQUERYINDEXEDIV
	gpGetQueryObjecti64v                          C.GPGETQUERYOBJECTI64V
	gpGetQueryObjectiv                            C.GPGETQUERYOBJECTIV
	gpGetQueryObjectui64v                         C.GPGETQUERYOBJECTUI64V
	gpGetQueryObjectuiv                           C.GPGETQUERYOBJECTUIV
	gpGetQueryiv                                  C.GPGETQUERYIV
	gpGetRenderbufferParameteriv                  C.GPGETRENDERBUFFERPARAMETERIV
	gpGetSamplerParameterIiv                      C.GPGETSAMPLERPARAMETERIIV
	gpGetSamplerParameterIuiv                     C.GPGETSAMPLERPARAMETERIUIV
	gpGetSamplerParameterfv                       C.GPGETSAMPLERPARAMETERFV
	gpGetSamplerParameteriv                       C.GPGETSAMPLERPARAMETERIV
	gpGetShaderInfoLog                            C.GPGETSHADERINFOLOG
	gpGetShaderPrecisionFormat                    C.GPGETSHADERPRECISIONFORMAT
	gpGetShaderSource                             C.GPGETSHADERSOURCE
	gpGetShaderiv                                 C.GPGETSHADERIV
	gpGetString                                   C.GPGETSTRING
	gpGetStringi                                  C.GPGETSTRINGI
	gpGetSubroutineIndex                          C.GPGETSUBROUTINEINDEX
	gpGetSubroutineUniformLocation                C.GPGETSUBROUTINEUNIFORMLOCATION
	gpGetSynciv                                   C.GPGETSYNCIV
	gpGetTexImage                                 C.GPGETTEXIMAGE
	gpGetTexLevelParameterfv                      C.GPGETTEXLEVELPARAMETERFV
	gpGetTexLevelParameteriv                      C.GPGETTEXLEVELPARAMETERIV
	gpGetTexParameterIiv                          C.GPGETTEXPARAMETERIIV
	gpGetTexParameterIuiv                         C.GPGETTEXPARAMETERIUIV
	gpGetTexParameterfv                           C.GPGETTEXPARAMETERFV
	gpGetTexParameteriv                           C.GPGETTEXPARAMETERIV
	gpGetTextureHandleARB                         C.GPGETTEXTUREHANDLEARB
	gpGetTextureImage                             C.GPGETTEXTUREIMAGE
	gpGetTextureLevelParameterfv                  C.GPGETTEXTURELEVELPARAMETERFV
	gpGetTextureLevelParameteriv                  C.GPGETTEXTURELEVELPARAMETERIV
	gpGetTextureParameterIiv                      C.GPGETTEXTUREPARAMETERIIV
	gpGetTextureParameterIuiv                     C.GPGETTEXTUREPARAMETERIUIV
	gpGetTextureParameterfv                       C.GPGETTEXTUREPARAMETERFV
	gpGetTextureParameteriv                       C.GPGETTEXTUREPARAMETERIV
	gpGetTextureSamplerHandleARB                  C.GPGETTEXTURESAMPLERHANDLEARB
	gpGetTextureSubImage                          C.GPGETTEXTURESUBIMAGE
	gpGetTransformFeedbackVarying                 C.GPGETTRANSFORMFEEDBACKVARYING
	gpGetTransformFeedbacki64_v                   C.GPGETTRANSFORMFEEDBACKI64_V
	gpGetTransformFeedbacki_v                     C.GPGETTRANSFORMFEEDBACKI_V
	gpGetTransformFeedbackiv                      C.GPGETTRANSFORMFEEDBACKIV
	gpGetUniformBlockIndex                        C.GPGETUNIFORMBLOCKINDEX
	gpGetUniformIndices                           C.GPGETUNIFORMINDICES
	gpGetUniformLocation                          C.GPGETUNIFORMLOCATION
	gpGetUniformSubroutineuiv                     C.GPGETUNIFORMSUBROUTINEUIV
	gpGetUniformdv                                C.GPGETUNIFORMDV
	gpGetUniformfv                                C.GPGETUNIFORMFV
	gpGetUniformiv                                C.GPGETUNIFORMIV
	gpGetUniformuiv                               C.GPGETUNIFORMUIV
	gpGetVertexArrayIndexed64iv                   C.GPGETVERTEXARRAYINDEXED64IV
	gpGetVertexArrayIndexediv                     C.GPGETVERTEXARRAYINDEXEDIV
	gpGetVertexArrayiv                            C.GPGETVERTEXARRAYIV
	gpGetVertexAttribIiv                          C.GPGETVERTEXATTRIBIIV
	gpGetVertexAttribIuiv                         C.GPGETVERTEXATTRIBIUIV
	gpGetVertexAttribLdv                          C.GPGETVERTEXATTRIBLDV
	gpGetVertexAttribLui64vARB                    C.GPGETVERTEXATTRIBLUI64VARB
	gpGetVertexAttribPointerv                     C.GPGETVERTEXATTRIBPOINTERV
	gpGetVertexAttribdv                           C.GPGETVERTEXATTRIBDV
	gpGetVertexAttribfv                           C.GPGETVERTEXATTRIBFV
	gpGetVertexAttribiv                           C.GPGETVERTEXATTRIBIV
	gpGetnCompressedTexImage                      C.GPGETNCOMPRESSEDTEXIMAGE
	gpGetnCompressedTexImageARB                   C.GPGETNCOMPRESSEDTEXIMAGEARB
	gpGetnTexImage                                C.GPGETNTEXIMAGE
	gpGetnTexImageARB                             C.GPGETNTEXIMAGEARB
	gpGetnUniformdv                               C.GPGETNUNIFORMDV
	gpGetnUniformdvARB                            C.GPGETNUNIFORMDVARB
	gpGetnUniformfv                               C.GPGETNUNIFORMFV
	gpGetnUniformfvARB                            C.GPGETNUNIFORMFVARB
	gpGetnUniformfvKHR                            C.GPGETNUNIFORMFVKHR
	gpGetnUniformiv                               C.GPGETNUNIFORMIV
	gpGetnUniformivARB                            C.GPGETNUNIFORMIVARB
	gpGetnUniformivKHR                            C.GPGETNUNIFORMIVKHR
	gpGetnUniformuiv                              C.GPGETNUNIFORMUIV
	gpGetnUniformuivARB                           C.GPGETNUNIFORMUIVARB
	gpGetnUniformuivKHR                           C.GPGETNUNIFORMUIVKHR
	gpHint                                        C.GPHINT
	gpInvalidateBufferData                        C.GPINVALIDATEBUFFERDATA
	gpInvalidateBufferSubData                     C.GPINVALIDATEBUFFERSUBDATA
	gpInvalidateFramebuffer                       C.GPINVALIDATEFRAMEBUFFER
	gpInvalidateNamedFramebufferData              C.GPINVALIDATENAMEDFRAMEBUFFERDATA
	gpInvalidateNamedFramebufferSubData           C.GPINVALIDATENAMEDFRAMEBUFFERSUBDATA
	gpInvalidateSubFramebuffer                    C.GPINVALIDATESUBFRAMEBUFFER
	gpInvalidateTexImage                          C.GPINVALIDATETEXIMAGE
	gpInvalidateTexSubImage                       C.GPINVALIDATETEXSUBIMAGE
	gpIsBuffer                                    C.GPISBUFFER
	gpIsEnabled                                   C.GPISENABLED
	gpIsEnabledi                                  C.GPISENABLEDI
	gpIsFramebuffer                               C.GPISFRAMEBUFFER
	gpIsImageHandleResidentARB                    C.GPISIMAGEHANDLERESIDENTARB
	gpIsNamedStringARB                            C.GPISNAMEDSTRINGARB
	gpIsProgram                                   C.GPISPROGRAM
	gpIsProgramPipeline                           C.GPISPROGRAMPIPELINE
	gpIsQuery                                     C.GPISQUERY
	gpIsRenderbuffer                              C.GPISRENDERBUFFER
	gpIsSampler                                   C.GPISSAMPLER
	gpIsShader                                    C.GPISSHADER
	gpIsSync                                      C.GPISSYNC
	gpIsTexture                                   C.GPISTEXTURE
	gpIsTextureHandleResidentARB                  C.GPISTEXTUREHANDLERESIDENTARB
	gpIsTransformFeedback                         C.GPISTRANSFORMFEEDBACK
	gpIsVertexArray                               C.GPISVERTEXARRAY
	gpLineWidth                                   C.GPLINEWIDTH
	gpLinkProgram                                 C.GPLINKPROGRAM
	gpLogicOp                                     C.GPLOGICOP
	gpMakeImageHandleNonResidentARB               C.GPMAKEIMAGEHANDLENONRESIDENTARB
	gpMakeImageHandleResidentARB                  C.GPMAKEIMAGEHANDLERESIDENTARB
	gpMakeTextureHandleNonResidentARB             C.GPMAKETEXTUREHANDLENONRESIDENTARB
	gpMakeTextureHandleResidentARB                C.GPMAKETEXTUREHANDLERESIDENTARB
	gpMapBuffer                                   C.GPMAPBUFFER
	gpMapBufferRange                              C.GPMAPBUFFERRANGE
	gpMapNamedBuffer                              C.GPMAPNAMEDBUFFER
	gpMapNamedBufferRange                         C.GPMAPNAMEDBUFFERRANGE
	gpMemoryBarrier                               C.GPMEMORYBARRIER
	gpMemoryBarrierByRegion                       C.GPMEMORYBARRIERBYREGION
	gpMinSampleShading                            C.GPMINSAMPLESHADING
	gpMinSampleShadingARB                         C.GPMINSAMPLESHADINGARB
	gpMultiDrawArrays                             C.GPMULTIDRAWARRAYS
	gpMultiDrawArraysIndirect                     C.GPMULTIDRAWARRAYSINDIRECT
	gpMultiDrawArraysIndirectCountARB             C.GPMULTIDRAWARRAYSINDIRECTCOUNTARB
	gpMultiDrawElements                           C.GPMULTIDRAWELEMENTS
	gpMultiDrawElementsBaseVertex                 C.GPMULTIDRAWELEMENTSBASEVERTEX
	gpMultiDrawElementsIndirect                   C.GPMULTIDRAWELEMENTSINDIRECT
	gpMultiDrawElementsIndirectCountARB           C.GPMULTIDRAWELEMENTSINDIRECTCOUNTARB
	gpNamedBufferData                             C.GPNAMEDBUFFERDATA
	gpNamedBufferPageCommitmentARB                C.GPNAMEDBUFFERPAGECOMMITMENTARB
	gpNamedBufferPageCommitmentEXT                C.GPNAMEDBUFFERPAGECOMMITMENTEXT
	gpNamedBufferStorage                          C.GPNAMEDBUFFERSTORAGE
	gpNamedBufferSubData                          C.GPNAMEDBUFFERSUBDATA
	gpNamedFramebufferDrawBuffer                  C.GPNAMEDFRAMEBUFFERDRAWBUFFER
	gpNamedFramebufferDrawBuffers                 C.GPNAMEDFRAMEBUFFERDRAWBUFFERS
	gpNamedFramebufferParameteri                  C.GPNAMEDFRAMEBUFFERPARAMETERI
	gpNamedFramebufferReadBuffer                  C.GPNAMEDFRAMEBUFFERREADBUFFER
	gpNamedFramebufferRenderbuffer                C.GPNAMEDFRAMEBUFFERRENDERBUFFER
	gpNamedFramebufferTexture                     C.GPNAMEDFRAMEBUFFERTEXTURE
	gpNamedFramebufferTextureLayer                C.GPNAMEDFRAMEBUFFERTEXTURELAYER
	gpNamedRenderbufferStorage                    C.GPNAMEDRENDERBUFFERSTORAGE
	gpNamedRenderbufferStorageMultisample         C.GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE
	gpNamedStringARB                              C.GPNAMEDSTRINGARB
	gpObjectLabel                                 C.GPOBJECTLABEL
	gpObjectLabelKHR                              C.GPOBJECTLABELKHR
	gpObjectPtrLabel                              C.GPOBJECTPTRLABEL
	gpObjectPtrLabelKHR                           C.GPOBJECTPTRLABELKHR
	gpPatchParameterfv                            C.GPPATCHPARAMETERFV
	gpPatchParameteri                             C.GPPATCHPARAMETERI
	gpPauseTransformFeedback                      C.GPPAUSETRANSFORMFEEDBACK
	gpPixelStoref                                 C.GPPIXELSTOREF
	gpPixelStorei                                 C.GPPIXELSTOREI
	gpPointParameterf                             C.GPPOINTPARAMETERF
	gpPointParameterfv                            C.GPPOINTPARAMETERFV
	gpPointParameteri                             C.GPPOINTPARAMETERI
	gpPointParameteriv                            C.GPPOINTPARAMETERIV
	gpPointSize                                   C.GPPOINTSIZE
	gpPolygonMode                                 C.GPPOLYGONMODE
	gpPolygonOffset                               C.GPPOLYGONOFFSET
	gpPopDebugGroup                               C.GPPOPDEBUGGROUP
	gpPopDebugGroupKHR                            C.GPPOPDEBUGGROUPKHR
	gpPrimitiveRestartIndex                       C.GPPRIMITIVERESTARTINDEX
	gpProgramBinary                               C.GPPROGRAMBINARY
	gpProgramParameteri                           C.GPPROGRAMPARAMETERI
	gpProgramUniform1d                            C.GPPROGRAMUNIFORM1D
	gpProgramUniform1dv                           C.GPPROGRAMUNIFORM1DV
	gpProgramUniform1f                            C.GPPROGRAMUNIFORM1F
	gpProgramUniform1fv                           C.GPPROGRAMUNIFORM1FV
	gpProgramUniform1i                            C.GPPROGRAMUNIFORM1I
	gpProgramUniform1iv                           C.GPPROGRAMUNIFORM1IV
	gpProgramUniform1ui                           C.GPPROGRAMUNIFORM1UI
	gpProgramUniform1uiv                          C.GPPROGRAMUNIFORM1UIV
	gpProgramUniform2d                            C.GPPROGRAMUNIFORM2D
	gpProgramUniform2dv                           C.GPPROGRAMUNIFORM2DV
	gpProgramUniform2f                            C.GPPROGRAMUNIFORM2F
	gpProgramUniform2fv                           C.GPPROGRAMUNIFORM2FV
	gpProgramUniform2i                            C.GPPROGRAMUNIFORM2I
	gpProgramUniform2iv                           C.GPPROGRAMUNIFORM2IV
	gpProgramUniform2ui                           C.GPPROGRAMUNIFORM2UI
	gpProgramUniform2uiv                          C.GPPROGRAMUNIFORM2UIV
	gpProgramUniform3d                            C.GPPROGRAMUNIFORM3D
	gpProgramUniform3dv                           C.GPPROGRAMUNIFORM3DV
	gpProgramUniform3f                            C.GPPROGRAMUNIFORM3F
	gpProgramUniform3fv                           C.GPPROGRAMUNIFORM3FV
	gpProgramUniform3i                            C.GPPROGRAMUNIFORM3I
	gpProgramUniform3iv                           C.GPPROGRAMUNIFORM3IV
	gpProgramUniform3ui                           C.GPPROGRAMUNIFORM3UI
	gpProgramUniform3uiv                          C.GPPROGRAMUNIFORM3UIV
	gpProgramUniform4d                            C.GPPROGRAMUNIFORM4D
	gpProgramUniform4dv                           C.GPPROGRAMUNIFORM4DV
	gpProgramUniform4f                            C.GPPROGRAMUNIFORM4F
	gpProgramUniform4fv                           C.GPPROGRAMUNIFORM4FV
	gpProgramUniform4i                            C.GPPROGRAMUNIFORM4I
	gpProgramUniform4iv                           C.GPPROGRAMUNIFORM4IV
	gpProgramUniform4ui                           C.GPPROGRAMUNIFORM4UI
	gpProgramUniform4uiv                          C.GPPROGRAMUNIFORM4UIV
	gpProgramUniformHandleui64ARB                 C.GPPROGRAMUNIFORMHANDLEUI64ARB
	gpProgramUniformHandleui64vARB                C.GPPROGRAMUNIFORMHANDLEUI64VARB
	gpProgramUniformMatrix2dv                     C.GPPROGRAMUNIFORMMATRIX2DV
	gpProgramUniformMatrix2fv                     C.GPPROGRAMUNIFORMMATRIX2FV
	gpProgramUniformMatrix2x3dv                   C.GPPROGRAMUNIFORMMATRIX2X3DV
	gpProgramUniformMatrix2x3fv                   C.GPPROGRAMUNIFORMMATRIX2X3FV
	gpProgramUniformMatrix2x4dv                   C.GPPROGRAMUNIFORMMATRIX2X4DV
	gpProgramUniformMatrix2x4fv                   C.GPPROGRAMUNIFORMMATRIX2X4FV
	gpProgramUniformMatrix3dv                     C.GPPROGRAMUNIFORMMATRIX3DV
	gpProgramUniformMatrix3fv                     C.GPPROGRAMUNIFORMMATRIX3FV
	gpProgramUniformMatrix3x2dv                   C.GPPROGRAMUNIFORMMATRIX3X2DV
	gpProgramUniformMatrix3x2fv                   C.GPPROGRAMUNIFORMMATRIX3X2FV
	gpProgramUniformMatrix3x4dv                   C.GPPROGRAMUNIFORMMATRIX3X4DV
	gpProgramUniformMatrix3x4fv                   C.GPPROGRAMUNIFORMMATRIX3X4FV
	gpProgramUniformMatrix4dv                     C.GPPROGRAMUNIFORMMATRIX4DV
	gpProgramUniformMatrix4fv                     C.GPPROGRAMUNIFORMMATRIX4FV
	gpProgramUniformMatrix4x2dv                   C.GPPROGRAMUNIFORMMATRIX4X2DV
	gpProgramUniformMatrix4x2fv                   C.GPPROGRAMUNIFORMMATRIX4X2FV
	gpProgramUniformMatrix4x3dv                   C.GPPROGRAMUNIFORMMATRIX4X3DV
	gpProgramUniformMatrix4x3fv                   C.GPPROGRAMUNIFORMMATRIX4X3FV
	gpProvokingVertex                             C.GPPROVOKINGVERTEX
	gpPushDebugGroup                              C.GPPUSHDEBUGGROUP
	gpPushDebugGroupKHR                           C.GPPUSHDEBUGGROUPKHR
	gpQueryCounter                                C.GPQUERYCOUNTER
	gpReadBuffer                                  C.GPREADBUFFER
	gpReadPixels                                  C.GPREADPIXELS
	gpReadnPixels                                 C.GPREADNPIXELS
	gpReadnPixelsARB                              C.GPREADNPIXELSARB
	gpReadnPixelsKHR                              C.GPREADNPIXELSKHR
	gpReleaseShaderCompiler                       C.GPRELEASESHADERCOMPILER
	gpRenderbufferStorage                         C.GPRENDERBUFFERSTORAGE
	gpRenderbufferStorageMultisample              C.GPRENDERBUFFERSTORAGEMULTISAMPLE
	gpResumeTransformFeedback                     C.GPRESUMETRANSFORMFEEDBACK
	gpSampleCoverage                              C.GPSAMPLECOVERAGE
	gpSampleMaski                                 C.GPSAMPLEMASKI
	gpSamplerParameterIiv                         C.GPSAMPLERPARAMETERIIV
	gpSamplerParameterIuiv                        C.GPSAMPLERPARAMETERIUIV
	gpSamplerParameterf                           C.GPSAMPLERPARAMETERF
	gpSamplerParameterfv                          C.GPSAMPLERPARAMETERFV
	gpSamplerParameteri                           C.GPSAMPLERPARAMETERI
	gpSamplerParameteriv                          C.GPSAMPLERPARAMETERIV
	gpScissor                                     C.GPSCISSOR
	gpScissorArrayv                               C.GPSCISSORARRAYV
	gpScissorIndexed                              C.GPSCISSORINDEXED
	gpScissorIndexedv                             C.GPSCISSORINDEXEDV
	gpShaderBinary                                C.GPSHADERBINARY
	gpShaderSource                                C.GPSHADERSOURCE
	gpShaderStorageBlockBinding                   C.GPSHADERSTORAGEBLOCKBINDING
	gpStencilFunc                                 C.GPSTENCILFUNC
	gpStencilFuncSeparate                         C.GPSTENCILFUNCSEPARATE
	gpStencilMask                                 C.GPSTENCILMASK
	gpStencilMaskSeparate                         C.GPSTENCILMASKSEPARATE
	gpStencilOp                                   C.GPSTENCILOP
	gpStencilOpSeparate                           C.GPSTENCILOPSEPARATE
	gpTexBuffer                                   C.GPTEXBUFFER
	gpTexBufferRange                              C.GPTEXBUFFERRANGE
	gpTexImage1D                                  C.GPTEXIMAGE1D
	gpTexImage2D                                  C.GPTEXIMAGE2D
	gpTexImage2DMultisample                       C.GPTEXIMAGE2DMULTISAMPLE
	gpTexImage3D                                  C.GPTEXIMAGE3D
	gpTexImage3DMultisample                       C.GPTEXIMAGE3DMULTISAMPLE
	gpTexPageCommitmentARB                        C.GPTEXPAGECOMMITMENTARB
	gpTexParameterIiv                             C.GPTEXPARAMETERIIV
	gpTexParameterIuiv                            C.GPTEXPARAMETERIUIV
	gpTexParameterf                               C.GPTEXPARAMETERF
	gpTexParameterfv                              C.GPTEXPARAMETERFV
	gpTexParameteri                               C.GPTEXPARAMETERI
	gpTexParameteriv                              C.GPTEXPARAMETERIV
	gpTexStorage1D                                C.GPTEXSTORAGE1D
	gpTexStorage2D                                C.GPTEXSTORAGE2D
	gpTexStorage2DMultisample                     C.GPTEXSTORAGE2DMULTISAMPLE
	gpTexStorage3D                                C.GPTEXSTORAGE3D
	gpTexStorage3DMultisample                     C.GPTEXSTORAGE3DMULTISAMPLE
	gpTexSubImage1D                               C.GPTEXSUBIMAGE1D
	gpTexSubImage2D                               C.GPTEXSUBIMAGE2D
	gpTexSubImage3D                               C.GPTEXSUBIMAGE3D
	gpTextureBarrier                              C.GPTEXTUREBARRIER
	gpTextureBuffer                               C.GPTEXTUREBUFFER
	gpTextureBufferRange                          C.GPTEXTUREBUFFERRANGE
	gpTextureParameterIiv                         C.GPTEXTUREPARAMETERIIV
	gpTextureParameterIuiv                        C.GPTEXTUREPARAMETERIUIV
	gpTextureParameterf                           C.GPTEXTUREPARAMETERF
	gpTextureParameterfv                          C.GPTEXTUREPARAMETERFV
	gpTextureParameteri                           C.GPTEXTUREPARAMETERI
	gpTextureParameteriv                          C.GPTEXTUREPARAMETERIV
	gpTextureStorage1D                            C.GPTEXTURESTORAGE1D
	gpTextureStorage2D                            C.GPTEXTURESTORAGE2D
	gpTextureStorage2DMultisample                 C.GPTEXTURESTORAGE2DMULTISAMPLE
	gpTextureStorage3D                            C.GPTEXTURESTORAGE3D
	gpTextureStorage3DMultisample                 C.GPTEXTURESTORAGE3DMULTISAMPLE
	gpTextureSubImage1D                           C.GPTEXTURESUBIMAGE1D
	gpTextureSubImage2D                           C.GPTEXTURESUBIMAGE2D
	gpTextureSubImage3D                           C.GPTEXTURESUBIMAGE3D
	gpTextureView                                 C.GPTEXTUREVIEW
	gpTransformFeedbackBufferBase                 C.GPTRANSFORMFEEDBACKBUFFERBASE
	gpTransformFeedbackBufferRange                C.GPTRANSFORMFEEDBACKBUFFERRANGE
	gpTransformFeedbackVaryings                   C.GPTRANSFORMFEEDBACKVARYINGS
	gpUniform1d                                   C.GPUNIFORM1D
	gpUniform1dv                                  C.GPUNIFORM1DV
	gpUniform1f                                   C.GPUNIFORM1F
	gpUniform1fv                                  C.GPUNIFORM1FV
	gpUniform1i                                   C.GPUNIFORM1I
	gpUniform1iv                                  C.GPUNIFORM1IV
	gpUniform1ui                                  C.GPUNIFORM1UI
	gpUniform1uiv                                 C.GPUNIFORM1UIV
	gpUniform2d                                   C.GPUNIFORM2D
	gpUniform2dv                                  C.GPUNIFORM2DV
	gpUniform2f                                   C.GPUNIFORM2F
	gpUniform2fv                                  C.GPUNIFORM2FV
	gpUniform2i                                   C.GPUNIFORM2I
	gpUniform2iv                                  C.GPUNIFORM2IV
	gpUniform2ui                                  C.GPUNIFORM2UI
	gpUniform2uiv                                 C.GPUNIFORM2UIV
	gpUniform3d                                   C.GPUNIFORM3D
	gpUniform3dv                                  C.GPUNIFORM3DV
	gpUniform3f                                   C.GPUNIFORM3F
	gpUniform3fv                                  C.GPUNIFORM3FV
	gpUniform3i                                   C.GPUNIFORM3I
	gpUniform3iv                                  C.GPUNIFORM3IV
	gpUniform3ui                                  C.GPUNIFORM3UI
	gpUniform3uiv                                 C.GPUNIFORM3UIV
	gpUniform4d                                   C.GPUNIFORM4D
	gpUniform4dv                                  C.GPUNIFORM4DV
	gpUniform4f                                   C.GPUNIFORM4F
	gpUniform4fv                                  C.GPUNIFORM4FV
	gpUniform4i                                   C.GPUNIFORM4I
	gpUniform4iv                                  C.GPUNIFORM4IV
	gpUniform4ui                                  C.GPUNIFORM4UI
	gpUniform4uiv                                 C.GPUNIFORM4UIV
	gpUniformBlockBinding                         C.GPUNIFORMBLOCKBINDING
	gpUniformHandleui64ARB                        C.GPUNIFORMHANDLEUI64ARB
	gpUniformHandleui64vARB                       C.GPUNIFORMHANDLEUI64VARB
	gpUniformMatrix2dv                            C.GPUNIFORMMATRIX2DV
	gpUniformMatrix2fv                            C.GPUNIFORMMATRIX2FV
	gpUniformMatrix2x3dv                          C.GPUNIFORMMATRIX2X3DV
	gpUniformMatrix2x3fv                          C.GPUNIFORMMATRIX2X3FV
	gpUniformMatrix2x4dv                          C.GPUNIFORMMATRIX2X4DV
	gpUniformMatrix2x4fv                          C.GPUNIFORMMATRIX2X4FV
	gpUniformMatrix3dv                            C.GPUNIFORMMATRIX3DV
	gpUniformMatrix3fv                            C.GPUNIFORMMATRIX3FV
	gpUniformMatrix3x2dv                          C.GPUNIFORMMATRIX3X2DV
	gpUniformMatrix3x2fv                          C.GPUNIFORMMATRIX3X2FV
	gpUniformMatrix3x4dv                          C.GPUNIFORMMATRIX3X4DV
	gpUniformMatrix3x4fv                          C.GPUNIFORMMATRIX3X4FV
	gpUniformMatrix4dv                            C.GPUNIFORMMATRIX4DV
	gpUniformMatrix4fv                            C.GPUNIFORMMATRIX4FV
	gpUniformMatrix4x2dv                          C.GPUNIFORMMATRIX4X2DV
	gpUniformMatrix4x2fv                          C.GPUNIFORMMATRIX4X2FV
	gpUniformMatrix4x3dv                          C.GPUNIFORMMATRIX4X3DV
	gpUniformMatrix4x3fv                          C.GPUNIFORMMATRIX4X3FV
	gpUniformSubroutinesuiv                       C.GPUNIFORMSUBROUTINESUIV
	gpUnmapBuffer                                 C.GPUNMAPBUFFER
	gpUnmapNamedBuffer                            C.GPUNMAPNAMEDBUFFER
	gpUseProgram                                  C.GPUSEPROGRAM
	gpUseProgramStages                            C.GPUSEPROGRAMSTAGES
	gpValidateProgram                             C.GPVALIDATEPROGRAM
	gpValidateProgramPipeline                     C.GPVALIDATEPROGRAMPIPELINE
	gpVertexArrayAttribBinding                    C.GPVERTEXARRAYATTRIBBINDING
	gpVertexArrayAttribFormat                     C.GPVERTEXARRAYATTRIBFORMAT
	gpVertexArrayAttribIFormat                    C.GPVERTEXARRAYATTRIBIFORMAT
	gpVertexArrayAttribLFormat                    C.GPVERTEXARRAYATTRIBLFORMAT
	gpVertexArrayBindingDivisor                   C.GPVERTEXARRAYBINDINGDIVISOR
	gpVertexArrayElementBuffer                    C.GPVERTEXARRAYELEMENTBUFFER
	gpVertexArrayVertexBuffer                     C.GPVERTEXARRAYVERTEXBUFFER
	gpVertexArrayVertexBuffers                    C.GPVERTEXARRAYVERTEXBUFFERS
	gpVertexAttrib1d                              C.GPVERTEXATTRIB1D
	gpVertexAttrib1dv                             C.GPVERTEXATTRIB1DV
	gpVertexAttrib1f                              C.GPVERTEXATTRIB1F
	gpVertexAttrib1fv                             C.GPVERTEXATTRIB1FV
	gpVertexAttrib1s                              C.GPVERTEXATTRIB1S
	gpVertexAttrib1sv                             C.GPVERTEXATTRIB1SV
	gpVertexAttrib2d                              C.GPVERTEXATTRIB2D
	gpVertexAttrib2dv                             C.GPVERTEXATTRIB2DV
	gpVertexAttrib2f                              C.GPVERTEXATTRIB2F
	gpVertexAttrib2fv                             C.GPVERTEXATTRIB2FV
	gpVertexAttrib2s                              C.GPVERTEXATTRIB2S
	gpVertexAttrib2sv                             C.GPVERTEXATTRIB2SV
	gpVertexAttrib3d                              C.GPVERTEXATTRIB3D
	gpVertexAttrib3dv                             C.GPVERTEXATTRIB3DV
	gpVertexAttrib3f                              C.GPVERTEXATTRIB3F
	gpVertexAttrib3fv                             C.GPVERTEXATTRIB3FV
	gpVertexAttrib3s                              C.GPVERTEXATTRIB3S
	gpVertexAttrib3sv                             C.GPVERTEXATTRIB3SV
	gpVertexAttrib4Nbv                            C.GPVERTEXATTRIB4NBV
	gpVertexAttrib4Niv                            C.GPVERTEXATTRIB4NIV
	gpVertexAttrib4Nsv                            C.GPVERTEXATTRIB4NSV
	gpVertexAttrib4Nub                            C.GPVERTEXATTRIB4NUB
	gpVertexAttrib4Nubv                           C.GPVERTEXATTRIB4NUBV
	gpVertexAttrib4Nuiv                           C.GPVERTEXATTRIB4NUIV
	gpVertexAttrib4Nusv                           C.GPVERTEXATTRIB4NUSV
	gpVertexAttrib4bv                             C.GPVERTEXATTRIB4BV
	gpVertexAttrib4d                              C.GPVERTEXATTRIB4D
	gpVertexAttrib4dv                             C.GPVERTEXATTRIB4DV
	gpVertexAttrib4f                              C.GPVERTEXATTRIB4F
	gpVertexAttrib4fv                             C.GPVERTEXATTRIB4FV
	gpVertexAttrib4iv                             C.GPVERTEXATTRIB4IV
	gpVertexAttrib4s                              C.GPVERTEXATTRIB4S
	gpVertexAttrib4sv                             C.GPVERTEXATTRIB4SV
	gpVertexAttrib4ubv                            C.GPVERTEXATTRIB4UBV
	gpVertexAttrib4uiv                            C.GPVERTEXATTRIB4UIV
	gpVertexAttrib4usv                            C.GPVERTEXATTRIB4USV
	gpVertexAttribBinding                         C.GPVERTEXATTRIBBINDING
	gpVertexAttribDivisor                         C.GPVERTEXATTRIBDIVISOR
	gpVertexAttribFormat                          C.GPVERTEXATTRIBFORMAT
	gpVertexAttribI1i                             C.GPVERTEXATTRIBI1I
	gpVertexAttribI1iv                            C.GPVERTEXATTRIBI1IV
	gpVertexAttribI1ui                            C.GPVERTEXATTRIBI1UI
	gpVertexAttribI1uiv                           C.GPVERTEXATTRIBI1UIV
	gpVertexAttribI2i                             C.GPVERTEXATTRIBI2I
	gpVertexAttribI2iv                            C.GPVERTEXATTRIBI2IV
	gpVertexAttribI2ui                            C.GPVERTEXATTRIBI2UI
	gpVertexAttribI2uiv                           C.GPVERTEXATTRIBI2UIV
	gpVertexAttribI3i                             C.GPVERTEXATTRIBI3I
	gpVertexAttribI3iv                            C.GPVERTEXATTRIBI3IV
	gpVertexAttribI3ui                            C.GPVERTEXATTRIBI3UI
	gpVertexAttribI3uiv                           C.GPVERTEXATTRIBI3UIV
	gpVertexAttribI4bv                            C.GPVERTEXATTRIBI4BV
	gpVertexAttribI4i                             C.GPVERTEXATTRIBI4I
	gpVertexAttribI4iv                            C.GPVERTEXATTRIBI4IV
	gpVertexAttribI4sv                            C.GPVERTEXATTRIBI4SV
	gpVertexAttribI4ubv                           C.GPVERTEXATTRIBI4UBV
	gpVertexAttribI4ui                            C.GPVERTEXATTRIBI4UI
	gpVertexAttribI4uiv                           C.GPVERTEXATTRIBI4UIV
	gpVertexAttribI4usv                           C.GPVERTEXATTRIBI4USV
	gpVertexAttribIFormat                         C.GPVERTEXATTRIBIFORMAT
	gpVertexAttribIPointer                        C.GPVERTEXATTRIBIPOINTER
	gpVertexAttribL1d                             C.GPVERTEXATTRIBL1D
	gpVertexAttribL1dv                            C.GPVERTEXATTRIBL1DV
	gpVertexAttribL1ui64ARB                       C.GPVERTEXATTRIBL1UI64ARB
	gpVertexAttribL1ui64vARB                      C.GPVERTEXATTRIBL1UI64VARB
	gpVertexAttribL2d                             C.GPVERTEXATTRIBL2D
	gpVertexAttribL2dv                            C.GPVERTEXATTRIBL2DV
	gpVertexAttribL3d                             C.GPVERTEXATTRIBL3D
	gpVertexAttribL3dv                            C.GPVERTEXATTRIBL3DV
	gpVertexAttribL4d                             C.GPVERTEXATTRIBL4D
	gpVertexAttribL4dv                            C.GPVERTEXATTRIBL4DV
	gpVertexAttribLFormat                         C.GPVERTEXATTRIBLFORMAT
	gpVertexAttribLPointer                        C.GPVERTEXATTRIBLPOINTER
	gpVertexAttribP1ui                            C.GPVERTEXATTRIBP1UI
	gpVertexAttribP1uiv                           C.GPVERTEXATTRIBP1UIV
	gpVertexAttribP2ui                            C.GPVERTEXATTRIBP2UI
	gpVertexAttribP2uiv                           C.GPVERTEXATTRIBP2UIV
	gpVertexAttribP3ui                            C.GPVERTEXATTRIBP3UI
	gpVertexAttribP3uiv                           C.GPVERTEXATTRIBP3UIV
	gpVertexAttribP4ui                            C.GPVERTEXATTRIBP4UI
	gpVertexAttribP4uiv                           C.GPVERTEXATTRIBP4UIV
	gpVertexAttribPointer                         C.GPVERTEXATTRIBPOINTER
	gpVertexBindingDivisor                        C.GPVERTEXBINDINGDIVISOR
	gpViewport                                    C.GPVIEWPORT
	gpViewportArrayv                              C.GPVIEWPORTARRAYV
	gpViewportIndexedf                            C.GPVIEWPORTINDEXEDF
	gpViewportIndexedfv                           C.GPVIEWPORTINDEXEDFV
	gpWaitSync                                    C.GPWAITSYNC
)

func ActiveShaderProgram(pipeline uint32, program uint32) {
	C.ptrActiveShaderProgram(gpActiveShaderProgram, (C.GLuint)(pipeline), (C.GLuint)(program))
}
func ActiveTexture(texture int32) {
	C.ptrActiveTexture(gpActiveTexture, (C.GLenum)(texture))
}
func AttachShader(program uint32, shader uint32) {
	C.ptrAttachShader(gpAttachShader, (C.GLuint)(program), (C.GLuint)(shader))
}
func BeginConditionalRender(id uint32, mode int32) {
	C.ptrBeginConditionalRender(gpBeginConditionalRender, (C.GLuint)(id), (C.GLenum)(mode))
}
func BeginQuery(target int32, id uint32) {
	C.ptrBeginQuery(gpBeginQuery, (C.GLenum)(target), (C.GLuint)(id))
}
func BeginQueryIndexed(target int32, index uint32, id uint32) {
	C.ptrBeginQueryIndexed(gpBeginQueryIndexed, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(id))
}
func BeginTransformFeedback(primitiveMode int32) {
	C.ptrBeginTransformFeedback(gpBeginTransformFeedback, (C.GLenum)(primitiveMode))
}
func BindAttribLocation(program uint32, index uint32, name string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrBindAttribLocation(gpBindAttribLocation, (C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func BindBuffer(target int32, buffer uint32) {
	C.ptrBindBuffer(gpBindBuffer, (C.GLenum)(target), (C.GLuint)(buffer))
}
func BindBufferBase(target int32, index uint32, buffer uint32) {
	C.ptrBindBufferBase(gpBindBufferBase, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}
func BindBufferRange(target int32, index uint32, buffer uint32, offset int64, size int64) {
	C.ptrBindBufferRange(gpBindBufferRange, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
func BindBuffersBase(target int32, first uint32, count int32, buffers *uint32) {
	C.ptrBindBuffersBase(gpBindBuffersBase, (C.GLenum)(target), (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func BindBuffersRange(target int32, first uint32, count int32, buffers *uint32, offsets *int64, sizes *int64) {
	C.ptrBindBuffersRange(gpBindBuffersRange, (C.GLenum)(target), (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)), (*C.GLintptr)(unsafe.Pointer(offsets)), (*C.GLsizeiptr)(unsafe.Pointer(sizes)))
}
func BindFragDataLocation(program uint32, color uint32, name string) {
	code := name + "\x00"
	cstrname := stringToUint8(code)
	C.ptrBindFragDataLocation(gpBindFragDataLocation, (C.GLuint)(program), (C.GLuint)(color), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func BindFragDataLocationIndexed(program uint32, colorNumber uint32, index uint32, name string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrBindFragDataLocationIndexed(gpBindFragDataLocationIndexed, (C.GLuint)(program), (C.GLuint)(colorNumber), (C.GLuint)(index), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func BindFramebuffer(target int32, framebuffer uint32) {
	C.ptrBindFramebuffer(gpBindFramebuffer, (C.GLenum)(target), (C.GLuint)(framebuffer))
}
func BindImageTexture(unit uint32, texture uint32, level int32, layered bool, layer int32, access int32, format int32) {
	C.ptrBindImageTexture(gpBindImageTexture, (C.GLuint)(unit), (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(boolToInt(layered)), (C.GLint)(layer), (C.GLenum)(access), (C.GLenum)(format))
}
func BindImageTextures(first uint32, count int32, textures *uint32) {
	C.ptrBindImageTextures(gpBindImageTextures, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(textures)))
}
func BindProgramPipeline(pipeline uint32) {
	C.ptrBindProgramPipeline(gpBindProgramPipeline, (C.GLuint)(pipeline))
}
func BindRenderbuffer(target int32, renderbuffer uint32) {
	C.ptrBindRenderbuffer(gpBindRenderbuffer, (C.GLenum)(target), (C.GLuint)(renderbuffer))
}
func BindSampler(unit uint32, sampler uint32) {
	C.ptrBindSampler(gpBindSampler, (C.GLuint)(unit), (C.GLuint)(sampler))
}
func BindSamplers(first uint32, count int32, samplers *uint32) {
	C.ptrBindSamplers(gpBindSamplers, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}
func BindTexture(target int32, texture uint32) {
	C.ptrBindTexture(gpBindTexture, (C.GLenum)(target), (C.GLuint)(texture))
}
func BindTextureUnit(unit uint32, texture uint32) {
	C.ptrBindTextureUnit(gpBindTextureUnit, (C.GLuint)(unit), (C.GLuint)(texture))
}
func BindTextures(first uint32, count int32, textures *uint32) {
	C.ptrBindTextures(gpBindTextures, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(textures)))
}
func BindTransformFeedback(target int32, id uint32) {
	C.ptrBindTransformFeedback(gpBindTransformFeedback, (C.GLenum)(target), (C.GLuint)(id))
}
func BindVertexArray(array uint32) {
	C.ptrBindVertexArray(gpBindVertexArray, (C.GLuint)(array))
}
func BindVertexBuffer(bindingindex uint32, buffer uint32, offset int64, stride int32) {
	C.ptrBindVertexBuffer(gpBindVertexBuffer, (C.GLuint)(bindingindex), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(stride))
}
func BindVertexBuffers(first uint32, count int32, buffers *uint32, offsets *int64, strides *int32) {
	C.ptrBindVertexBuffers(gpBindVertexBuffers, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)), (*C.GLintptr)(unsafe.Pointer(offsets)), (*C.GLsizei)(unsafe.Pointer(strides)))
}
func BlendColor(red float32, green float32, blue float32, alpha float32) {
	C.ptrBlendColor(gpBlendColor, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
func BlendEquation(mode int32) {
	C.ptrBlendEquation(gpBlendEquation, (C.GLenum)(mode))
}
func BlendEquationSeparate(modeRGB int32, modeAlpha int32) {
	C.ptrBlendEquationSeparate(gpBlendEquationSeparate, (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationSeparatei(buf uint32, modeRGB int32, modeAlpha int32) {
	C.ptrBlendEquationSeparatei(gpBlendEquationSeparatei, (C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationSeparateiARB(buf uint32, modeRGB int32, modeAlpha int32) {
	C.ptrBlendEquationSeparateiARB(gpBlendEquationSeparateiARB, (C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationi(buf uint32, mode int32) {
	C.ptrBlendEquationi(gpBlendEquationi, (C.GLuint)(buf), (C.GLenum)(mode))
}
func BlendEquationiARB(buf uint32, mode int32) {
	C.ptrBlendEquationiARB(gpBlendEquationiARB, (C.GLuint)(buf), (C.GLenum)(mode))
}
func BlendFunc(sfactor int32, dfactor int32) {
	C.ptrBlendFunc(gpBlendFunc, (C.GLenum)(sfactor), (C.GLenum)(dfactor))
}
func BlendFuncSeparate(sfactorRGB int32, dfactorRGB int32, sfactorAlpha int32, dfactorAlpha int32) {
	C.ptrBlendFuncSeparate(gpBlendFuncSeparate, (C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
func BlendFuncSeparatei(buf uint32, srcRGB int32, dstRGB int32, srcAlpha int32, dstAlpha int32) {
	C.ptrBlendFuncSeparatei(gpBlendFuncSeparatei, (C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
func BlendFuncSeparateiARB(buf uint32, srcRGB int32, dstRGB int32, srcAlpha int32, dstAlpha int32) {
	C.ptrBlendFuncSeparateiARB(gpBlendFuncSeparateiARB, (C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
func BlendFunci(buf uint32, src int32, dst int32) {
	C.ptrBlendFunci(gpBlendFunci, (C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}
func BlendFunciARB(buf uint32, src int32, dst int32) {
	C.ptrBlendFunciARB(gpBlendFunciARB, (C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}
func BlitFramebuffer(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter int32) {
	C.ptrBlitFramebuffer(gpBlitFramebuffer, (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}
func BlitNamedFramebuffer(readFramebuffer uint32, drawFramebuffer uint32, srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter int32) {
	C.ptrBlitNamedFramebuffer(gpBlitNamedFramebuffer, (C.GLuint)(readFramebuffer), (C.GLuint)(drawFramebuffer), (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}
func BufferData(target int32, size int64, data interface{}, usage int32) {
	C.ptrBufferData(gpBufferData, (C.GLenum)(target), (C.GLsizeiptr)(size), ptr(data), (C.GLenum)(usage))
}
func BufferPageCommitmentARB(target int32, offset int64, size int32, commit bool) {
	C.ptrBufferPageCommitmentARB(gpBufferPageCommitmentARB, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLboolean)(boolToInt(commit)))
}
func BufferStorage(target int32, size int64, data interface{}, flags uint32) {
	C.ptrBufferStorage(gpBufferStorage, (C.GLenum)(target), (C.GLsizeiptr)(size), ptr(data), (C.GLbitfield)(flags))
}
func BufferSubData(target int32, offset int64, size int64, data interface{}) {
	C.ptrBufferSubData(gpBufferSubData, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), ptr(data))
}
func CheckFramebufferStatus(target int32) int32 {
	ret := C.ptrCheckFramebufferStatus(gpCheckFramebufferStatus, (C.GLenum)(target))
	return (int32)(ret)
}
func CheckNamedFramebufferStatus(framebuffer uint32, target int32) int32 {
	ret := C.ptrCheckNamedFramebufferStatus(gpCheckNamedFramebufferStatus, (C.GLuint)(framebuffer), (C.GLenum)(target))
	return (int32)(ret)
}
func ClampColor(target int32, clamp int32) {
	C.ptrClampColor(gpClampColor, (C.GLenum)(target), (C.GLenum)(clamp))
}
func Clear(mask uint32) {
	C.ptrClear(gpClear, (C.GLbitfield)(mask))
}
func ClearBufferData(target int32, internalformat int32, format int32, gltype int32, data interface{}) {
	C.ptrClearBufferData(gpClearBufferData, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(format), (C.GLenum)(gltype), ptr(data))
}
func ClearBufferSubData(target int32, internalformat int32, offset int64, size int64, format int32, gltype int32, data interface{}) {
	C.ptrClearBufferSubData(gpClearBufferSubData, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (C.GLenum)(format), (C.GLenum)(gltype), ptr(data))
}
func ClearBufferfi(buffer int32, drawbuffer int32, depth float32, stencil int32) {
	C.ptrClearBufferfi(gpClearBufferfi, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
func ClearBufferfv(buffer int32, drawbuffer int32, value *float32) {
	C.ptrClearBufferfv(gpClearBufferfv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ClearBufferiv(buffer int32, drawbuffer int32, value *int32) {
	C.ptrClearBufferiv(gpClearBufferiv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(unsafe.Pointer(value)))
}
func ClearBufferuiv(buffer int32, drawbuffer int32, value *uint32) {
	C.ptrClearBufferuiv(gpClearBufferuiv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(unsafe.Pointer(value)))
}
func ClearColor(red float32, green float32, blue float32, alpha float32) {
	C.ptrClearColor(gpClearColor, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
func ClearDepth(depth float64) {
	C.ptrClearDepth(gpClearDepth, (C.GLdouble)(depth))
}
func ClearDepthf(d float32) {
	C.ptrClearDepthf(gpClearDepthf, (C.GLfloat)(d))
}
func ClearNamedBufferData(buffer uint32, internalformat int32, format int32, gltype int32, data interface{}) {
	C.ptrClearNamedBufferData(gpClearNamedBufferData, (C.GLuint)(buffer), (C.GLenum)(internalformat), (C.GLenum)(format), (C.GLenum)(gltype), ptr(data))
}
func ClearNamedBufferSubData(buffer uint32, internalformat int32, offset int64, size int32, format int32, gltype int32, data interface{}) {
	C.ptrClearNamedBufferSubData(gpClearNamedBufferSubData, (C.GLuint)(buffer), (C.GLenum)(internalformat), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLenum)(format), (C.GLenum)(gltype), ptr(data))
}
func ClearNamedFramebufferfi(framebuffer uint32, buffer int32, depth float32, stencil int32) {
	C.ptrClearNamedFramebufferfi(gpClearNamedFramebufferfi, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
func ClearNamedFramebufferfv(framebuffer uint32, buffer int32, drawbuffer int32, value *float32) {
	C.ptrClearNamedFramebufferfv(gpClearNamedFramebufferfv, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ClearNamedFramebufferiv(framebuffer uint32, buffer int32, drawbuffer int32, value *int32) {
	C.ptrClearNamedFramebufferiv(gpClearNamedFramebufferiv, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(unsafe.Pointer(value)))
}
func ClearNamedFramebufferuiv(framebuffer uint32, buffer int32, drawbuffer int32, value *uint32) {
	C.ptrClearNamedFramebufferuiv(gpClearNamedFramebufferuiv, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(unsafe.Pointer(value)))
}
func ClearStencil(s int32) {
	C.ptrClearStencil(gpClearStencil, (C.GLint)(s))
}
func ClearTexImage(texture uint32, level int32, format int32, gltype int32, data interface{}) {
	C.ptrClearTexImage(gpClearTexImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(gltype), ptr(data))
}
func ClearTexSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format int32, gltype int32, data interface{}) {
	C.ptrClearTexSubImage(gpClearTexSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(gltype), ptr(data))
}
func ClientWaitSync(sync unsafe.Pointer, flags uint32, timeout uint64) int32 {
	ret := C.ptrClientWaitSync(gpClientWaitSync, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
	return (int32)(ret)
}
func ClipControl(origin int32, depth int32) {
	C.ptrClipControl(gpClipControl, (C.GLenum)(origin), (C.GLenum)(depth))
}
func ColorMask(red bool, green bool, blue bool, alpha bool) {
	C.ptrColorMask(gpColorMask, (C.GLboolean)(boolToInt(red)), (C.GLboolean)(boolToInt(green)), (C.GLboolean)(boolToInt(blue)), (C.GLboolean)(boolToInt(alpha)))
}
func ColorMaski(index uint32, r bool, g bool, b bool, a bool) {
	C.ptrColorMaski(gpColorMaski, (C.GLuint)(index), (C.GLboolean)(boolToInt(r)), (C.GLboolean)(boolToInt(g)), (C.GLboolean)(boolToInt(b)), (C.GLboolean)(boolToInt(a)))
}
func CompileShader(shader uint32) {
	C.ptrCompileShader(gpCompileShader, (C.GLuint)(shader))
}
func CompileShaderIncludeARB(shader uint32, count int32, path string, length *int32) {
	cstrpath, freepath := stringToChar(path)
	defer freepath()
	C.ptrCompileShaderIncludeARB(gpCompileShaderIncludeARB, (C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(cstrpath)), (*C.GLint)(unsafe.Pointer(length)))
}
func CompressedTexImage1D(target int32, level int32, internalformat int32, width int32, border int32, imageSize int32, data interface{}) {
	C.ptrCompressedTexImage1D(gpCompressedTexImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), ptr(data))
}
func CompressedTexImage2D(target int32, level int32, internalformat int32, width int32, height int32, border int32, imageSize int32, data interface{}) {
	C.ptrCompressedTexImage2D(gpCompressedTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), ptr(data))
}
func CompressedTexImage3D(target int32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, imageSize int32, data interface{}) {
	C.ptrCompressedTexImage3D(gpCompressedTexImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), ptr(data))
}
func CompressedTexSubImage1D(target int32, level int32, xoffset int32, width int32, format int32, imageSize int32, data interface{}) {
	C.ptrCompressedTexSubImage1D(gpCompressedTexSubImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), ptr(data))
}
func CompressedTexSubImage2D(target int32, level int32, xoffset int32, yoffset int32, width int32, height int32, format int32, imageSize int32, data interface{}) {
	C.ptrCompressedTexSubImage2D(gpCompressedTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), ptr(data))
}
func CompressedTexSubImage3D(target int32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format int32, imageSize int32, data interface{}) {
	C.ptrCompressedTexSubImage3D(gpCompressedTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), ptr(data))
}
func CompressedTextureSubImage1D(texture uint32, level int32, xoffset int32, width int32, format int32, imageSize int32, data interface{}) {
	C.ptrCompressedTextureSubImage1D(gpCompressedTextureSubImage1D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), ptr(data))
}
func CompressedTextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format int32, imageSize int32, data interface{}) {
	C.ptrCompressedTextureSubImage2D(gpCompressedTextureSubImage2D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), ptr(data))
}
func CompressedTextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format int32, imageSize int32, data interface{}) {
	C.ptrCompressedTextureSubImage3D(gpCompressedTextureSubImage3D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), ptr(data))
}
func CopyBufferSubData(readTarget int32, writeTarget int32, readOffset int64, writeOffset int64, size int64) {
	C.ptrCopyBufferSubData(gpCopyBufferSubData, (C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}
func CopyImageSubData(srcName uint32, srcTarget int32, srcLevel int32, srcX int32, srcY int32, srcZ int32, dstName uint32, dstTarget int32, dstLevel int32, dstX int32, dstY int32, dstZ int32, srcWidth int32, srcHeight int32, srcDepth int32) {
	C.ptrCopyImageSubData(gpCopyImageSubData, (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(srcWidth), (C.GLsizei)(srcHeight), (C.GLsizei)(srcDepth))
}
func CopyNamedBufferSubData(readBuffer uint32, writeBuffer uint32, readOffset int64, writeOffset int64, size int32) {
	C.ptrCopyNamedBufferSubData(gpCopyNamedBufferSubData, (C.GLuint)(readBuffer), (C.GLuint)(writeBuffer), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizei)(size))
}
func CopyTexImage1D(target int32, level int32, internalformat int32, x int32, y int32, width int32, border int32) {
	C.ptrCopyTexImage1D(gpCopyTexImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}
func CopyTexImage2D(target int32, level int32, internalformat int32, x int32, y int32, width int32, height int32, border int32) {
	C.ptrCopyTexImage2D(gpCopyTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}
func CopyTexSubImage1D(target int32, level int32, xoffset int32, x int32, y int32, width int32) {
	C.ptrCopyTexSubImage1D(gpCopyTexSubImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
func CopyTexSubImage2D(target int32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	C.ptrCopyTexSubImage2D(gpCopyTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func CopyTexSubImage3D(target int32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	C.ptrCopyTexSubImage3D(gpCopyTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func CopyTextureSubImage1D(texture uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	C.ptrCopyTextureSubImage1D(gpCopyTextureSubImage1D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}
func CopyTextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	C.ptrCopyTextureSubImage2D(gpCopyTextureSubImage2D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func CopyTextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	C.ptrCopyTextureSubImage3D(gpCopyTextureSubImage3D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func CreateBuffers(n int32, buffers *uint32) {
	C.ptrCreateBuffers(gpCreateBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func CreateFramebuffers(n int32, framebuffers *uint32) {
	C.ptrCreateFramebuffers(gpCreateFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}
func CreateProgram() uint32 {
	ret := C.ptrCreateProgram(gpCreateProgram)
	return (uint32)(ret)
}
func CreateProgramPipelines(n int32, pipelines *uint32) {
	C.ptrCreateProgramPipelines(gpCreateProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}
func CreateQueries(target int32, n int32, ids *uint32) {
	C.ptrCreateQueries(gpCreateQueries, (C.GLenum)(target), (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}
func CreateRenderbuffers(n int32, renderbuffers *uint32) {
	C.ptrCreateRenderbuffers(gpCreateRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}
func CreateSamplers(n int32, samplers *uint32) {
	C.ptrCreateSamplers(gpCreateSamplers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(samplers)))
}
func CreateShader(gltype int32) uint32 {
	ret := C.ptrCreateShader(gpCreateShader, (C.GLenum)(gltype))
	return (uint32)(ret)
}
func CreateShaderProgramv(gltype int32, count int32, strings string) uint32 {
	cstrstrings, freestrings := stringToChar(strings)
	defer freestrings()
	ret := C.ptrCreateShaderProgramv(gpCreateShaderProgramv, (C.GLenum)(gltype), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(cstrstrings)))
	return (uint32)(ret)
}
func CreateSyncFromCLeventARB(context unsafe.Pointer, event unsafe.Pointer, flags uint32) unsafe.Pointer {
	ret := C.ptrCreateSyncFromCLeventARB(gpCreateSyncFromCLeventARB, (*C.struct__cl_context)(context), (*C.struct__cl_event)(event), (C.GLbitfield)(flags))
	return (unsafe.Pointer)(ret)
}
func CreateTextures(target int32, n int32, textures *uint32) {
	C.ptrCreateTextures(gpCreateTextures, (C.GLenum)(target), (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}
func CreateTransformFeedbacks(n int32, ids *uint32) {
	C.ptrCreateTransformFeedbacks(gpCreateTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}
func CreateVertexArrays(n int32, arrays *uint32) {
	C.ptrCreateVertexArrays(gpCreateVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}
func CullFace(mode int32) {
	C.ptrCullFace(gpCullFace, (C.GLenum)(mode))
}
func DebugMessageCallback(callback DebugProc, userParam unsafe.Pointer) {
	C.ptrDebugMessageCallback(gpDebugMessageCallback, (C.GLDEBUGPROC)(unsafe.Pointer(&callback)), userParam)
}
func DebugMessageCallbackARB(callback DebugProc, userParam unsafe.Pointer) {
	C.ptrDebugMessageCallbackARB(gpDebugMessageCallbackARB, (C.GLDEBUGPROCARB)(unsafe.Pointer(&callback)), userParam)
}
func DebugMessageCallbackKHR(callback DebugProc, userParam unsafe.Pointer) {
	C.ptrDebugMessageCallbackKHR(gpDebugMessageCallbackKHR, (C.GLDEBUGPROCKHR)(unsafe.Pointer(&callback)), userParam)
}
func DebugMessageControl(source int32, gltype int32, severity int32, count int32, ids *uint32, enabled bool) {
	C.ptrDebugMessageControl(gpDebugMessageControl, (C.GLenum)(source), (C.GLenum)(gltype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}
func DebugMessageControlARB(source int32, gltype int32, severity int32, count int32, ids *uint32, enabled bool) {
	C.ptrDebugMessageControlARB(gpDebugMessageControlARB, (C.GLenum)(source), (C.GLenum)(gltype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}
func DebugMessageControlKHR(source int32, gltype int32, severity int32, count int32, ids *uint32, enabled bool) {
	C.ptrDebugMessageControlKHR(gpDebugMessageControlKHR, (C.GLenum)(source), (C.GLenum)(gltype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}
func DebugMessageInsert(source int32, gltype int32, id uint32, severity int32, length int32, buf string) {
	cstrbuf, freebuf := stringToChar(buf)
	defer freebuf()
	C.ptrDebugMessageInsert(gpDebugMessageInsert, (C.GLenum)(source), (C.GLenum)(gltype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(cstrbuf)))
}
func DebugMessageInsertARB(source int32, gltype int32, id uint32, severity int32, length int32, buf string) {
	cstrbuf, freebuf := stringToChar(buf)
	defer freebuf()
	C.ptrDebugMessageInsertARB(gpDebugMessageInsertARB, (C.GLenum)(source), (C.GLenum)(gltype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(cstrbuf)))
}
func DebugMessageInsertKHR(source int32, gltype int32, id uint32, severity int32, length int32, buf string) {
	cstrbuf, freebuf := stringToChar(buf)
	defer freebuf()
	C.ptrDebugMessageInsertKHR(gpDebugMessageInsertKHR, (C.GLenum)(source), (C.GLenum)(gltype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(cstrbuf)))
}
func DeleteBuffers(n int32, buffers *uint32) {
	C.ptrDeleteBuffers(gpDeleteBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func DeleteFramebuffers(n int32, framebuffers *uint32) {
	C.ptrDeleteFramebuffers(gpDeleteFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}
func DeleteNamedStringARB(namelen int32, name string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrDeleteNamedStringARB(gpDeleteNamedStringARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func DeleteProgram(program uint32) {
	C.ptrDeleteProgram(gpDeleteProgram, (C.GLuint)(program))
}
func DeleteProgramPipelines(n int32, pipelines *uint32) {
	C.ptrDeleteProgramPipelines(gpDeleteProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}
func DeleteQueries(n int32, ids *uint32) {
	C.ptrDeleteQueries(gpDeleteQueries, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}
func DeleteRenderbuffers(n int32, renderbuffers *uint32) {
	C.ptrDeleteRenderbuffers(gpDeleteRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}
func DeleteSamplers(count int32, samplers *uint32) {
	C.ptrDeleteSamplers(gpDeleteSamplers, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}
func DeleteShader(shader uint32) {
	C.ptrDeleteShader(gpDeleteShader, (C.GLuint)(shader))
}
func DeleteSync(sync unsafe.Pointer) {
	C.ptrDeleteSync(gpDeleteSync, (C.GLsync)(sync))
}
func DeleteTextures(n int32, textures *uint32) {
	C.ptrDeleteTextures(gpDeleteTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}
func DeleteTransformFeedbacks(n int32, ids *uint32) {
	C.ptrDeleteTransformFeedbacks(gpDeleteTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}
func DeleteVertexArrays(n int32, arrays *uint32) {
	C.ptrDeleteVertexArrays(gpDeleteVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}
func DepthFunc(glfunc int32) {
	C.ptrDepthFunc(gpDepthFunc, (C.GLenum)(glfunc))
}
func DepthMask(flag bool) {
	C.ptrDepthMask(gpDepthMask, (C.GLboolean)(boolToInt(flag)))
}
func DepthRange(xnear float64, xfar float64) {
	C.ptrDepthRange(gpDepthRange, (C.GLdouble)(xnear), (C.GLdouble)(xfar))
}
func DepthRangeArrayv(first uint32, count int32, v *float64) {
	C.ptrDepthRangeArrayv(gpDepthRangeArrayv, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(v)))
}
func DepthRangeIndexed(index uint32, n float64, f float64) {
	C.ptrDepthRangeIndexed(gpDepthRangeIndexed, (C.GLuint)(index), (C.GLdouble)(n), (C.GLdouble)(f))
}
func DepthRangef(n float32, f float32) {
	C.ptrDepthRangef(gpDepthRangef, (C.GLfloat)(n), (C.GLfloat)(f))
}
func DetachShader(program uint32, shader uint32) {
	C.ptrDetachShader(gpDetachShader, (C.GLuint)(program), (C.GLuint)(shader))
}
func Disable(cap int32) {
	C.ptrDisable(gpDisable, (C.GLenum)(cap))
}
func DisableVertexArrayAttrib(vaobj uint32, index uint32) {
	C.ptrDisableVertexArrayAttrib(gpDisableVertexArrayAttrib, (C.GLuint)(vaobj), (C.GLuint)(index))
}
func DisableVertexAttribArray(index uint32) {
	C.ptrDisableVertexAttribArray(gpDisableVertexAttribArray, (C.GLuint)(index))
}
func Disablei(target int32, index uint32) {
	C.ptrDisablei(gpDisablei, (C.GLenum)(target), (C.GLuint)(index))
}
func DispatchCompute(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32) {
	C.ptrDispatchCompute(gpDispatchCompute, (C.GLuint)(num_groups_x), (C.GLuint)(num_groups_y), (C.GLuint)(num_groups_z))
}
func DispatchComputeGroupSizeARB(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32, group_size_x uint32, group_size_y uint32, group_size_z uint32) {
	C.ptrDispatchComputeGroupSizeARB(gpDispatchComputeGroupSizeARB, (C.GLuint)(num_groups_x), (C.GLuint)(num_groups_y), (C.GLuint)(num_groups_z), (C.GLuint)(group_size_x), (C.GLuint)(group_size_y), (C.GLuint)(group_size_z))
}
func DispatchComputeIndirect(indirect int64) {
	C.ptrDispatchComputeIndirect(gpDispatchComputeIndirect, (C.GLintptr)(indirect))
}
func DrawArrays(mode int32, first int32, count int32) {
	C.ptrDrawArrays(gpDrawArrays, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}
func DrawArraysIndirect(mode int32, indirect unsafe.Pointer) {
	C.ptrDrawArraysIndirect(gpDrawArraysIndirect, (C.GLenum)(mode), indirect)
}
func DrawArraysInstanced(mode int32, first int32, count int32, instancecount int32) {
	C.ptrDrawArraysInstanced(gpDrawArraysInstanced, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount))
}
func DrawArraysInstancedBaseInstance(mode int32, first int32, count int32, instancecount int32, baseinstance uint32) {
	C.ptrDrawArraysInstancedBaseInstance(gpDrawArraysInstancedBaseInstance, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}
func DrawBuffer(buf int32) {
	C.ptrDrawBuffer(gpDrawBuffer, (C.GLenum)(buf))
}
func DrawBuffers(n int32, bufs *int32) {
	C.ptrDrawBuffers(gpDrawBuffers, (C.GLsizei)(n), (*C.GLenum)(unsafe.Pointer(bufs)))
}
func DrawElements(mode int32, count int32, gltype int32, indices unsafe.Pointer) {
	C.ptrDrawElements(gpDrawElements, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(gltype), indices)
}
func DrawElementsBaseVertex(mode int32, count int32, gltype int32, indices unsafe.Pointer, basevertex int32) {
	C.ptrDrawElementsBaseVertex(gpDrawElementsBaseVertex, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(gltype), indices, (C.GLint)(basevertex))
}
func DrawElementsIndirect(mode int32, gltype int32, indirect unsafe.Pointer) {
	C.ptrDrawElementsIndirect(gpDrawElementsIndirect, (C.GLenum)(mode), (C.GLenum)(gltype), indirect)
}
func DrawElementsInstanced(mode int32, count int32, gltype int32, indices unsafe.Pointer, instancecount int32) {
	C.ptrDrawElementsInstanced(gpDrawElementsInstanced, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(gltype), indices, (C.GLsizei)(instancecount))
}
func DrawElementsInstancedBaseInstance(mode int32, count int32, gltype int32, indices unsafe.Pointer, instancecount int32, baseinstance uint32) {
	C.ptrDrawElementsInstancedBaseInstance(gpDrawElementsInstancedBaseInstance, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(gltype), indices, (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}
func DrawElementsInstancedBaseVertex(mode int32, count int32, gltype int32, indices unsafe.Pointer, instancecount int32, basevertex int32) {
	C.ptrDrawElementsInstancedBaseVertex(gpDrawElementsInstancedBaseVertex, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(gltype), indices, (C.GLsizei)(instancecount), (C.GLint)(basevertex))
}
func DrawElementsInstancedBaseVertexBaseInstance(mode int32, count int32, gltype int32, indices unsafe.Pointer, instancecount int32, basevertex int32, baseinstance uint32) {
	C.ptrDrawElementsInstancedBaseVertexBaseInstance(gpDrawElementsInstancedBaseVertexBaseInstance, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(gltype), indices, (C.GLsizei)(instancecount), (C.GLint)(basevertex), (C.GLuint)(baseinstance))
}
func DrawRangeElements(mode int32, start uint32, end uint32, count int32, gltype int32, indices unsafe.Pointer) {
	C.ptrDrawRangeElements(gpDrawRangeElements, (C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(gltype), indices)
}
func DrawRangeElementsBaseVertex(mode int32, start uint32, end uint32, count int32, gltype int32, indices unsafe.Pointer, basevertex int32) {
	C.ptrDrawRangeElementsBaseVertex(gpDrawRangeElementsBaseVertex, (C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(gltype), indices, (C.GLint)(basevertex))
}
func DrawTransformFeedback(mode int32, id uint32) {
	C.ptrDrawTransformFeedback(gpDrawTransformFeedback, (C.GLenum)(mode), (C.GLuint)(id))
}
func DrawTransformFeedbackInstanced(mode int32, id uint32, instancecount int32) {
	C.ptrDrawTransformFeedbackInstanced(gpDrawTransformFeedbackInstanced, (C.GLenum)(mode), (C.GLuint)(id), (C.GLsizei)(instancecount))
}
func DrawTransformFeedbackStream(mode int32, id uint32, stream uint32) {
	C.ptrDrawTransformFeedbackStream(gpDrawTransformFeedbackStream, (C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream))
}
func DrawTransformFeedbackStreamInstanced(mode int32, id uint32, stream uint32, instancecount int32) {
	C.ptrDrawTransformFeedbackStreamInstanced(gpDrawTransformFeedbackStreamInstanced, (C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream), (C.GLsizei)(instancecount))
}
func Enable(cap int32) {
	C.ptrEnable(gpEnable, (C.GLenum)(cap))
}
func EnableVertexArrayAttrib(vaobj uint32, index uint32) {
	C.ptrEnableVertexArrayAttrib(gpEnableVertexArrayAttrib, (C.GLuint)(vaobj), (C.GLuint)(index))
}
func EnableVertexAttribArray(index uint32) {
	C.ptrEnableVertexAttribArray(gpEnableVertexAttribArray, (C.GLuint)(index))
}
func Enablei(target int32, index uint32) {
	C.ptrEnablei(gpEnablei, (C.GLenum)(target), (C.GLuint)(index))
}
func EndConditionalRender() {
	C.ptrEndConditionalRender(gpEndConditionalRender)
}
func EndQuery(target int32) {
	C.ptrEndQuery(gpEndQuery, (C.GLenum)(target))
}
func EndQueryIndexed(target int32, index uint32) {
	C.ptrEndQueryIndexed(gpEndQueryIndexed, (C.GLenum)(target), (C.GLuint)(index))
}
func EndTransformFeedback() {
	C.ptrEndTransformFeedback(gpEndTransformFeedback)
}
func FenceSync(condition int32, flags uint32) unsafe.Pointer {
	ret := C.ptrFenceSync(gpFenceSync, (C.GLenum)(condition), (C.GLbitfield)(flags))
	return (unsafe.Pointer)(ret)
}
func Finish() {
	C.ptrFinish(gpFinish)
}
func Flush() {
	C.ptrFlush(gpFlush)
}
func FlushMappedBufferRange(target int32, offset int64, length int64) {
	C.ptrFlushMappedBufferRange(gpFlushMappedBufferRange, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}
func FlushMappedNamedBufferRange(buffer uint32, offset int64, length int32) {
	C.ptrFlushMappedNamedBufferRange(gpFlushMappedNamedBufferRange, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(length))
}
func FramebufferParameteri(target int32, pname int32, param int32) {
	C.ptrFramebufferParameteri(gpFramebufferParameteri, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func FramebufferRenderbuffer(target int32, attachment int32, renderbuffertarget int32, renderbuffer uint32) {
	C.ptrFramebufferRenderbuffer(gpFramebufferRenderbuffer, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}
func FramebufferTexture(target int32, attachment int32, texture uint32, level int32) {
	C.ptrFramebufferTexture(gpFramebufferTexture, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture1D(target int32, attachment int32, textarget int32, texture uint32, level int32) {
	C.ptrFramebufferTexture1D(gpFramebufferTexture1D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture2D(target int32, attachment int32, textarget int32, texture uint32, level int32) {
	C.ptrFramebufferTexture2D(gpFramebufferTexture2D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture3D(target int32, attachment int32, textarget int32, texture uint32, level int32, zoffset int32) {
	C.ptrFramebufferTexture3D(gpFramebufferTexture3D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(zoffset))
}
func FramebufferTextureLayer(target int32, attachment int32, texture uint32, level int32, layer int32) {
	C.ptrFramebufferTextureLayer(gpFramebufferTextureLayer, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}
func FrontFace(mode int32) {
	C.ptrFrontFace(gpFrontFace, (C.GLenum)(mode))
}
func GenBuffers(n int32, buffers *uint32) {
	C.ptrGenBuffers(gpGenBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func GenFramebuffers(n int32, framebuffers *uint32) {
	C.ptrGenFramebuffers(gpGenFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}
func GenProgramPipelines(n int32, pipelines *uint32) {
	C.ptrGenProgramPipelines(gpGenProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}
func GenQueries(n int32, ids *uint32) {
	C.ptrGenQueries(gpGenQueries, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}
func GenRenderbuffers(n int32, renderbuffers *uint32) {
	C.ptrGenRenderbuffers(gpGenRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}
func GenSamplers(count int32, samplers *uint32) {
	C.ptrGenSamplers(gpGenSamplers, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}
func GenTextures(n int32, textures *uint32) {
	C.ptrGenTextures(gpGenTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}
func GenTransformFeedbacks(n int32, ids *uint32) {
	C.ptrGenTransformFeedbacks(gpGenTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}
func GenVertexArrays(n int32, arrays *uint32) {
	C.ptrGenVertexArrays(gpGenVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}
func GenerateMipmap(target int32) {
	C.ptrGenerateMipmap(gpGenerateMipmap, (C.GLenum)(target))
}
func GenerateTextureMipmap(texture uint32) {
	C.ptrGenerateTextureMipmap(gpGenerateTextureMipmap, (C.GLuint)(texture))
}
func GetActiveAtomicCounterBufferiv(program uint32, bufferIndex uint32, pname int32, params *int32) {
	C.ptrGetActiveAtomicCounterBufferiv(gpGetActiveAtomicCounterBufferiv, (C.GLuint)(program), (C.GLuint)(bufferIndex), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetActiveAttrib(program uint32, index uint32, bufSize int32, length *int32, size *int32, gltype *int32, name string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrGetActiveAttrib(gpGetActiveAttrib, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(gltype)), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func GetActiveSubroutineName(program uint32, shadertype int32, index uint32, bufsize int32, length *int32, name string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrGetActiveSubroutineName(gpGetActiveSubroutineName, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func GetActiveSubroutineUniformName(program uint32, shadertype int32, index uint32, bufsize int32, length *int32, name string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrGetActiveSubroutineUniformName(gpGetActiveSubroutineUniformName, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func GetActiveSubroutineUniformiv(program uint32, shadertype int32, index uint32, pname int32, values *int32) {
	C.ptrGetActiveSubroutineUniformiv(gpGetActiveSubroutineUniformiv, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(values)))
}
func GetActiveUniform(program uint32, index uint32, bufSize int32, length *int32, size *int32, gltype *int32, name string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrGetActiveUniform(gpGetActiveUniform, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(gltype)), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func GetActiveUniformBlockName(program uint32, uniformBlockIndex uint32, bufSize int32, length *int32, uniformBlockName string) {
	cstruniformBlockName, freeuniformBlockName := stringToChar(uniformBlockName)
	defer freeuniformBlockName()
	C.ptrGetActiveUniformBlockName(gpGetActiveUniformBlockName, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstruniformBlockName)))
}
func GetActiveUniformBlockiv(program uint32, uniformBlockIndex uint32, pname int32, params *int32) {
	C.ptrGetActiveUniformBlockiv(gpGetActiveUniformBlockiv, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetActiveUniformName(program uint32, uniformIndex uint32, bufSize int32, length *int32, uniformName string) {
	cstruniformName, freeuniformName := stringToChar(uniformName)
	defer freeuniformName()
	C.ptrGetActiveUniformName(gpGetActiveUniformName, (C.GLuint)(program), (C.GLuint)(uniformIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstruniformName)))
}
func GetActiveUniformsiv(program uint32, uniformCount int32, uniformIndices *[]uint32, pname int32, params *[]int32) {
	C.ptrGetActiveUniformsiv(gpGetActiveUniformsiv, (C.GLuint)(program), (C.GLsizei)(uniformCount), (*C.GLuint)(unsafe.Pointer(uniformIndices)), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetAttachedShaders(program uint32, maxCount int32, count *int32, shaders *uint32) {
	C.ptrGetAttachedShaders(gpGetAttachedShaders, (C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
}
func GetAttribLocation(program uint32, name string) int32 {
	code := name + "\x00"
	cstrname := stringToUint8(code)
	ret := C.ptrGetAttribLocation(gpGetAttribLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return (int32)(ret)
}
func GetBooleani_v(target int32, index uint32, data bool) {
	bdata := boolToInt(data)
	C.ptrGetBooleani_v(gpGetBooleani_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLboolean)(unsafe.Pointer(&bdata)))
}
func GetBooleanv(pname int32, data bool) {
	bdata := boolToInt(data)
	C.ptrGetBooleanv(gpGetBooleanv, (C.GLenum)(pname), (*C.GLboolean)(unsafe.Pointer(&bdata)))
}
func GetBufferParameteri64v(target int32, pname int32, params *int64) {
	C.ptrGetBufferParameteri64v(gpGetBufferParameteri64v, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetBufferParameteriv(target int32, pname int32, params *int32) {
	C.ptrGetBufferParameteriv(gpGetBufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetBufferPointerv(target int32, pname int32, params *unsafe.Pointer) {
	C.ptrGetBufferPointerv(gpGetBufferPointerv, (C.GLenum)(target), (C.GLenum)(pname), params)
}
func GetBufferSubData(target int32, offset int64, size int64, data interface{}) {
	C.ptrGetBufferSubData(gpGetBufferSubData, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), ptr(data))
}
func GetCompressedTexImage(target int32, level int32, img unsafe.Pointer) {
	C.ptrGetCompressedTexImage(gpGetCompressedTexImage, (C.GLenum)(target), (C.GLint)(level), img)
}
func GetCompressedTextureImage(texture uint32, level int32, bufSize int32, pixels unsafe.Pointer) {
	C.ptrGetCompressedTextureImage(gpGetCompressedTextureImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLsizei)(bufSize), pixels)
}
func GetCompressedTextureSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, bufSize int32, pixels unsafe.Pointer) {
	C.ptrGetCompressedTextureSubImage(gpGetCompressedTextureSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLsizei)(bufSize), pixels)
}
func GetDebugMessageLog(count uint32, bufSize int32, sources *int32, types *int32, ids *uint32, severities *int32, lengths *int32, messageLog string) uint32 {
	cstrmessageLog, freemessageLog := stringToChar(messageLog)
	defer freemessageLog()
	ret := C.ptrGetDebugMessageLog(gpGetDebugMessageLog, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(cstrmessageLog)))
	return (uint32)(ret)
}
func GetDebugMessageLogARB(count uint32, bufSize int32, sources *int32, types *int32, ids *uint32, severities *int32, lengths *int32, messageLog string) uint32 {
	cstrmessageLog, freemessageLog := stringToChar(messageLog)
	defer freemessageLog()
	ret := C.ptrGetDebugMessageLogARB(gpGetDebugMessageLogARB, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(cstrmessageLog)))
	return (uint32)(ret)
}
func GetDebugMessageLogKHR(count uint32, bufSize int32, sources *int32, types *int32, ids *uint32, severities *int32, lengths *int32, messageLog string) uint32 {
	cstrmessageLog, freemessageLog := stringToChar(messageLog)
	defer freemessageLog()
	ret := C.ptrGetDebugMessageLogKHR(gpGetDebugMessageLogKHR, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(cstrmessageLog)))
	return (uint32)(ret)
}
func GetDoublei_v(target int32, index uint32, data *float64) {
	C.ptrGetDoublei_v(gpGetDoublei_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(data)))
}
func GetDoublev(pname int32, data *float64) {
	C.ptrGetDoublev(gpGetDoublev, (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(data)))
}
func GetError() int32 {
	ret := C.ptrGetError(gpGetError)
	return (int32)(ret)
}
func GetFloati_v(target int32, index uint32, data *float32) {
	C.ptrGetFloati_v(gpGetFloati_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(data)))
}
func GetFloatv(pname int32, data *float32) {
	C.ptrGetFloatv(gpGetFloatv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(data)))
}
func GetFragDataIndex(program uint32, name string) int32 {
	cstrname, freename := stringToChar(name)
	defer freename()
	ret := C.ptrGetFragDataIndex(gpGetFragDataIndex, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return (int32)(ret)
}
func GetFragDataLocation(program uint32, name string) int32 {
	cstrname, freename := stringToChar(name)
	defer freename()
	ret := C.ptrGetFragDataLocation(gpGetFragDataLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return (int32)(ret)
}
func GetFramebufferAttachmentParameteriv(target int32, attachment int32, pname int32, params *int32) {
	C.ptrGetFramebufferAttachmentParameteriv(gpGetFramebufferAttachmentParameteriv, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetFramebufferParameteriv(target int32, pname int32, params *int32) {
	C.ptrGetFramebufferParameteriv(gpGetFramebufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetGraphicsResetStatus() int32 {
	ret := C.ptrGetGraphicsResetStatus(gpGetGraphicsResetStatus)
	return (int32)(ret)
}
func GetGraphicsResetStatusARB() int32 {
	ret := C.ptrGetGraphicsResetStatusARB(gpGetGraphicsResetStatusARB)
	return (int32)(ret)
}
func GetGraphicsResetStatusKHR() int32 {
	ret := C.ptrGetGraphicsResetStatusKHR(gpGetGraphicsResetStatusKHR)
	return (int32)(ret)
}
func GetImageHandleARB(texture uint32, level int32, layered bool, layer int32, format int32) uint64 {
	ret := C.ptrGetImageHandleARB(gpGetImageHandleARB, (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(boolToInt(layered)), (C.GLint)(layer), (C.GLenum)(format))
	return (uint64)(ret)
}
func GetInteger64i_v(target int32, index uint32, data *int64) {
	C.ptrGetInteger64i_v(gpGetInteger64i_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLint64)(unsafe.Pointer(data)))
}
func GetInteger64v(pname int32, data *int64) {
	C.ptrGetInteger64v(gpGetInteger64v, (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(data)))
}
func GetIntegeri_v(target int32, index uint32, data *int32) {
	C.ptrGetIntegeri_v(gpGetIntegeri_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(data)))
}
func GetIntegerv(pname int32, data *int32) {
	C.ptrGetIntegerv(gpGetIntegerv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(data)))
}
func GetInternalformati64v(target int32, internalformat int32, pname int32, bufSize int32, params *int64) {
	C.ptrGetInternalformati64v(gpGetInternalformati64v, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetInternalformativ(target int32, internalformat int32, pname int32, bufSize int32, params *int32) {
	C.ptrGetInternalformativ(gpGetInternalformativ, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetMultisamplefv(pname int32, index uint32, val *float32) {
	C.ptrGetMultisamplefv(gpGetMultisamplefv, (C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(val)))
}
func GetNamedBufferParameteri64v(buffer uint32, pname int32, params *int64) {
	C.ptrGetNamedBufferParameteri64v(gpGetNamedBufferParameteri64v, (C.GLuint)(buffer), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetNamedBufferParameteriv(buffer uint32, pname int32, params *int32) {
	C.ptrGetNamedBufferParameteriv(gpGetNamedBufferParameteriv, (C.GLuint)(buffer), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetNamedBufferPointerv(buffer uint32, pname int32, params *unsafe.Pointer) {
	C.ptrGetNamedBufferPointerv(gpGetNamedBufferPointerv, (C.GLuint)(buffer), (C.GLenum)(pname), params)
}
func GetNamedBufferSubData(buffer uint32, offset int64, size int32, data interface{}) {
	C.ptrGetNamedBufferSubData(gpGetNamedBufferSubData, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), ptr(data))
}
func GetNamedFramebufferAttachmentParameteriv(framebuffer uint32, attachment int32, pname int32, params *int32) {
	C.ptrGetNamedFramebufferAttachmentParameteriv(gpGetNamedFramebufferAttachmentParameteriv, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetNamedFramebufferParameteriv(framebuffer uint32, pname int32, param *int32) {
	C.ptrGetNamedFramebufferParameteriv(gpGetNamedFramebufferParameteriv, (C.GLuint)(framebuffer), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func GetNamedRenderbufferParameteriv(renderbuffer uint32, pname int32, params *int32) {
	C.ptrGetNamedRenderbufferParameteriv(gpGetNamedRenderbufferParameteriv, (C.GLuint)(renderbuffer), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetNamedStringARB(namelen int32, name string, bufSize int32, stringlen *int32, glstring string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	cstrglstring, freeglstring := stringToChar(glstring)
	defer freeglstring()
	C.ptrGetNamedStringARB(gpGetNamedStringARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(cstrname)), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(stringlen)), (*C.GLchar)(unsafe.Pointer(cstrglstring)))
}
func GetNamedStringivARB(namelen int32, name string, pname int32, params *int32) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrGetNamedStringivARB(gpGetNamedStringivARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(cstrname)), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetObjectLabel(identifier int32, name uint32, bufSize int32, length *int32, label string) {
	cstrlabel, freelabel := stringToChar(label)
	defer freelabel()
	C.ptrGetObjectLabel(gpGetObjectLabel, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrlabel)))
}
func GetObjectLabelKHR(identifier int32, name uint32, bufSize int32, length *int32, label string) {
	cstrlabel, freelabel := stringToChar(label)
	defer freelabel()
	C.ptrGetObjectLabelKHR(gpGetObjectLabelKHR, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrlabel)))
}
func GetObjectPtrLabel(ptr unsafe.Pointer, bufSize int32, length *int32, label string) {
	cstrlabel, freelabel := stringToChar(label)
	defer freelabel()
	C.ptrGetObjectPtrLabel(gpGetObjectPtrLabel, ptr, (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrlabel)))
}
func GetObjectPtrLabelKHR(ptr unsafe.Pointer, bufSize int32, length *int32, label string) {
	cstrlabel, freelabel := stringToChar(label)
	defer freelabel()
	C.ptrGetObjectPtrLabelKHR(gpGetObjectPtrLabelKHR, ptr, (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrlabel)))
}
func GetPointerv(pname int32, params *unsafe.Pointer) {
	C.ptrGetPointerv(gpGetPointerv, (C.GLenum)(pname), params)
}
func GetPointervKHR(pname int32, params *unsafe.Pointer) {
	C.ptrGetPointervKHR(gpGetPointervKHR, (C.GLenum)(pname), params)
}
func GetProgramBinary(program uint32, bufSize int32, length *int32, binaryFormat *int32, binary unsafe.Pointer) {
	C.ptrGetProgramBinary(gpGetProgramBinary, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLenum)(unsafe.Pointer(binaryFormat)), binary)
}
func GetProgramInfoLog(program uint32, bufSize int32, length *int32) string {
	infoLog := strings.Repeat("\x00", int(bufSize+1))
	cstrinfoLog, freeinfoLog := stringToChar(infoLog)
	defer freeinfoLog()
	C.ptrGetProgramInfoLog(gpGetProgramInfoLog, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrinfoLog)))
	return C.GoString((*C.char)(unsafe.Pointer(unsafe.Pointer(cstrinfoLog))))
}
func GetProgramInterfaceiv(program uint32, programInterface int32, pname int32, params *int32) {
	C.ptrGetProgramInterfaceiv(gpGetProgramInterfaceiv, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetProgramPipelineInfoLog(pipeline uint32, bufSize int32, length *int32, infoLog string) {
	cstrinfoLog, freeinfoLog := stringToChar(infoLog)
	defer freeinfoLog()
	C.ptrGetProgramPipelineInfoLog(gpGetProgramPipelineInfoLog, (C.GLuint)(pipeline), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrinfoLog)))
}
func GetProgramPipelineiv(pipeline uint32, pname int32, params *int32) {
	C.ptrGetProgramPipelineiv(gpGetProgramPipelineiv, (C.GLuint)(pipeline), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetProgramResourceIndex(program uint32, programInterface int32, name string) uint32 {
	cstrname, freename := stringToChar(name)
	defer freename()
	ret := C.ptrGetProgramResourceIndex(gpGetProgramResourceIndex, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return (uint32)(ret)
}
func GetProgramResourceLocation(program uint32, programInterface int32, name string) int32 {
	cstrname, freename := stringToChar(name)
	defer freename()
	ret := C.ptrGetProgramResourceLocation(gpGetProgramResourceLocation, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return (int32)(ret)
}
func GetProgramResourceLocationIndex(program uint32, programInterface int32, name string) int32 {
	cstrname, freename := stringToChar(name)
	defer freename()
	ret := C.ptrGetProgramResourceLocationIndex(gpGetProgramResourceLocationIndex, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return (int32)(ret)
}
func GetProgramResourceName(program uint32, programInterface int32, index uint32, bufSize int32, length *int32, name string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrGetProgramResourceName(gpGetProgramResourceName, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func GetProgramResourceiv(program uint32, programInterface int32, index uint32, propCount int32, props *int32, bufSize int32, length *int32, params *int32) {
	C.ptrGetProgramResourceiv(gpGetProgramResourceiv, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(propCount), (*C.GLenum)(unsafe.Pointer(props)), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(params)))
}
func GetProgramStageiv(program uint32, shadertype int32, pname int32, values *int32) {
	C.ptrGetProgramStageiv(gpGetProgramStageiv, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(values)))
}
func GetProgramiv(program uint32, pname int32, params *int32) {
	C.ptrGetProgramiv(gpGetProgramiv, (C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryIndexediv(target int32, index uint32, pname int32, params *int32) {
	C.ptrGetQueryIndexediv(gpGetQueryIndexediv, (C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryObjecti64v(id uint32, pname int32, params *int64) {
	C.ptrGetQueryObjecti64v(gpGetQueryObjecti64v, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetQueryObjectiv(id uint32, pname int32, params *int32) {
	C.ptrGetQueryObjectiv(gpGetQueryObjectiv, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryObjectui64v(id uint32, pname int32, params *uint64) {
	C.ptrGetQueryObjectui64v(gpGetQueryObjectui64v, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint64)(unsafe.Pointer(params)))
}
func GetQueryObjectuiv(id uint32, pname int32, params *uint32) {
	C.ptrGetQueryObjectuiv(gpGetQueryObjectuiv, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetQueryiv(target int32, pname int32, params *int32) {
	C.ptrGetQueryiv(gpGetQueryiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetRenderbufferParameteriv(target int32, pname int32, params *int32) {
	C.ptrGetRenderbufferParameteriv(gpGetRenderbufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIiv(sampler uint32, pname int32, params *int32) {
	C.ptrGetSamplerParameterIiv(gpGetSamplerParameterIiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIuiv(sampler uint32, pname int32, params *uint32) {
	C.ptrGetSamplerParameterIuiv(gpGetSamplerParameterIuiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetSamplerParameterfv(sampler uint32, pname int32, params *float32) {
	C.ptrGetSamplerParameterfv(gpGetSamplerParameterfv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetSamplerParameteriv(sampler uint32, pname int32, params *int32) {
	C.ptrGetSamplerParameteriv(gpGetSamplerParameteriv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetShaderInfoLog(shader uint32, bufSize int32, length *int32) string {
	infoLog := strings.Repeat("\x00", int(bufSize+1))
	cstrinfoLog, freeinfoLog := stringToChar(infoLog)
	defer freeinfoLog()
	C.ptrGetShaderInfoLog(gpGetShaderInfoLog, (C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrinfoLog)))
	return C.GoString((*C.char)(unsafe.Pointer(unsafe.Pointer(cstrinfoLog))))
}
func GetShaderPrecisionFormat(shadertype int32, precisiontype int32, glrange *int32, precision *int32) {
	C.ptrGetShaderPrecisionFormat(gpGetShaderPrecisionFormat, (C.GLenum)(shadertype), (C.GLenum)(precisiontype), (*C.GLint)(unsafe.Pointer(glrange)), (*C.GLint)(unsafe.Pointer(precision)))
}
func GetShaderSource(shader uint32, bufSize int32, length *int32, source string) {
	cstrsource, freesource := stringToChar(source)
	defer freesource()
	C.ptrGetShaderSource(gpGetShaderSource, (C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(cstrsource)))
}
func GetShaderiv(shader uint32, pname int32, params *int32) {
	C.ptrGetShaderiv(gpGetShaderiv, (C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetString(name int32) string {
	ret := C.ptrGetString(gpGetString, (C.GLenum)(name))
	return charToString((*uint8)(ret))
}
func GetStringi(name int32, index uint32) string {
	ret := C.ptrGetStringi(gpGetStringi, (C.GLenum)(name), (C.GLuint)(index))
	return charToString((*uint8)(ret))
}
func GetSubroutineIndex(program uint32, shadertype int32, name string) uint32 {
	cstrname, freename := stringToChar(name)
	defer freename()
	ret := C.ptrGetSubroutineIndex(gpGetSubroutineIndex, (C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return (uint32)(ret)
}
func GetSubroutineUniformLocation(program uint32, shadertype int32, name string) int32 {
	cstrname, freename := stringToChar(name)
	defer freename()
	ret := C.ptrGetSubroutineUniformLocation(gpGetSubroutineUniformLocation, (C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return (int32)(ret)
}
func GetSynciv(sync unsafe.Pointer, pname int32, bufSize int32, length *int32, values *int32) {
	C.ptrGetSynciv(gpGetSynciv, (C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(values)))
}
func GetTexImage(target int32, level int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrGetTexImage(gpGetTexImage, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func GetTexLevelParameterfv(target int32, level int32, pname int32, params *float32) {
	C.ptrGetTexLevelParameterfv(gpGetTexLevelParameterfv, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexLevelParameteriv(target int32, level int32, pname int32, params *int32) {
	C.ptrGetTexLevelParameteriv(gpGetTexLevelParameteriv, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexParameterIiv(target int32, pname int32, params *int32) {
	C.ptrGetTexParameterIiv(gpGetTexParameterIiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexParameterIuiv(target int32, pname int32, params *uint32) {
	C.ptrGetTexParameterIuiv(gpGetTexParameterIuiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetTexParameterfv(target int32, pname int32, params *float32) {
	C.ptrGetTexParameterfv(gpGetTexParameterfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexParameteriv(target int32, pname int32, params *int32) {
	C.ptrGetTexParameteriv(gpGetTexParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureHandleARB(texture uint32) uint64 {
	ret := C.ptrGetTextureHandleARB(gpGetTextureHandleARB, (C.GLuint)(texture))
	return (uint64)(ret)
}
func GetTextureImage(texture uint32, level int32, format int32, gltype int32, bufSize int32, pixels unsafe.Pointer) {
	C.ptrGetTextureImage(gpGetTextureImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(gltype), (C.GLsizei)(bufSize), pixels)
}
func GetTextureLevelParameterfv(texture uint32, level int32, pname int32, params *float32) {
	C.ptrGetTextureLevelParameterfv(gpGetTextureLevelParameterfv, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTextureLevelParameteriv(texture uint32, level int32, pname int32, params *int32) {
	C.ptrGetTextureLevelParameteriv(gpGetTextureLevelParameteriv, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureParameterIiv(texture uint32, pname int32, params *int32) {
	C.ptrGetTextureParameterIiv(gpGetTextureParameterIiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureParameterIuiv(texture uint32, pname int32, params *uint32) {
	C.ptrGetTextureParameterIuiv(gpGetTextureParameterIuiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetTextureParameterfv(texture uint32, pname int32, params *float32) {
	C.ptrGetTextureParameterfv(gpGetTextureParameterfv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTextureParameteriv(texture uint32, pname int32, params *int32) {
	C.ptrGetTextureParameteriv(gpGetTextureParameteriv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureSamplerHandleARB(texture uint32, sampler uint32) uint64 {
	ret := C.ptrGetTextureSamplerHandleARB(gpGetTextureSamplerHandleARB, (C.GLuint)(texture), (C.GLuint)(sampler))
	return (uint64)(ret)
}
func GetTextureSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format int32, gltype int32, bufSize int32, pixels unsafe.Pointer) {
	C.ptrGetTextureSubImage(gpGetTextureSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(gltype), (C.GLsizei)(bufSize), pixels)
}
func GetTransformFeedbackVarying(program uint32, index uint32, bufSize int32, length *int32, size *int32, gltype *int32, name string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	C.ptrGetTransformFeedbackVarying(gpGetTransformFeedbackVarying, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLsizei)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(gltype)), (*C.GLchar)(unsafe.Pointer(cstrname)))
}
func GetTransformFeedbacki64_v(xfb uint32, pname int32, index uint32, param *int64) {
	C.ptrGetTransformFeedbacki64_v(gpGetTransformFeedbacki64_v, (C.GLuint)(xfb), (C.GLenum)(pname), (C.GLuint)(index), (*C.GLint64)(unsafe.Pointer(param)))
}
func GetTransformFeedbacki_v(xfb uint32, pname int32, index uint32, param *int32) {
	C.ptrGetTransformFeedbacki_v(gpGetTransformFeedbacki_v, (C.GLuint)(xfb), (C.GLenum)(pname), (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(param)))
}
func GetTransformFeedbackiv(xfb uint32, pname int32, param *int32) {
	C.ptrGetTransformFeedbackiv(gpGetTransformFeedbackiv, (C.GLuint)(xfb), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func GetUniformBlockIndex(program uint32, uniformBlockName string) uint32 {
	code := uniformBlockName + "\x00"
	cstruniformBlockName := stringToUint8(code)
	ret := C.ptrGetUniformBlockIndex(gpGetUniformBlockIndex, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(cstruniformBlockName)))
	return (uint32)(ret)
}
func GetUniformIndices(program uint32, uniformCount int32, uniformNames []string, uniformIndices *uint32) {
	fmt.Println("GetUniformIndices=>", program, uniformCount, uniformNames, uniformIndices, len(uniformNames))
	uniformNamesArray := make([]*uint8, len(uniformNames))
	for i := range uniformNames {
		code := uniformNames[i] + "\x00"
		uniformNamesArray[i] = stringToUint8(code)
	}
	C.ptrGetUniformIndices(gpGetUniformIndices, (C.GLuint)(program), (C.GLsizei)(uniformCount), (**C.GLchar)(unsafe.Pointer(&uniformNamesArray)), (*C.GLuint)(unsafe.Pointer(uniformIndices)))
}
func GetUniformLocation(program uint32, name string) int32 {
	code := name + "\x00"
	cstrname := stringToUint8(code)
	ret := C.ptrGetUniformLocation(gpGetUniformLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return (int32)(ret)
}
func GetUniformSubroutineuiv(shadertype int32, location int32, params *uint32) {
	C.ptrGetUniformSubroutineuiv(gpGetUniformSubroutineuiv, (C.GLenum)(shadertype), (C.GLint)(location), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetUniformdv(program uint32, location int32, params *float64) {
	C.ptrGetUniformdv(gpGetUniformdv, (C.GLuint)(program), (C.GLint)(location), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetUniformfv(program uint32, location int32, params *float32) {
	C.ptrGetUniformfv(gpGetUniformfv, (C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetUniformiv(program uint32, location int32, params *int32) {
	C.ptrGetUniformiv(gpGetUniformiv, (C.GLuint)(program), (C.GLint)(location), (*C.GLint)(unsafe.Pointer(params)))
}
func GetUniformuiv(program uint32, location int32, params *uint32) {
	C.ptrGetUniformuiv(gpGetUniformuiv, (C.GLuint)(program), (C.GLint)(location), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetVertexArrayIndexed64iv(vaobj uint32, index uint32, pname int32, param *int64) {
	C.ptrGetVertexArrayIndexed64iv(gpGetVertexArrayIndexed64iv, (C.GLuint)(vaobj), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(param)))
}
func GetVertexArrayIndexediv(vaobj uint32, index uint32, pname int32, param *int32) {
	C.ptrGetVertexArrayIndexediv(gpGetVertexArrayIndexediv, (C.GLuint)(vaobj), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func GetVertexArrayiv(vaobj uint32, pname int32, param *int32) {
	C.ptrGetVertexArrayiv(gpGetVertexArrayiv, (C.GLuint)(vaobj), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func GetVertexAttribIiv(index uint32, pname int32, params *int32) {
	C.ptrGetVertexAttribIiv(gpGetVertexAttribIiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetVertexAttribIuiv(index uint32, pname int32, params *uint32) {
	C.ptrGetVertexAttribIuiv(gpGetVertexAttribIuiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetVertexAttribLdv(index uint32, pname int32, params *float64) {
	C.ptrGetVertexAttribLdv(gpGetVertexAttribLdv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetVertexAttribLui64vARB(index uint32, pname int32, params *uint64) {
	C.ptrGetVertexAttribLui64vARB(gpGetVertexAttribLui64vARB, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint64EXT)(unsafe.Pointer(params)))
}
func GetVertexAttribPointerv(index uint32, pname int32, pointer *unsafe.Pointer) {
	C.ptrGetVertexAttribPointerv(gpGetVertexAttribPointerv, (C.GLuint)(index), (C.GLenum)(pname), pointer)
}
func GetVertexAttribdv(index uint32, pname int32, params *float64) {
	C.ptrGetVertexAttribdv(gpGetVertexAttribdv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetVertexAttribfv(index uint32, pname int32, params *float32) {
	C.ptrGetVertexAttribfv(gpGetVertexAttribfv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetVertexAttribiv(index uint32, pname int32, params *int32) {
	C.ptrGetVertexAttribiv(gpGetVertexAttribiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnCompressedTexImage(target int32, lod int32, bufSize int32, pixels unsafe.Pointer) {
	C.ptrGetnCompressedTexImage(gpGetnCompressedTexImage, (C.GLenum)(target), (C.GLint)(lod), (C.GLsizei)(bufSize), pixels)
}
func GetnCompressedTexImageARB(target int32, lod int32, bufSize int32, img unsafe.Pointer) {
	C.ptrGetnCompressedTexImageARB(gpGetnCompressedTexImageARB, (C.GLenum)(target), (C.GLint)(lod), (C.GLsizei)(bufSize), img)
}
func GetnTexImage(target int32, level int32, format int32, gltype int32, bufSize int32, pixels unsafe.Pointer) {
	C.ptrGetnTexImage(gpGetnTexImage, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(gltype), (C.GLsizei)(bufSize), pixels)
}
func GetnTexImageARB(target int32, level int32, format int32, gltype int32, bufSize int32, img unsafe.Pointer) {
	C.ptrGetnTexImageARB(gpGetnTexImageARB, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(gltype), (C.GLsizei)(bufSize), img)
}
func GetnUniformdv(program uint32, location int32, bufSize int32, params *float64) {
	C.ptrGetnUniformdv(gpGetnUniformdv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetnUniformdvARB(program uint32, location int32, bufSize int32, params *float64) {
	C.ptrGetnUniformdvARB(gpGetnUniformdvARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetnUniformfv(program uint32, location int32, bufSize int32, params *float32) {
	C.ptrGetnUniformfv(gpGetnUniformfv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformfvARB(program uint32, location int32, bufSize int32, params *float32) {
	C.ptrGetnUniformfvARB(gpGetnUniformfvARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformfvKHR(program uint32, location int32, bufSize int32, params *float32) {
	C.ptrGetnUniformfvKHR(gpGetnUniformfvKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformiv(program uint32, location int32, bufSize int32, params *int32) {
	C.ptrGetnUniformiv(gpGetnUniformiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformivARB(program uint32, location int32, bufSize int32, params *int32) {
	C.ptrGetnUniformivARB(gpGetnUniformivARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformivKHR(program uint32, location int32, bufSize int32, params *int32) {
	C.ptrGetnUniformivKHR(gpGetnUniformivKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformuiv(program uint32, location int32, bufSize int32, params *uint32) {
	C.ptrGetnUniformuiv(gpGetnUniformuiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetnUniformuivARB(program uint32, location int32, bufSize int32, params *uint32) {
	C.ptrGetnUniformuivARB(gpGetnUniformuivARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetnUniformuivKHR(program uint32, location int32, bufSize int32, params *uint32) {
	C.ptrGetnUniformuivKHR(gpGetnUniformuivKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}
func Hint(target int32, mode int32) {
	C.ptrHint(gpHint, (C.GLenum)(target), (C.GLenum)(mode))
}
func InvalidateBufferData(buffer uint32) {
	C.ptrInvalidateBufferData(gpInvalidateBufferData, (C.GLuint)(buffer))
}
func InvalidateBufferSubData(buffer uint32, offset int64, length int64) {
	C.ptrInvalidateBufferSubData(gpInvalidateBufferSubData, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}
func InvalidateFramebuffer(target int32, numAttachments int32, attachments *int32) {
	C.ptrInvalidateFramebuffer(gpInvalidateFramebuffer, (C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)))
}
func InvalidateNamedFramebufferData(framebuffer uint32, numAttachments int32, attachments *int32) {
	C.ptrInvalidateNamedFramebufferData(gpInvalidateNamedFramebufferData, (C.GLuint)(framebuffer), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)))
}
func InvalidateNamedFramebufferSubData(framebuffer uint32, numAttachments int32, attachments *int32, x int32, y int32, width int32, height int32) {
	C.ptrInvalidateNamedFramebufferSubData(gpInvalidateNamedFramebufferSubData, (C.GLuint)(framebuffer), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func InvalidateSubFramebuffer(target int32, numAttachments int32, attachments *int32, x int32, y int32, width int32, height int32) {
	C.ptrInvalidateSubFramebuffer(gpInvalidateSubFramebuffer, (C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func InvalidateTexImage(texture uint32, level int32) {
	C.ptrInvalidateTexImage(gpInvalidateTexImage, (C.GLuint)(texture), (C.GLint)(level))
}
func InvalidateTexSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32) {
	C.ptrInvalidateTexSubImage(gpInvalidateTexSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
func IsBuffer(buffer uint32) bool {
	ret := C.ptrIsBuffer(gpIsBuffer, (C.GLuint)(buffer))
	return intToBool((int8)(ret))
}
func IsEnabled(cap int32) bool {
	ret := C.ptrIsEnabled(gpIsEnabled, (C.GLenum)(cap))
	return intToBool((int8)(ret))
}
func IsEnabledi(target int32, index uint32) bool {
	ret := C.ptrIsEnabledi(gpIsEnabledi, (C.GLenum)(target), (C.GLuint)(index))
	return intToBool((int8)(ret))
}
func IsFramebuffer(framebuffer uint32) bool {
	ret := C.ptrIsFramebuffer(gpIsFramebuffer, (C.GLuint)(framebuffer))
	return intToBool((int8)(ret))
}
func IsImageHandleResidentARB(handle uint64) bool {
	ret := C.ptrIsImageHandleResidentARB(gpIsImageHandleResidentARB, (C.GLuint64)(handle))
	return intToBool((int8)(ret))
}
func IsNamedStringARB(namelen int32, name string) bool {
	cstrname, freename := stringToChar(name)
	defer freename()
	ret := C.ptrIsNamedStringARB(gpIsNamedStringARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(cstrname)))
	return intToBool((int8)(ret))
}
func IsProgram(program uint32) bool {
	ret := C.ptrIsProgram(gpIsProgram, (C.GLuint)(program))
	return intToBool((int8)(ret))
}
func IsProgramPipeline(pipeline uint32) bool {
	ret := C.ptrIsProgramPipeline(gpIsProgramPipeline, (C.GLuint)(pipeline))
	return intToBool((int8)(ret))
}
func IsQuery(id uint32) bool {
	ret := C.ptrIsQuery(gpIsQuery, (C.GLuint)(id))
	return intToBool((int8)(ret))
}
func IsRenderbuffer(renderbuffer uint32) bool {
	ret := C.ptrIsRenderbuffer(gpIsRenderbuffer, (C.GLuint)(renderbuffer))
	return intToBool((int8)(ret))
}
func IsSampler(sampler uint32) bool {
	ret := C.ptrIsSampler(gpIsSampler, (C.GLuint)(sampler))
	return intToBool((int8)(ret))
}
func IsShader(shader uint32) bool {
	ret := C.ptrIsShader(gpIsShader, (C.GLuint)(shader))
	return intToBool((int8)(ret))
}
func IsSync(sync unsafe.Pointer) bool {
	ret := C.ptrIsSync(gpIsSync, (C.GLsync)(sync))
	return intToBool((int8)(ret))
}
func IsTexture(texture uint32) bool {
	ret := C.ptrIsTexture(gpIsTexture, (C.GLuint)(texture))
	return intToBool((int8)(ret))
}
func IsTextureHandleResidentARB(handle uint64) bool {
	ret := C.ptrIsTextureHandleResidentARB(gpIsTextureHandleResidentARB, (C.GLuint64)(handle))
	return intToBool((int8)(ret))
}
func IsTransformFeedback(id uint32) bool {
	ret := C.ptrIsTransformFeedback(gpIsTransformFeedback, (C.GLuint)(id))
	return intToBool((int8)(ret))
}
func IsVertexArray(array uint32) bool {
	ret := C.ptrIsVertexArray(gpIsVertexArray, (C.GLuint)(array))
	return intToBool((int8)(ret))
}
func LineWidth(width float32) {
	C.ptrLineWidth(gpLineWidth, (C.GLfloat)(width))
}
func LinkProgram(program uint32) {
	C.ptrLinkProgram(gpLinkProgram, (C.GLuint)(program))
}
func LogicOp(opcode int32) {
	C.ptrLogicOp(gpLogicOp, (C.GLenum)(opcode))
}
func MakeImageHandleNonResidentARB(handle uint64) {
	C.ptrMakeImageHandleNonResidentARB(gpMakeImageHandleNonResidentARB, (C.GLuint64)(handle))
}
func MakeImageHandleResidentARB(handle uint64, access int32) {
	C.ptrMakeImageHandleResidentARB(gpMakeImageHandleResidentARB, (C.GLuint64)(handle), (C.GLenum)(access))
}
func MakeTextureHandleNonResidentARB(handle uint64) {
	C.ptrMakeTextureHandleNonResidentARB(gpMakeTextureHandleNonResidentARB, (C.GLuint64)(handle))
}
func MakeTextureHandleResidentARB(handle uint64) {
	C.ptrMakeTextureHandleResidentARB(gpMakeTextureHandleResidentARB, (C.GLuint64)(handle))
}
func MapBuffer(target int32, access int32) unsafe.Pointer {
	ret := C.ptrMapBuffer(gpMapBuffer, (C.GLenum)(target), (C.GLenum)(access))
	return (unsafe.Pointer)(ret)
}
func MapBufferRange(target int32, offset int64, length int64, access uint32) unsafe.Pointer {
	ret := C.ptrMapBufferRange(gpMapBufferRange, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access))
	return (unsafe.Pointer)(ret)
}
func MapNamedBuffer(buffer uint32, access int32) unsafe.Pointer {
	ret := C.ptrMapNamedBuffer(gpMapNamedBuffer, (C.GLuint)(buffer), (C.GLenum)(access))
	return (unsafe.Pointer)(ret)
}
func MapNamedBufferRange(buffer uint32, offset int64, length int32, access uint32) unsafe.Pointer {
	ret := C.ptrMapNamedBufferRange(gpMapNamedBufferRange, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(length), (C.GLbitfield)(access))
	return (unsafe.Pointer)(ret)
}
func MemoryBarrier(barriers uint32) {
	C.ptrMemoryBarrier(gpMemoryBarrier, (C.GLbitfield)(barriers))
}
func MemoryBarrierByRegion(barriers uint32) {
	C.ptrMemoryBarrierByRegion(gpMemoryBarrierByRegion, (C.GLbitfield)(barriers))
}
func MinSampleShading(value float32) {
	C.ptrMinSampleShading(gpMinSampleShading, (C.GLfloat)(value))
}
func MinSampleShadingARB(value float32) {
	C.ptrMinSampleShadingARB(gpMinSampleShadingARB, (C.GLfloat)(value))
}
func MultiDrawArrays(mode int32, first *int32, count *int32, drawcount int32) {
	C.ptrMultiDrawArrays(gpMultiDrawArrays, (C.GLenum)(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLsizei)(drawcount))
}
func MultiDrawArraysIndirect(mode int32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	C.ptrMultiDrawArraysIndirect(gpMultiDrawArraysIndirect, (C.GLenum)(mode), indirect, (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}
func MultiDrawArraysIndirectCountARB(mode int32, indirect int64, drawcount int64, maxdrawcount int32, stride int32) {
	C.ptrMultiDrawArraysIndirectCountARB(gpMultiDrawArraysIndirectCountARB, (C.GLenum)(mode), (C.GLintptr)(indirect), (C.GLintptr)(drawcount), (C.GLsizei)(maxdrawcount), (C.GLsizei)(stride))
}
func MultiDrawElements(mode int32, count *int32, gltype int32, indices *unsafe.Pointer, drawcount int32) {
	C.ptrMultiDrawElements(gpMultiDrawElements, (C.GLenum)(mode), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLenum)(gltype), indices, (C.GLsizei)(drawcount))
}
func MultiDrawElementsBaseVertex(mode int32, count *int32, gltype int32, indices *unsafe.Pointer, drawcount int32, basevertex *int32) {
	C.ptrMultiDrawElementsBaseVertex(gpMultiDrawElementsBaseVertex, (C.GLenum)(mode), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLenum)(gltype), indices, (C.GLsizei)(drawcount), (*C.GLint)(unsafe.Pointer(basevertex)))
}
func MultiDrawElementsIndirect(mode int32, gltype int32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	C.ptrMultiDrawElementsIndirect(gpMultiDrawElementsIndirect, (C.GLenum)(mode), (C.GLenum)(gltype), indirect, (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}
func MultiDrawElementsIndirectCountARB(mode int32, gltype int32, indirect int64, drawcount int64, maxdrawcount int32, stride int32) {
	C.ptrMultiDrawElementsIndirectCountARB(gpMultiDrawElementsIndirectCountARB, (C.GLenum)(mode), (C.GLenum)(gltype), (C.GLintptr)(indirect), (C.GLintptr)(drawcount), (C.GLsizei)(maxdrawcount), (C.GLsizei)(stride))
}
func NamedBufferData(buffer uint32, size int32, data interface{}, usage int32) {
	C.ptrNamedBufferData(gpNamedBufferData, (C.GLuint)(buffer), (C.GLsizei)(size), ptr(data), (C.GLenum)(usage))
}
func NamedBufferPageCommitmentARB(buffer uint32, offset int64, size int32, commit bool) {
	C.ptrNamedBufferPageCommitmentARB(gpNamedBufferPageCommitmentARB, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLboolean)(boolToInt(commit)))
}
func NamedBufferPageCommitmentEXT(buffer uint32, offset int64, size int32, commit bool) {
	C.ptrNamedBufferPageCommitmentEXT(gpNamedBufferPageCommitmentEXT, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLboolean)(boolToInt(commit)))
}
func NamedBufferStorage(buffer uint32, size int32, data interface{}, flags uint32) {
	C.ptrNamedBufferStorage(gpNamedBufferStorage, (C.GLuint)(buffer), (C.GLsizei)(size), ptr(data), (C.GLbitfield)(flags))
}
func NamedBufferSubData(buffer uint32, offset int64, size int32, data interface{}) {
	C.ptrNamedBufferSubData(gpNamedBufferSubData, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), ptr(data))
}
func NamedFramebufferDrawBuffer(framebuffer uint32, buf int32) {
	C.ptrNamedFramebufferDrawBuffer(gpNamedFramebufferDrawBuffer, (C.GLuint)(framebuffer), (C.GLenum)(buf))
}
func NamedFramebufferDrawBuffers(framebuffer uint32, n int32, bufs *int32) {
	C.ptrNamedFramebufferDrawBuffers(gpNamedFramebufferDrawBuffers, (C.GLuint)(framebuffer), (C.GLsizei)(n), (*C.GLenum)(unsafe.Pointer(bufs)))
}
func NamedFramebufferParameteri(framebuffer uint32, pname int32, param int32) {
	C.ptrNamedFramebufferParameteri(gpNamedFramebufferParameteri, (C.GLuint)(framebuffer), (C.GLenum)(pname), (C.GLint)(param))
}
func NamedFramebufferReadBuffer(framebuffer uint32, src int32) {
	C.ptrNamedFramebufferReadBuffer(gpNamedFramebufferReadBuffer, (C.GLuint)(framebuffer), (C.GLenum)(src))
}
func NamedFramebufferRenderbuffer(framebuffer uint32, attachment int32, renderbuffertarget int32, renderbuffer uint32) {
	C.ptrNamedFramebufferRenderbuffer(gpNamedFramebufferRenderbuffer, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}
func NamedFramebufferTexture(framebuffer uint32, attachment int32, texture uint32, level int32) {
	C.ptrNamedFramebufferTexture(gpNamedFramebufferTexture, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}
func NamedFramebufferTextureLayer(framebuffer uint32, attachment int32, texture uint32, level int32, layer int32) {
	C.ptrNamedFramebufferTextureLayer(gpNamedFramebufferTextureLayer, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}
func NamedRenderbufferStorage(renderbuffer uint32, internalformat int32, width int32, height int32) {
	C.ptrNamedRenderbufferStorage(gpNamedRenderbufferStorage, (C.GLuint)(renderbuffer), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func NamedRenderbufferStorageMultisample(renderbuffer uint32, samples int32, internalformat int32, width int32, height int32) {
	C.ptrNamedRenderbufferStorageMultisample(gpNamedRenderbufferStorageMultisample, (C.GLuint)(renderbuffer), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func NamedStringARB(gltype int32, namelen int32, name string, stringlen int32, glstring string) {
	cstrname, freename := stringToChar(name)
	defer freename()
	cstrglstring, freeglstring := stringToChar(glstring)
	defer freeglstring()
	C.ptrNamedStringARB(gpNamedStringARB, (C.GLenum)(gltype), (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(cstrname)), (C.GLint)(stringlen), (*C.GLchar)(unsafe.Pointer(cstrglstring)))
}
func ObjectLabel(identifier int32, name uint32, length int32, label string) {
	cstrlabel, freelabel := stringToChar(label)
	defer freelabel()
	C.ptrObjectLabel(gpObjectLabel, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(cstrlabel)))
}
func ObjectLabelKHR(identifier int32, name uint32, length int32, label string) {
	cstrlabel, freelabel := stringToChar(label)
	defer freelabel()
	C.ptrObjectLabelKHR(gpObjectLabelKHR, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(cstrlabel)))
}
func ObjectPtrLabel(ptr unsafe.Pointer, length int32, label string) {
	cstrlabel, freelabel := stringToChar(label)
	defer freelabel()
	C.ptrObjectPtrLabel(gpObjectPtrLabel, ptr, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(cstrlabel)))
}
func ObjectPtrLabelKHR(ptr unsafe.Pointer, length int32, label string) {
	cstrlabel, freelabel := stringToChar(label)
	defer freelabel()
	C.ptrObjectPtrLabelKHR(gpObjectPtrLabelKHR, ptr, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(cstrlabel)))
}
func PatchParameterfv(pname int32, values *float32) {
	C.ptrPatchParameterfv(gpPatchParameterfv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(values)))
}
func PatchParameteri(pname int32, value int32) {
	C.ptrPatchParameteri(gpPatchParameteri, (C.GLenum)(pname), (C.GLint)(value))
}
func PauseTransformFeedback() {
	C.ptrPauseTransformFeedback(gpPauseTransformFeedback)
}
func PixelStoref(pname int32, param float32) {
	C.ptrPixelStoref(gpPixelStoref, (C.GLenum)(pname), (C.GLfloat)(param))
}
func PixelStorei(pname int32, param int32) {
	C.ptrPixelStorei(gpPixelStorei, (C.GLenum)(pname), (C.GLint)(param))
}
func PointParameterf(pname int32, param float32) {
	C.ptrPointParameterf(gpPointParameterf, (C.GLenum)(pname), (C.GLfloat)(param))
}
func PointParameterfv(pname int32, params *float32) {
	C.ptrPointParameterfv(gpPointParameterfv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func PointParameteri(pname int32, param int32) {
	C.ptrPointParameteri(gpPointParameteri, (C.GLenum)(pname), (C.GLint)(param))
}
func PointParameteriv(pname int32, params *int32) {
	C.ptrPointParameteriv(gpPointParameteriv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func PointSize(size float32) {
	C.ptrPointSize(gpPointSize, (C.GLfloat)(size))
}
func PolygonMode(face int32, mode int32) {
	C.ptrPolygonMode(gpPolygonMode, (C.GLenum)(face), (C.GLenum)(mode))
}
func PolygonOffset(factor float32, units float32) {
	C.ptrPolygonOffset(gpPolygonOffset, (C.GLfloat)(factor), (C.GLfloat)(units))
}
func PopDebugGroup() {
	C.ptrPopDebugGroup(gpPopDebugGroup)
}
func PopDebugGroupKHR() {
	C.ptrPopDebugGroupKHR(gpPopDebugGroupKHR)
}
func PrimitiveRestartIndex(index uint32) {
	C.ptrPrimitiveRestartIndex(gpPrimitiveRestartIndex, (C.GLuint)(index))
}
func ProgramBinary(program uint32, binaryFormat int32, binary unsafe.Pointer, length int32) {
	C.ptrProgramBinary(gpProgramBinary, (C.GLuint)(program), (C.GLenum)(binaryFormat), binary, (C.GLsizei)(length))
}
func ProgramParameteri(program uint32, pname int32, value int32) {
	C.ptrProgramParameteri(gpProgramParameteri, (C.GLuint)(program), (C.GLenum)(pname), (C.GLint)(value))
}
func ProgramUniform1d(program uint32, location int32, v0 float64) {
	C.ptrProgramUniform1d(gpProgramUniform1d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0))
}
func ProgramUniform1dv(program uint32, location int32, count int32, value *float64) {
	C.ptrProgramUniform1dv(gpProgramUniform1dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniform1f(program uint32, location int32, v0 float32) {
	C.ptrProgramUniform1f(gpProgramUniform1f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0))
}
func ProgramUniform1fv(program uint32, location int32, count int32, value *float32) {
	C.ptrProgramUniform1fv(gpProgramUniform1fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform1i(program uint32, location int32, v0 int32) {
	C.ptrProgramUniform1i(gpProgramUniform1i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0))
}
func ProgramUniform1iv(program uint32, location int32, count int32, value *int32) {
	C.ptrProgramUniform1iv(gpProgramUniform1iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform1ui(program uint32, location int32, v0 uint32) {
	C.ptrProgramUniform1ui(gpProgramUniform1ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0))
}
func ProgramUniform1uiv(program uint32, location int32, count int32, value *uint32) {
	C.ptrProgramUniform1uiv(gpProgramUniform1uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform2d(program uint32, location int32, v0 float64, v1 float64) {
	C.ptrProgramUniform2d(gpProgramUniform2d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1))
}
func ProgramUniform2dv(program uint32, location int32, count int32, value *float64) {
	C.ptrProgramUniform2dv(gpProgramUniform2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniform2f(program uint32, location int32, v0 float32, v1 float32) {
	C.ptrProgramUniform2f(gpProgramUniform2f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
func ProgramUniform2fv(program uint32, location int32, count int32, value *float32) {
	C.ptrProgramUniform2fv(gpProgramUniform2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform2i(program uint32, location int32, v0 int32, v1 int32) {
	C.ptrProgramUniform2i(gpProgramUniform2i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
func ProgramUniform2iv(program uint32, location int32, count int32, value *int32) {
	C.ptrProgramUniform2iv(gpProgramUniform2iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform2ui(program uint32, location int32, v0 uint32, v1 uint32) {
	C.ptrProgramUniform2ui(gpProgramUniform2ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}
func ProgramUniform2uiv(program uint32, location int32, count int32, value *uint32) {
	C.ptrProgramUniform2uiv(gpProgramUniform2uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform3d(program uint32, location int32, v0 float64, v1 float64, v2 float64) {
	C.ptrProgramUniform3d(gpProgramUniform3d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2))
}
func ProgramUniform3dv(program uint32, location int32, count int32, value *float64) {
	C.ptrProgramUniform3dv(gpProgramUniform3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniform3f(program uint32, location int32, v0 float32, v1 float32, v2 float32) {
	C.ptrProgramUniform3f(gpProgramUniform3f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
func ProgramUniform3fv(program uint32, location int32, count int32, value *float32) {
	C.ptrProgramUniform3fv(gpProgramUniform3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform3i(program uint32, location int32, v0 int32, v1 int32, v2 int32) {
	C.ptrProgramUniform3i(gpProgramUniform3i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
func ProgramUniform3iv(program uint32, location int32, count int32, value *int32) {
	C.ptrProgramUniform3iv(gpProgramUniform3iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform3ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.ptrProgramUniform3ui(gpProgramUniform3ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}
func ProgramUniform3uiv(program uint32, location int32, count int32, value *uint32) {
	C.ptrProgramUniform3uiv(gpProgramUniform3uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform4d(program uint32, location int32, v0 float64, v1 float64, v2 float64, v3 float64) {
	C.ptrProgramUniform4d(gpProgramUniform4d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLdouble)(v3))
}
func ProgramUniform4dv(program uint32, location int32, count int32, value *float64) {
	C.ptrProgramUniform4dv(gpProgramUniform4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniform4f(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.ptrProgramUniform4f(gpProgramUniform4f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
func ProgramUniform4fv(program uint32, location int32, count int32, value *float32) {
	C.ptrProgramUniform4fv(gpProgramUniform4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform4i(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.ptrProgramUniform4i(gpProgramUniform4i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
func ProgramUniform4iv(program uint32, location int32, count int32, value *int32) {
	C.ptrProgramUniform4iv(gpProgramUniform4iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform4ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.ptrProgramUniform4ui(gpProgramUniform4ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}
func ProgramUniform4uiv(program uint32, location int32, count int32, value *uint32) {
	C.ptrProgramUniform4uiv(gpProgramUniform4uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniformHandleui64ARB(program uint32, location int32, value uint64) {
	C.ptrProgramUniformHandleui64ARB(gpProgramUniformHandleui64ARB, (C.GLuint)(program), (C.GLint)(location), (C.GLuint64)(value))
}
func ProgramUniformHandleui64vARB(program uint32, location int32, count int32, values *uint64) {
	C.ptrProgramUniformHandleui64vARB(gpProgramUniformHandleui64vARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(values)))
}
func ProgramUniformMatrix2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.ptrProgramUniformMatrix2dv(gpProgramUniformMatrix2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.ptrProgramUniformMatrix2fv(gpProgramUniformMatrix2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.ptrProgramUniformMatrix2x3dv(gpProgramUniformMatrix2x3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.ptrProgramUniformMatrix2x3fv(gpProgramUniformMatrix2x3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.ptrProgramUniformMatrix2x4dv(gpProgramUniformMatrix2x4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.ptrProgramUniformMatrix2x4fv(gpProgramUniformMatrix2x4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.ptrProgramUniformMatrix3dv(gpProgramUniformMatrix3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.ptrProgramUniformMatrix3fv(gpProgramUniformMatrix3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.ptrProgramUniformMatrix3x2dv(gpProgramUniformMatrix3x2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.ptrProgramUniformMatrix3x2fv(gpProgramUniformMatrix3x2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.ptrProgramUniformMatrix3x4dv(gpProgramUniformMatrix3x4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.ptrProgramUniformMatrix3x4fv(gpProgramUniformMatrix3x4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.ptrProgramUniformMatrix4dv(gpProgramUniformMatrix4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.ptrProgramUniformMatrix4fv(gpProgramUniformMatrix4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.ptrProgramUniformMatrix4x2dv(gpProgramUniformMatrix4x2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.ptrProgramUniformMatrix4x2fv(gpProgramUniformMatrix4x2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.ptrProgramUniformMatrix4x3dv(gpProgramUniformMatrix4x3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.ptrProgramUniformMatrix4x3fv(gpProgramUniformMatrix4x3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProvokingVertex(mode int32) {
	C.ptrProvokingVertex(gpProvokingVertex, (C.GLenum)(mode))
}
func PushDebugGroup(source int32, id uint32, length int32, message string) {
	cstrmessage, freemessage := stringToChar(message)
	defer freemessage()
	C.ptrPushDebugGroup(gpPushDebugGroup, (C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(cstrmessage)))
}
func PushDebugGroupKHR(source int32, id uint32, length int32, message string) {
	cstrmessage, freemessage := stringToChar(message)
	defer freemessage()
	C.ptrPushDebugGroupKHR(gpPushDebugGroupKHR, (C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(cstrmessage)))
}
func QueryCounter(id uint32, target int32) {
	C.ptrQueryCounter(gpQueryCounter, (C.GLuint)(id), (C.GLenum)(target))
}
func ReadBuffer(src int32) {
	C.ptrReadBuffer(gpReadBuffer, (C.GLenum)(src))
}
func ReadPixels(x int32, y int32, width int32, height int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrReadPixels(gpReadPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func ReadnPixels(x int32, y int32, width int32, height int32, format int32, gltype int32, bufSize int32, data interface{}) {
	C.ptrReadnPixels(gpReadnPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(gltype), (C.GLsizei)(bufSize), ptr(data))
}
func ReadnPixelsARB(x int32, y int32, width int32, height int32, format int32, gltype int32, bufSize int32, data interface{}) {
	C.ptrReadnPixelsARB(gpReadnPixelsARB, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(gltype), (C.GLsizei)(bufSize), ptr(data))
}
func ReadnPixelsKHR(x int32, y int32, width int32, height int32, format int32, gltype int32, bufSize int32, data interface{}) {
	C.ptrReadnPixelsKHR(gpReadnPixelsKHR, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(gltype), (C.GLsizei)(bufSize), ptr(data))
}
func ReleaseShaderCompiler() {
	C.ptrReleaseShaderCompiler(gpReleaseShaderCompiler)
}
func RenderbufferStorage(target int32, internalformat int32, width int32, height int32) {
	C.ptrRenderbufferStorage(gpRenderbufferStorage, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func RenderbufferStorageMultisample(target int32, samples int32, internalformat int32, width int32, height int32) {
	C.ptrRenderbufferStorageMultisample(gpRenderbufferStorageMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ResumeTransformFeedback() {
	C.ptrResumeTransformFeedback(gpResumeTransformFeedback)
}
func SampleCoverage(value float32, invert bool) {
	C.ptrSampleCoverage(gpSampleCoverage, (C.GLfloat)(value), (C.GLboolean)(boolToInt(invert)))
}
func SampleMaski(maskNumber uint32, mask uint32) {
	C.ptrSampleMaski(gpSampleMaski, (C.GLuint)(maskNumber), (C.GLbitfield)(mask))
}
func SamplerParameterIiv(sampler uint32, pname int32, param *int32) {
	C.ptrSamplerParameterIiv(gpSamplerParameterIiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func SamplerParameterIuiv(sampler uint32, pname int32, param *uint32) {
	C.ptrSamplerParameterIuiv(gpSamplerParameterIuiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(param)))
}
func SamplerParameterf(sampler uint32, pname int32, param float32) {
	C.ptrSamplerParameterf(gpSamplerParameterf, (C.GLuint)(sampler), (C.GLenum)(pname), (C.GLfloat)(param))
}
func SamplerParameterfv(sampler uint32, pname int32, param *float32) {
	C.ptrSamplerParameterfv(gpSamplerParameterfv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(param)))
}
func SamplerParameteri(sampler uint32, pname int32, param int32) {
	C.ptrSamplerParameteri(gpSamplerParameteri, (C.GLuint)(sampler), (C.GLenum)(pname), (C.GLint)(param))
}
func SamplerParameteriv(sampler uint32, pname int32, param *int32) {
	C.ptrSamplerParameteriv(gpSamplerParameteriv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func Scissor(x int32, y int32, width int32, height int32) {
	C.ptrScissor(gpScissor, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ScissorArrayv(first uint32, count int32, v *int32) {
	C.ptrScissorArrayv(gpScissorArrayv, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(v)))
}
func ScissorIndexed(index uint32, left int32, bottom int32, width int32, height int32) {
	C.ptrScissorIndexed(gpScissorIndexed, (C.GLuint)(index), (C.GLint)(left), (C.GLint)(bottom), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ScissorIndexedv(index uint32, v *int32) {
	C.ptrScissorIndexedv(gpScissorIndexedv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func ShaderBinary(count int32, shaders *uint32, binaryformat int32, binary unsafe.Pointer, length int32) {
	C.ptrShaderBinary(gpShaderBinary, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(shaders)), (C.GLenum)(binaryformat), binary, (C.GLsizei)(length))
}
func ShaderSource(shader uint32, count int32, glstring string, length *int32) {
	code := glstring + "\x00"
	cstrglstring, freeglstring := stringToChar(code)
	defer freeglstring()
	C.ptrShaderSource(gpShaderSource, (C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(cstrglstring)), (*C.GLint)(unsafe.Pointer(length)))
}
func ShaderStorageBlockBinding(program uint32, storageBlockIndex uint32, storageBlockBinding uint32) {
	C.ptrShaderStorageBlockBinding(gpShaderStorageBlockBinding, (C.GLuint)(program), (C.GLuint)(storageBlockIndex), (C.GLuint)(storageBlockBinding))
}
func StencilFunc(glfunc int32, ref int32, mask uint32) {
	C.ptrStencilFunc(gpStencilFunc, (C.GLenum)(glfunc), (C.GLint)(ref), (C.GLuint)(mask))
}
func StencilFuncSeparate(face int32, glfunc int32, ref int32, mask uint32) {
	C.ptrStencilFuncSeparate(gpStencilFuncSeparate, (C.GLenum)(face), (C.GLenum)(glfunc), (C.GLint)(ref), (C.GLuint)(mask))
}
func StencilMask(mask uint32) {
	C.ptrStencilMask(gpStencilMask, (C.GLuint)(mask))
}
func StencilMaskSeparate(face int32, mask uint32) {
	C.ptrStencilMaskSeparate(gpStencilMaskSeparate, (C.GLenum)(face), (C.GLuint)(mask))
}
func StencilOp(fail int32, zfail int32, zpass int32) {
	C.ptrStencilOp(gpStencilOp, (C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}
func StencilOpSeparate(face int32, sfail int32, dpfail int32, dppass int32) {
	C.ptrStencilOpSeparate(gpStencilOpSeparate, (C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}
func TexBuffer(target int32, internalformat int32, buffer uint32) {
	C.ptrTexBuffer(gpTexBuffer, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}
func TexBufferRange(target int32, internalformat int32, buffer uint32, offset int64, size int64) {
	C.ptrTexBufferRange(gpTexBufferRange, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
func TexImage1D(target int32, level int32, internalformat int32, width int32, border int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrTexImage1D(gpTexImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func TexImage2D(target int32, level int32, internalformat int32, width int32, height int32, border int32, format int32, gltype int32, pixels interface{}) {
	C.ptrTexImage2D(gpTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(gltype), ptr(pixels))
}
func TexImage2DMultisample(target int32, samples int32, internalformat int32, width int32, height int32, fixedsamplelocations bool) {
	C.ptrTexImage2DMultisample(gpTexImage2DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}
func TexImage3D(target int32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrTexImage3D(gpTexImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func TexImage3DMultisample(target int32, samples int32, internalformat int32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.ptrTexImage3DMultisample(gpTexImage3DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}
func TexPageCommitmentARB(target int32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, resident bool) {
	C.ptrTexPageCommitmentARB(gpTexPageCommitmentARB, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(resident)))
}
func TexParameterIiv(target int32, pname int32, params *int32) {
	C.ptrTexParameterIiv(gpTexParameterIiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TexParameterIuiv(target int32, pname int32, params *uint32) {
	C.ptrTexParameterIuiv(gpTexParameterIuiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func TexParameterf(target int32, pname int32, param float32) {
	C.ptrTexParameterf(gpTexParameterf, (C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TexParameterfv(target int32, pname int32, params *float32) {
	C.ptrTexParameterfv(gpTexParameterfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func TexParameteri(target int32, pname int32, param int32) {
	C.ptrTexParameteri(gpTexParameteri, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func TexParameteriv(target int32, pname int32, params *int32) {
	C.ptrTexParameteriv(gpTexParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TexStorage1D(target int32, levels int32, internalformat int32, width int32) {
	C.ptrTexStorage1D(gpTexStorage1D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}
func TexStorage2D(target int32, levels int32, internalformat int32, width int32, height int32) {
	C.ptrTexStorage2D(gpTexStorage2D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func TexStorage2DMultisample(target int32, samples int32, internalformat int32, width int32, height int32, fixedsamplelocations bool) {
	C.ptrTexStorage2DMultisample(gpTexStorage2DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}
func TexStorage3D(target int32, levels int32, internalformat int32, width int32, height int32, depth int32) {
	C.ptrTexStorage3D(gpTexStorage3D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
func TexStorage3DMultisample(target int32, samples int32, internalformat int32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.ptrTexStorage3DMultisample(gpTexStorage3DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}
func TexSubImage1D(target int32, level int32, xoffset int32, width int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrTexSubImage1D(gpTexSubImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func TexSubImage2D(target int32, level int32, xoffset int32, yoffset int32, width int32, height int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrTexSubImage2D(gpTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func TexSubImage3D(target int32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrTexSubImage3D(gpTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func TextureBarrier() {
	C.ptrTextureBarrier(gpTextureBarrier)
}
func TextureBuffer(texture uint32, internalformat int32, buffer uint32) {
	C.ptrTextureBuffer(gpTextureBuffer, (C.GLuint)(texture), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}
func TextureBufferRange(texture uint32, internalformat int32, buffer uint32, offset int64, size int32) {
	C.ptrTextureBufferRange(gpTextureBufferRange, (C.GLuint)(texture), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size))
}
func TextureParameterIiv(texture uint32, pname int32, params *int32) {
	C.ptrTextureParameterIiv(gpTextureParameterIiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TextureParameterIuiv(texture uint32, pname int32, params *uint32) {
	C.ptrTextureParameterIuiv(gpTextureParameterIuiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func TextureParameterf(texture uint32, pname int32, param float32) {
	C.ptrTextureParameterf(gpTextureParameterf, (C.GLuint)(texture), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TextureParameterfv(texture uint32, pname int32, param *float32) {
	C.ptrTextureParameterfv(gpTextureParameterfv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(param)))
}
func TextureParameteri(texture uint32, pname int32, param int32) {
	C.ptrTextureParameteri(gpTextureParameteri, (C.GLuint)(texture), (C.GLenum)(pname), (C.GLint)(param))
}
func TextureParameteriv(texture uint32, pname int32, param *int32) {
	C.ptrTextureParameteriv(gpTextureParameteriv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func TextureStorage1D(texture uint32, levels int32, internalformat int32, width int32) {
	C.ptrTextureStorage1D(gpTextureStorage1D, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}
func TextureStorage2D(texture uint32, levels int32, internalformat int32, width int32, height int32) {
	C.ptrTextureStorage2D(gpTextureStorage2D, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func TextureStorage2DMultisample(texture uint32, samples int32, internalformat int32, width int32, height int32, fixedsamplelocations bool) {
	C.ptrTextureStorage2DMultisample(gpTextureStorage2DMultisample, (C.GLuint)(texture), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}
func TextureStorage3D(texture uint32, levels int32, internalformat int32, width int32, height int32, depth int32) {
	C.ptrTextureStorage3D(gpTextureStorage3D, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}
func TextureStorage3DMultisample(texture uint32, samples int32, internalformat int32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.ptrTextureStorage3DMultisample(gpTextureStorage3DMultisample, (C.GLuint)(texture), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}
func TextureSubImage1D(texture uint32, level int32, xoffset int32, width int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrTextureSubImage1D(gpTextureSubImage1D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func TextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrTextureSubImage2D(gpTextureSubImage2D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func TextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format int32, gltype int32, pixels unsafe.Pointer) {
	C.ptrTextureSubImage3D(gpTextureSubImage3D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(gltype), pixels)
}
func TextureView(texture uint32, target int32, origtexture uint32, internalformat int32, minlevel uint32, numlevels uint32, minlayer uint32, numlayers uint32) {
	C.ptrTextureView(gpTextureView, (C.GLuint)(texture), (C.GLenum)(target), (C.GLuint)(origtexture), (C.GLenum)(internalformat), (C.GLuint)(minlevel), (C.GLuint)(numlevels), (C.GLuint)(minlayer), (C.GLuint)(numlayers))
}
func TransformFeedbackBufferBase(xfb uint32, index uint32, buffer uint32) {
	C.ptrTransformFeedbackBufferBase(gpTransformFeedbackBufferBase, (C.GLuint)(xfb), (C.GLuint)(index), (C.GLuint)(buffer))
}
func TransformFeedbackBufferRange(xfb uint32, index uint32, buffer uint32, offset int64, size int32) {
	C.ptrTransformFeedbackBufferRange(gpTransformFeedbackBufferRange, (C.GLuint)(xfb), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size))
}
func TransformFeedbackVaryings(program uint32, count int32, varyings string, bufferMode int32) {
	cstrvaryings, freevaryings := stringToChar(varyings)
	defer freevaryings()
	C.ptrTransformFeedbackVaryings(gpTransformFeedbackVaryings, (C.GLuint)(program), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(cstrvaryings)), (C.GLenum)(bufferMode))
}
func Uniform1d(location int32, x float64) {
	C.ptrUniform1d(gpUniform1d, (C.GLint)(location), (C.GLdouble)(x))
}
func Uniform1dv(location int32, count int32, value *float64) {
	C.ptrUniform1dv(gpUniform1dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}
func Uniform1f(location int32, v0 float32) {
	C.ptrUniform1f(gpUniform1f, (C.GLint)(location), (C.GLfloat)(v0))
}
func Uniform1fv(location int32, count int32, value *float32) {
	C.ptrUniform1fv(gpUniform1fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func Uniform1i(location int32, v0 int32) {
	C.ptrUniform1i(gpUniform1i, (C.GLint)(location), (C.GLint)(v0))
}
func Uniform1iv(location int32, count int32, value *int32) {
	C.ptrUniform1iv(gpUniform1iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func Uniform1ui(location int32, v0 uint32) {
	C.ptrUniform1ui(gpUniform1ui, (C.GLint)(location), (C.GLuint)(v0))
}
func Uniform1uiv(location int32, count int32, value *uint32) {
	C.ptrUniform1uiv(gpUniform1uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func Uniform2d(location int32, x float64, y float64) {
	C.ptrUniform2d(gpUniform2d, (C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y))
}
func Uniform2dv(location int32, count int32, value *float64) {
	C.ptrUniform2dv(gpUniform2dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}
func Uniform2f(location int32, v0 float32, v1 float32) {
	C.ptrUniform2f(gpUniform2f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
func Uniform2fv(location int32, count int32, value *float32) {
	C.ptrUniform2fv(gpUniform2fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func Uniform2i(location int32, v0 int32, v1 int32) {
	C.ptrUniform2i(gpUniform2i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
func Uniform2iv(location int32, count int32, value *int32) {
	C.ptrUniform2iv(gpUniform2iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func Uniform2ui(location int32, v0 uint32, v1 uint32) {
	C.ptrUniform2ui(gpUniform2ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}
func Uniform2uiv(location int32, count int32, value *uint32) {
	C.ptrUniform2uiv(gpUniform2uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func Uniform3d(location int32, x float64, y float64, z float64) {
	C.ptrUniform3d(gpUniform3d, (C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Uniform3dv(location int32, count int32, value *float64) {
	C.ptrUniform3dv(gpUniform3dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}
func Uniform3f(location int32, v0 float32, v1 float32, v2 float32) {
	C.ptrUniform3f(gpUniform3f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
func Uniform3fv(location int32, count int32, value *float32) {
	C.ptrUniform3fv(gpUniform3fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func Uniform3i(location int32, v0 int32, v1 int32, v2 int32) {
	C.ptrUniform3i(gpUniform3i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
func Uniform3iv(location int32, count int32, value *int32) {
	C.ptrUniform3iv(gpUniform3iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func Uniform3ui(location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.ptrUniform3ui(gpUniform3ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}
func Uniform3uiv(location int32, count int32, value *uint32) {
	C.ptrUniform3uiv(gpUniform3uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func Uniform4d(location int32, x float64, y float64, z float64, w float64) {
	C.ptrUniform4d(gpUniform4d, (C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func Uniform4dv(location int32, count int32, value *float64) {
	C.ptrUniform4dv(gpUniform4dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}
func Uniform4f(location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.ptrUniform4f(gpUniform4f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
func Uniform4fv(location int32, count int32, value *float32) {
	C.ptrUniform4fv(gpUniform4fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func Uniform4i(location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.ptrUniform4i(gpUniform4i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
func Uniform4iv(location int32, count int32, value *int32) {
	C.ptrUniform4iv(gpUniform4iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func Uniform4ui(location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.ptrUniform4ui(gpUniform4ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}
func Uniform4uiv(location int32, count int32, value *uint32) {
	C.ptrUniform4uiv(gpUniform4uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func UniformBlockBinding(program uint32, uniformBlockIndex uint32, uniformBlockBinding uint32) {
	C.ptrUniformBlockBinding(gpUniformBlockBinding, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLuint)(uniformBlockBinding))
}
func UniformHandleui64ARB(location int32, value uint64) {
	C.ptrUniformHandleui64ARB(gpUniformHandleui64ARB, (C.GLint)(location), (C.GLuint64)(value))
}
func UniformHandleui64vARB(location int32, count int32, value *uint64) {
	C.ptrUniformHandleui64vARB(gpUniformHandleui64vARB, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(value)))
}
func UniformMatrix2dv(location int32, count int32, transpose bool, value *float64) {
	C.ptrUniformMatrix2dv(gpUniformMatrix2dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func UniformMatrix2fv(location int32, count int32, transpose bool, value *float32) {
	C.ptrUniformMatrix2fv(gpUniformMatrix2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix2x3dv(location int32, count int32, transpose bool, value *float64) {
	C.ptrUniformMatrix2x3dv(gpUniformMatrix2x3dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func UniformMatrix2x3fv(location int32, count int32, transpose bool, value *float32) {
	C.ptrUniformMatrix2x3fv(gpUniformMatrix2x3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix2x4dv(location int32, count int32, transpose bool, value *float64) {
	C.ptrUniformMatrix2x4dv(gpUniformMatrix2x4dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func UniformMatrix2x4fv(location int32, count int32, transpose bool, value *float32) {
	C.ptrUniformMatrix2x4fv(gpUniformMatrix2x4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3dv(location int32, count int32, transpose bool, value *float64) {
	C.ptrUniformMatrix3dv(gpUniformMatrix3dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func UniformMatrix3fv(location int32, count int32, transpose bool, value *float32) {
	C.ptrUniformMatrix3fv(gpUniformMatrix3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3x2dv(location int32, count int32, transpose bool, value *float64) {
	C.ptrUniformMatrix3x2dv(gpUniformMatrix3x2dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func UniformMatrix3x2fv(location int32, count int32, transpose bool, value *float32) {
	C.ptrUniformMatrix3x2fv(gpUniformMatrix3x2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3x4dv(location int32, count int32, transpose bool, value *float64) {
	C.ptrUniformMatrix3x4dv(gpUniformMatrix3x4dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func UniformMatrix3x4fv(location int32, count int32, transpose bool, value *float32) {
	C.ptrUniformMatrix3x4fv(gpUniformMatrix3x4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4dv(location int32, count int32, transpose bool, value *float64) {
	C.ptrUniformMatrix4dv(gpUniformMatrix4dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func UniformMatrix4fv(location int32, count int32, transpose bool, value *float32) {
	C.ptrUniformMatrix4fv(gpUniformMatrix4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4x2dv(location int32, count int32, transpose bool, value *float64) {
	C.ptrUniformMatrix4x2dv(gpUniformMatrix4x2dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func UniformMatrix4x2fv(location int32, count int32, transpose bool, value *float32) {
	C.ptrUniformMatrix4x2fv(gpUniformMatrix4x2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4x3dv(location int32, count int32, transpose bool, value *float64) {
	C.ptrUniformMatrix4x3dv(gpUniformMatrix4x3dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}
func UniformMatrix4x3fv(location int32, count int32, transpose bool, value *float32) {
	C.ptrUniformMatrix4x3fv(gpUniformMatrix4x3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformSubroutinesuiv(shadertype int32, count int32, indices *uint32) {
	C.ptrUniformSubroutinesuiv(gpUniformSubroutinesuiv, (C.GLenum)(shadertype), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(indices)))
}
func UnmapBuffer(target int32) bool {
	ret := C.ptrUnmapBuffer(gpUnmapBuffer, (C.GLenum)(target))
	return intToBool((int8)(ret))
}
func UnmapNamedBuffer(buffer uint32) bool {
	ret := C.ptrUnmapNamedBuffer(gpUnmapNamedBuffer, (C.GLuint)(buffer))
	return intToBool((int8)(ret))
}
func UseProgram(program uint32) {
	C.ptrUseProgram(gpUseProgram, (C.GLuint)(program))
}
func UseProgramStages(pipeline uint32, stages uint32, program uint32) {
	C.ptrUseProgramStages(gpUseProgramStages, (C.GLuint)(pipeline), (C.GLbitfield)(stages), (C.GLuint)(program))
}
func ValidateProgram(program uint32) {
	C.ptrValidateProgram(gpValidateProgram, (C.GLuint)(program))
}
func ValidateProgramPipeline(pipeline uint32) {
	C.ptrValidateProgramPipeline(gpValidateProgramPipeline, (C.GLuint)(pipeline))
}
func VertexArrayAttribBinding(vaobj uint32, attribindex uint32, bindingindex uint32) {
	C.ptrVertexArrayAttribBinding(gpVertexArrayAttribBinding, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLuint)(bindingindex))
}
func VertexArrayAttribFormat(vaobj uint32, attribindex uint32, size int32, gltype int32, normalized bool, relativeoffset uint32) {
	C.ptrVertexArrayAttribFormat(gpVertexArrayAttribFormat, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(relativeoffset))
}
func VertexArrayAttribIFormat(vaobj uint32, attribindex uint32, size int32, gltype int32, relativeoffset uint32) {
	C.ptrVertexArrayAttribIFormat(gpVertexArrayAttribIFormat, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(gltype), (C.GLuint)(relativeoffset))
}
func VertexArrayAttribLFormat(vaobj uint32, attribindex uint32, size int32, gltype int32, relativeoffset uint32) {
	C.ptrVertexArrayAttribLFormat(gpVertexArrayAttribLFormat, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(gltype), (C.GLuint)(relativeoffset))
}
func VertexArrayBindingDivisor(vaobj uint32, bindingindex uint32, divisor uint32) {
	C.ptrVertexArrayBindingDivisor(gpVertexArrayBindingDivisor, (C.GLuint)(vaobj), (C.GLuint)(bindingindex), (C.GLuint)(divisor))
}
func VertexArrayElementBuffer(vaobj uint32, buffer uint32) {
	C.ptrVertexArrayElementBuffer(gpVertexArrayElementBuffer, (C.GLuint)(vaobj), (C.GLuint)(buffer))
}
func VertexArrayVertexBuffer(vaobj uint32, bindingindex uint32, buffer uint32, offset int64, stride int32) {
	C.ptrVertexArrayVertexBuffer(gpVertexArrayVertexBuffer, (C.GLuint)(vaobj), (C.GLuint)(bindingindex), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(stride))
}
func VertexArrayVertexBuffers(vaobj uint32, first uint32, count int32, buffers *uint32, offsets *int64, strides *int32) {
	C.ptrVertexArrayVertexBuffers(gpVertexArrayVertexBuffers, (C.GLuint)(vaobj), (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)), (*C.GLintptr)(unsafe.Pointer(offsets)), (*C.GLsizei)(unsafe.Pointer(strides)))
}
func VertexAttrib1d(index uint32, x float64) {
	C.ptrVertexAttrib1d(gpVertexAttrib1d, (C.GLuint)(index), (C.GLdouble)(x))
}
func VertexAttrib1dv(index uint32, v *float64) {
	C.ptrVertexAttrib1dv(gpVertexAttrib1dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib1f(index uint32, x float32) {
	C.ptrVertexAttrib1f(gpVertexAttrib1f, (C.GLuint)(index), (C.GLfloat)(x))
}
func VertexAttrib1fv(index uint32, v *float32) {
	C.ptrVertexAttrib1fv(gpVertexAttrib1fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib1s(index uint32, x int16) {
	C.ptrVertexAttrib1s(gpVertexAttrib1s, (C.GLuint)(index), (C.GLshort)(x))
}
func VertexAttrib1sv(index uint32, v *int16) {
	C.ptrVertexAttrib1sv(gpVertexAttrib1sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib2d(index uint32, x float64, y float64) {
	C.ptrVertexAttrib2d(gpVertexAttrib2d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
func VertexAttrib2dv(index uint32, v *float64) {
	C.ptrVertexAttrib2dv(gpVertexAttrib2dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib2f(index uint32, x float32, y float32) {
	C.ptrVertexAttrib2f(gpVertexAttrib2f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
}
func VertexAttrib2fv(index uint32, v *float32) {
	C.ptrVertexAttrib2fv(gpVertexAttrib2fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib2s(index uint32, x int16, y int16) {
	C.ptrVertexAttrib2s(gpVertexAttrib2s, (C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y))
}
func VertexAttrib2sv(index uint32, v *int16) {
	C.ptrVertexAttrib2sv(gpVertexAttrib2sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib3d(index uint32, x float64, y float64, z float64) {
	C.ptrVertexAttrib3d(gpVertexAttrib3d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func VertexAttrib3dv(index uint32, v *float64) {
	C.ptrVertexAttrib3dv(gpVertexAttrib3dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib3f(index uint32, x float32, y float32, z float32) {
	C.ptrVertexAttrib3f(gpVertexAttrib3f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func VertexAttrib3fv(index uint32, v *float32) {
	C.ptrVertexAttrib3fv(gpVertexAttrib3fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib3s(index uint32, x int16, y int16, z int16) {
	C.ptrVertexAttrib3s(gpVertexAttrib3s, (C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func VertexAttrib3sv(index uint32, v *int16) {
	C.ptrVertexAttrib3sv(gpVertexAttrib3sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib4Nbv(index uint32, v *int8) {
	C.ptrVertexAttrib4Nbv(gpVertexAttrib4Nbv, (C.GLuint)(index), (*C.GLbyte)(unsafe.Pointer(v)))
}
func VertexAttrib4Niv(index uint32, v *int32) {
	C.ptrVertexAttrib4Niv(gpVertexAttrib4Niv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttrib4Nsv(index uint32, v *int16) {
	C.ptrVertexAttrib4Nsv(gpVertexAttrib4Nsv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib4Nub(index uint32, x uint8, y uint8, z uint8, w uint8) {
	C.ptrVertexAttrib4Nub(gpVertexAttrib4Nub, (C.GLuint)(index), (C.GLubyte)(x), (C.GLubyte)(y), (C.GLubyte)(z), (C.GLubyte)(w))
}
func VertexAttrib4Nubv(index uint32, v *uint8) {
	C.ptrVertexAttrib4Nubv(gpVertexAttrib4Nubv, (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(v)))
}
func VertexAttrib4Nuiv(index uint32, v *uint32) {
	C.ptrVertexAttrib4Nuiv(gpVertexAttrib4Nuiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttrib4Nusv(index uint32, v *uint16) {
	C.ptrVertexAttrib4Nusv(gpVertexAttrib4Nusv, (C.GLuint)(index), (*C.GLushort)(unsafe.Pointer(v)))
}
func VertexAttrib4bv(index uint32, v *int8) {
	C.ptrVertexAttrib4bv(gpVertexAttrib4bv, (C.GLuint)(index), (*C.GLbyte)(unsafe.Pointer(v)))
}
func VertexAttrib4d(index uint32, x float64, y float64, z float64, w float64) {
	C.ptrVertexAttrib4d(gpVertexAttrib4d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func VertexAttrib4dv(index uint32, v *float64) {
	C.ptrVertexAttrib4dv(gpVertexAttrib4dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib4f(index uint32, x float32, y float32, z float32, w float32) {
	C.ptrVertexAttrib4f(gpVertexAttrib4f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func VertexAttrib4fv(index uint32, v *float32) {
	C.ptrVertexAttrib4fv(gpVertexAttrib4fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib4iv(index uint32, v *int32) {
	C.ptrVertexAttrib4iv(gpVertexAttrib4iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttrib4s(index uint32, x int16, y int16, z int16, w int16) {
	C.ptrVertexAttrib4s(gpVertexAttrib4s, (C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func VertexAttrib4sv(index uint32, v *int16) {
	C.ptrVertexAttrib4sv(gpVertexAttrib4sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib4ubv(index uint32, v *uint8) {
	C.ptrVertexAttrib4ubv(gpVertexAttrib4ubv, (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(v)))
}
func VertexAttrib4uiv(index uint32, v *uint32) {
	C.ptrVertexAttrib4uiv(gpVertexAttrib4uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttrib4usv(index uint32, v *uint16) {
	C.ptrVertexAttrib4usv(gpVertexAttrib4usv, (C.GLuint)(index), (*C.GLushort)(unsafe.Pointer(v)))
}
func VertexAttribBinding(attribindex uint32, bindingindex uint32) {
	C.ptrVertexAttribBinding(gpVertexAttribBinding, (C.GLuint)(attribindex), (C.GLuint)(bindingindex))
}
func VertexAttribDivisor(index uint32, divisor uint32) {
	C.ptrVertexAttribDivisor(gpVertexAttribDivisor, (C.GLuint)(index), (C.GLuint)(divisor))
}
func VertexAttribFormat(attribindex uint32, size int32, gltype int32, normalized bool, relativeoffset uint32) {
	C.ptrVertexAttribFormat(gpVertexAttribFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(relativeoffset))
}
func VertexAttribI1i(index uint32, x int32) {
	C.ptrVertexAttribI1i(gpVertexAttribI1i, (C.GLuint)(index), (C.GLint)(x))
}
func VertexAttribI1iv(index uint32, v *int32) {
	C.ptrVertexAttribI1iv(gpVertexAttribI1iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI1ui(index uint32, x uint32) {
	C.ptrVertexAttribI1ui(gpVertexAttribI1ui, (C.GLuint)(index), (C.GLuint)(x))
}
func VertexAttribI1uiv(index uint32, v *uint32) {
	C.ptrVertexAttribI1uiv(gpVertexAttribI1uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI2i(index uint32, x int32, y int32) {
	C.ptrVertexAttribI2i(gpVertexAttribI2i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y))
}
func VertexAttribI2iv(index uint32, v *int32) {
	C.ptrVertexAttribI2iv(gpVertexAttribI2iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI2ui(index uint32, x uint32, y uint32) {
	C.ptrVertexAttribI2ui(gpVertexAttribI2ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y))
}
func VertexAttribI2uiv(index uint32, v *uint32) {
	C.ptrVertexAttribI2uiv(gpVertexAttribI2uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI3i(index uint32, x int32, y int32, z int32) {
	C.ptrVertexAttribI3i(gpVertexAttribI3i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func VertexAttribI3iv(index uint32, v *int32) {
	C.ptrVertexAttribI3iv(gpVertexAttribI3iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI3ui(index uint32, x uint32, y uint32, z uint32) {
	C.ptrVertexAttribI3ui(gpVertexAttribI3ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z))
}
func VertexAttribI3uiv(index uint32, v *uint32) {
	C.ptrVertexAttribI3uiv(gpVertexAttribI3uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI4bv(index uint32, v *int8) {
	C.ptrVertexAttribI4bv(gpVertexAttribI4bv, (C.GLuint)(index), (*C.GLbyte)(unsafe.Pointer(v)))
}
func VertexAttribI4i(index uint32, x int32, y int32, z int32, w int32) {
	C.ptrVertexAttribI4i(gpVertexAttribI4i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func VertexAttribI4iv(index uint32, v *int32) {
	C.ptrVertexAttribI4iv(gpVertexAttribI4iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI4sv(index uint32, v *int16) {
	C.ptrVertexAttribI4sv(gpVertexAttribI4sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttribI4ubv(index uint32, v *uint8) {
	C.ptrVertexAttribI4ubv(gpVertexAttribI4ubv, (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(v)))
}
func VertexAttribI4ui(index uint32, x uint32, y uint32, z uint32, w uint32) {
	C.ptrVertexAttribI4ui(gpVertexAttribI4ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
func VertexAttribI4uiv(index uint32, v *uint32) {
	C.ptrVertexAttribI4uiv(gpVertexAttribI4uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI4usv(index uint32, v *uint16) {
	C.ptrVertexAttribI4usv(gpVertexAttribI4usv, (C.GLuint)(index), (*C.GLushort)(unsafe.Pointer(v)))
}
func VertexAttribIFormat(attribindex uint32, size int32, gltype int32, relativeoffset uint32) {
	C.ptrVertexAttribIFormat(gpVertexAttribIFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(gltype), (C.GLuint)(relativeoffset))
}
func VertexAttribIPointer(index uint32, size int32, gltype int32, stride int32, pointer unsafe.Pointer) {
	C.ptrVertexAttribIPointer(gpVertexAttribIPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(gltype), (C.GLsizei)(stride), pointer)
}
func VertexAttribL1d(index uint32, x float64) {
	C.ptrVertexAttribL1d(gpVertexAttribL1d, (C.GLuint)(index), (C.GLdouble)(x))
}
func VertexAttribL1dv(index uint32, v *float64) {
	C.ptrVertexAttribL1dv(gpVertexAttribL1dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribL1ui64ARB(index uint32, x uint64) {
	C.ptrVertexAttribL1ui64ARB(gpVertexAttribL1ui64ARB, (C.GLuint)(index), (C.GLuint64EXT)(x))
}
func VertexAttribL1ui64vARB(index uint32, v *uint64) {
	C.ptrVertexAttribL1ui64vARB(gpVertexAttribL1ui64vARB, (C.GLuint)(index), (*C.GLuint64EXT)(unsafe.Pointer(v)))
}
func VertexAttribL2d(index uint32, x float64, y float64) {
	C.ptrVertexAttribL2d(gpVertexAttribL2d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
func VertexAttribL2dv(index uint32, v *float64) {
	C.ptrVertexAttribL2dv(gpVertexAttribL2dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribL3d(index uint32, x float64, y float64, z float64) {
	C.ptrVertexAttribL3d(gpVertexAttribL3d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func VertexAttribL3dv(index uint32, v *float64) {
	C.ptrVertexAttribL3dv(gpVertexAttribL3dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribL4d(index uint32, x float64, y float64, z float64, w float64) {
	C.ptrVertexAttribL4d(gpVertexAttribL4d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func VertexAttribL4dv(index uint32, v *float64) {
	C.ptrVertexAttribL4dv(gpVertexAttribL4dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribLFormat(attribindex uint32, size int32, gltype int32, relativeoffset uint32) {
	C.ptrVertexAttribLFormat(gpVertexAttribLFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(gltype), (C.GLuint)(relativeoffset))
}
func VertexAttribLPointer(index uint32, size int32, gltype int32, stride int32, pointer unsafe.Pointer) {
	C.ptrVertexAttribLPointer(gpVertexAttribLPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(gltype), (C.GLsizei)(stride), pointer)
}
func VertexAttribP1ui(index uint32, gltype int32, normalized bool, value uint32) {
	C.ptrVertexAttribP1ui(gpVertexAttribP1ui, (C.GLuint)(index), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP1uiv(index uint32, gltype int32, normalized bool, value *uint32) {
	C.ptrVertexAttribP1uiv(gpVertexAttribP1uiv, (C.GLuint)(index), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribP2ui(index uint32, gltype int32, normalized bool, value uint32) {
	C.ptrVertexAttribP2ui(gpVertexAttribP2ui, (C.GLuint)(index), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP2uiv(index uint32, gltype int32, normalized bool, value *uint32) {
	C.ptrVertexAttribP2uiv(gpVertexAttribP2uiv, (C.GLuint)(index), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribP3ui(index uint32, gltype int32, normalized bool, value uint32) {
	C.ptrVertexAttribP3ui(gpVertexAttribP3ui, (C.GLuint)(index), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP3uiv(index uint32, gltype int32, normalized bool, value *uint32) {
	C.ptrVertexAttribP3uiv(gpVertexAttribP3uiv, (C.GLuint)(index), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribP4ui(index uint32, gltype int32, normalized bool, value uint32) {
	C.ptrVertexAttribP4ui(gpVertexAttribP4ui, (C.GLuint)(index), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP4uiv(index uint32, gltype int32, normalized bool, value *uint32) {
	C.ptrVertexAttribP4uiv(gpVertexAttribP4uiv, (C.GLuint)(index), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribPointer(index uint32, size int32, gltype int32, normalized bool, stride int32, offset int32) {
	C.ptrVertexAttribPointer(gpVertexAttribPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(gltype), (C.GLboolean)(boolToInt(normalized)), (C.GLsizei)(stride), unsafe.Pointer(uintptr(offset)))
}
func VertexBindingDivisor(bindingindex uint32, divisor uint32) {
	C.ptrVertexBindingDivisor(gpVertexBindingDivisor, (C.GLuint)(bindingindex), (C.GLuint)(divisor))
}
func Viewport(x int32, y int32, width int32, height int32) {
	C.ptrViewport(gpViewport, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ViewportArrayv(first uint32, count int32, v *float32) {
	C.ptrViewportArrayv(gpViewportArrayv, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(v)))
}
func ViewportIndexedf(index uint32, x float32, y float32, w float32, h float32) {
	C.ptrViewportIndexedf(gpViewportIndexedf, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(w), (C.GLfloat)(h))
}
func ViewportIndexedfv(index uint32, v *float32) {
	C.ptrViewportIndexedfv(gpViewportIndexedfv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func WaitSync(sync unsafe.Pointer, flags uint32, timeout uint64) {
	C.ptrWaitSync(gpWaitSync, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}

// Init initializes the OpenGL bindings by loading the function pointers (for
// each OpenGL function) from the active OpenGL context.
//
// It must be called under the presence of an active OpenGL context, e.g.,
// always after calling window.MakeContextCurrent() and always before calling
// any OpenGL functions exported by this package.
//
// On Windows, Init loads pointers that are context-specific (and hence you
// must re-init if switching between OpenGL contexts, although not calling Init
// again after switching between OpenGL contexts may work if the contexts belong
// to the same graphics driver/device).
//
// On macOS and the other POSIX systems, the behavior is different, but code
// written compatible with the Windows behavior is compatible with macOS and the
// other POSIX systems. That is, always Init under an active OpenGL context, and
// always re-init after switching graphics contexts.
//
// For information about caveats of Init, you should read the "Platform Specific
// Function Retrieval" section of https://www.opengl.org/wiki/Load_OpenGL_Functions.
func Init() error {
	return InitWithProcAddrFunc(getProcAddress)
}

// InitWithProcAddrFunc intializes the package using the specified OpenGL
// function pointer loading function. For more cases Init should be used
// instead.
func InitWithProcAddrFunc(getProcAddr func(name string) unsafe.Pointer) error {
	missList := make([]string, 0)
	gpActiveShaderProgram = (C.GPACTIVESHADERPROGRAM)(getProcAddr("glActiveShaderProgram"))
	if gpActiveShaderProgram == nil {
		return errors.New("glActiveShaderProgram")
	}
	gpActiveTexture = (C.GPACTIVETEXTURE)(getProcAddr("glActiveTexture"))
	if gpActiveTexture == nil {
		return errors.New("glActiveTexture")
	}
	gpAttachShader = (C.GPATTACHSHADER)(getProcAddr("glAttachShader"))
	if gpAttachShader == nil {
		return errors.New("glAttachShader")
	}
	gpBeginConditionalRender = (C.GPBEGINCONDITIONALRENDER)(getProcAddr("glBeginConditionalRender"))
	if gpBeginConditionalRender == nil {
		return errors.New("glBeginConditionalRender")
	}
	gpBeginQuery = (C.GPBEGINQUERY)(getProcAddr("glBeginQuery"))
	if gpBeginQuery == nil {
		return errors.New("glBeginQuery")
	}
	gpBeginQueryIndexed = (C.GPBEGINQUERYINDEXED)(getProcAddr("glBeginQueryIndexed"))
	if gpBeginQueryIndexed == nil {
		return errors.New("glBeginQueryIndexed")
	}
	gpBeginTransformFeedback = (C.GPBEGINTRANSFORMFEEDBACK)(getProcAddr("glBeginTransformFeedback"))
	if gpBeginTransformFeedback == nil {
		return errors.New("glBeginTransformFeedback")
	}
	gpBindAttribLocation = (C.GPBINDATTRIBLOCATION)(getProcAddr("glBindAttribLocation"))
	if gpBindAttribLocation == nil {
		return errors.New("glBindAttribLocation")
	}
	gpBindBuffer = (C.GPBINDBUFFER)(getProcAddr("glBindBuffer"))
	if gpBindBuffer == nil {
		return errors.New("glBindBuffer")
	}
	gpBindBufferBase = (C.GPBINDBUFFERBASE)(getProcAddr("glBindBufferBase"))
	if gpBindBufferBase == nil {
		return errors.New("glBindBufferBase")
	}
	gpBindBufferRange = (C.GPBINDBUFFERRANGE)(getProcAddr("glBindBufferRange"))
	if gpBindBufferRange == nil {
		return errors.New("glBindBufferRange")
	}
	gpBindBuffersBase = (C.GPBINDBUFFERSBASE)(getProcAddr("glBindBuffersBase"))
	if gpBindBuffersBase == nil {
		return errors.New("glBindBuffersBase")
	}
	gpBindBuffersRange = (C.GPBINDBUFFERSRANGE)(getProcAddr("glBindBuffersRange"))
	if gpBindBuffersRange == nil {
		return errors.New("glBindBuffersRange")
	}
	gpBindFragDataLocation = (C.GPBINDFRAGDATALOCATION)(getProcAddr("glBindFragDataLocation"))
	if gpBindFragDataLocation == nil {
		return errors.New("glBindFragDataLocation")
	}
	gpBindFragDataLocationIndexed = (C.GPBINDFRAGDATALOCATIONINDEXED)(getProcAddr("glBindFragDataLocationIndexed"))
	if gpBindFragDataLocationIndexed == nil {
		return errors.New("glBindFragDataLocationIndexed")
	}
	gpBindFramebuffer = (C.GPBINDFRAMEBUFFER)(getProcAddr("glBindFramebuffer"))
	if gpBindFramebuffer == nil {
		return errors.New("glBindFramebuffer")
	}
	gpBindImageTexture = (C.GPBINDIMAGETEXTURE)(getProcAddr("glBindImageTexture"))
	if gpBindImageTexture == nil {
		return errors.New("glBindImageTexture")
	}
	gpBindImageTextures = (C.GPBINDIMAGETEXTURES)(getProcAddr("glBindImageTextures"))
	if gpBindImageTextures == nil {
		return errors.New("glBindImageTextures")
	}
	gpBindProgramPipeline = (C.GPBINDPROGRAMPIPELINE)(getProcAddr("glBindProgramPipeline"))
	if gpBindProgramPipeline == nil {
		return errors.New("glBindProgramPipeline")
	}
	gpBindRenderbuffer = (C.GPBINDRENDERBUFFER)(getProcAddr("glBindRenderbuffer"))
	if gpBindRenderbuffer == nil {
		return errors.New("glBindRenderbuffer")
	}
	gpBindSampler = (C.GPBINDSAMPLER)(getProcAddr("glBindSampler"))
	if gpBindSampler == nil {
		return errors.New("glBindSampler")
	}
	gpBindSamplers = (C.GPBINDSAMPLERS)(getProcAddr("glBindSamplers"))
	if gpBindSamplers == nil {
		return errors.New("glBindSamplers")
	}
	gpBindTexture = (C.GPBINDTEXTURE)(getProcAddr("glBindTexture"))
	if gpBindTexture == nil {
		return errors.New("glBindTexture")
	}
	gpBindTextureUnit = (C.GPBINDTEXTUREUNIT)(getProcAddr("glBindTextureUnit"))
	if gpBindTextureUnit == nil {
		return errors.New("glBindTextureUnit")
	}
	gpBindTextures = (C.GPBINDTEXTURES)(getProcAddr("glBindTextures"))
	if gpBindTextures == nil {
		return errors.New("glBindTextures")
	}
	gpBindTransformFeedback = (C.GPBINDTRANSFORMFEEDBACK)(getProcAddr("glBindTransformFeedback"))
	if gpBindTransformFeedback == nil {
		return errors.New("glBindTransformFeedback")
	}
	gpBindVertexArray = (C.GPBINDVERTEXARRAY)(getProcAddr("glBindVertexArray"))
	if gpBindVertexArray == nil {
		return errors.New("glBindVertexArray")
	}
	gpBindVertexBuffer = (C.GPBINDVERTEXBUFFER)(getProcAddr("glBindVertexBuffer"))
	if gpBindVertexBuffer == nil {
		return errors.New("glBindVertexBuffer")
	}
	gpBindVertexBuffers = (C.GPBINDVERTEXBUFFERS)(getProcAddr("glBindVertexBuffers"))
	if gpBindVertexBuffers == nil {
		return errors.New("glBindVertexBuffers")
	}
	gpBlendColor = (C.GPBLENDCOLOR)(getProcAddr("glBlendColor"))
	if gpBlendColor == nil {
		return errors.New("glBlendColor")
	}
	gpBlendEquation = (C.GPBLENDEQUATION)(getProcAddr("glBlendEquation"))
	if gpBlendEquation == nil {
		return errors.New("glBlendEquation")
	}
	gpBlendEquationSeparate = (C.GPBLENDEQUATIONSEPARATE)(getProcAddr("glBlendEquationSeparate"))
	if gpBlendEquationSeparate == nil {
		return errors.New("glBlendEquationSeparate")
	}
	gpBlendEquationSeparatei = (C.GPBLENDEQUATIONSEPARATEI)(getProcAddr("glBlendEquationSeparatei"))
	if gpBlendEquationSeparatei == nil {
		return errors.New("glBlendEquationSeparatei")
	}
	gpBlendEquationSeparateiARB = (C.GPBLENDEQUATIONSEPARATEIARB)(getProcAddr("glBlendEquationSeparateiARB"))
	gpBlendEquationi = (C.GPBLENDEQUATIONI)(getProcAddr("glBlendEquationi"))
	if gpBlendEquationi == nil {
		return errors.New("glBlendEquationi")
	}
	gpBlendEquationiARB = (C.GPBLENDEQUATIONIARB)(getProcAddr("glBlendEquationiARB"))
	gpBlendFunc = (C.GPBLENDFUNC)(getProcAddr("glBlendFunc"))
	if gpBlendFunc == nil {
		return errors.New("glBlendFunc")
	}
	gpBlendFuncSeparate = (C.GPBLENDFUNCSEPARATE)(getProcAddr("glBlendFuncSeparate"))
	if gpBlendFuncSeparate == nil {
		return errors.New("glBlendFuncSeparate")
	}
	gpBlendFuncSeparatei = (C.GPBLENDFUNCSEPARATEI)(getProcAddr("glBlendFuncSeparatei"))
	if gpBlendFuncSeparatei == nil {
		return errors.New("glBlendFuncSeparatei")
	}
	gpBlendFuncSeparateiARB = (C.GPBLENDFUNCSEPARATEIARB)(getProcAddr("glBlendFuncSeparateiARB"))
	gpBlendFunci = (C.GPBLENDFUNCI)(getProcAddr("glBlendFunci"))
	if gpBlendFunci == nil {
		return errors.New("glBlendFunci")
	}
	gpBlendFunciARB = (C.GPBLENDFUNCIARB)(getProcAddr("glBlendFunciARB"))
	gpBlitFramebuffer = (C.GPBLITFRAMEBUFFER)(getProcAddr("glBlitFramebuffer"))
	if gpBlitFramebuffer == nil {
		return errors.New("glBlitFramebuffer")
	}
	gpBlitNamedFramebuffer = (C.GPBLITNAMEDFRAMEBUFFER)(getProcAddr("glBlitNamedFramebuffer"))
	if gpBlitNamedFramebuffer == nil {
		return errors.New("glBlitNamedFramebuffer")
	}
	gpBufferData = (C.GPBUFFERDATA)(getProcAddr("glBufferData"))
	if gpBufferData == nil {
		return errors.New("glBufferData")
	}
	gpBufferPageCommitmentARB = (C.GPBUFFERPAGECOMMITMENTARB)(getProcAddr("glBufferPageCommitmentARB"))
	gpBufferStorage = (C.GPBUFFERSTORAGE)(getProcAddr("glBufferStorage"))
	if gpBufferStorage == nil {
		return errors.New("glBufferStorage")
	}
	gpBufferSubData = (C.GPBUFFERSUBDATA)(getProcAddr("glBufferSubData"))
	if gpBufferSubData == nil {
		return errors.New("glBufferSubData")
	}
	gpCheckFramebufferStatus = (C.GPCHECKFRAMEBUFFERSTATUS)(getProcAddr("glCheckFramebufferStatus"))
	if gpCheckFramebufferStatus == nil {
		return errors.New("glCheckFramebufferStatus")
	}
	gpCheckNamedFramebufferStatus = (C.GPCHECKNAMEDFRAMEBUFFERSTATUS)(getProcAddr("glCheckNamedFramebufferStatus"))
	if gpCheckNamedFramebufferStatus == nil {
		return errors.New("glCheckNamedFramebufferStatus")
	}
	gpClampColor = (C.GPCLAMPCOLOR)(getProcAddr("glClampColor"))
	if gpClampColor == nil {
		return errors.New("glClampColor")
	}
	gpClear = (C.GPCLEAR)(getProcAddr("glClear"))
	if gpClear == nil {
		return errors.New("glClear")
	}
	gpClearBufferData = (C.GPCLEARBUFFERDATA)(getProcAddr("glClearBufferData"))
	if gpClearBufferData == nil {
		return errors.New("glClearBufferData")
	}
	gpClearBufferSubData = (C.GPCLEARBUFFERSUBDATA)(getProcAddr("glClearBufferSubData"))
	if gpClearBufferSubData == nil {
		return errors.New("glClearBufferSubData")
	}
	gpClearBufferfi = (C.GPCLEARBUFFERFI)(getProcAddr("glClearBufferfi"))
	if gpClearBufferfi == nil {
		return errors.New("glClearBufferfi")
	}
	gpClearBufferfv = (C.GPCLEARBUFFERFV)(getProcAddr("glClearBufferfv"))
	if gpClearBufferfv == nil {
		return errors.New("glClearBufferfv")
	}
	gpClearBufferiv = (C.GPCLEARBUFFERIV)(getProcAddr("glClearBufferiv"))
	if gpClearBufferiv == nil {
		return errors.New("glClearBufferiv")
	}
	gpClearBufferuiv = (C.GPCLEARBUFFERUIV)(getProcAddr("glClearBufferuiv"))
	if gpClearBufferuiv == nil {
		return errors.New("glClearBufferuiv")
	}
	gpClearColor = (C.GPCLEARCOLOR)(getProcAddr("glClearColor"))
	if gpClearColor == nil {
		return errors.New("glClearColor")
	}
	gpClearDepth = (C.GPCLEARDEPTH)(getProcAddr("glClearDepth"))
	if gpClearDepth == nil {
		return errors.New("glClearDepth")
	}
	gpClearDepthf = (C.GPCLEARDEPTHF)(getProcAddr("glClearDepthf"))
	if gpClearDepthf == nil {
		return errors.New("glClearDepthf")
	}
	gpClearNamedBufferData = (C.GPCLEARNAMEDBUFFERDATA)(getProcAddr("glClearNamedBufferData"))
	if gpClearNamedBufferData == nil {
		return errors.New("glClearNamedBufferData")
	}
	gpClearNamedBufferSubData = (C.GPCLEARNAMEDBUFFERSUBDATA)(getProcAddr("glClearNamedBufferSubData"))
	if gpClearNamedBufferSubData == nil {
		return errors.New("glClearNamedBufferSubData")
	}
	gpClearNamedFramebufferfi = (C.GPCLEARNAMEDFRAMEBUFFERFI)(getProcAddr("glClearNamedFramebufferfi"))
	if gpClearNamedFramebufferfi == nil {
		return errors.New("glClearNamedFramebufferfi")
	}
	gpClearNamedFramebufferfv = (C.GPCLEARNAMEDFRAMEBUFFERFV)(getProcAddr("glClearNamedFramebufferfv"))
	if gpClearNamedFramebufferfv == nil {
		return errors.New("glClearNamedFramebufferfv")
	}
	gpClearNamedFramebufferiv = (C.GPCLEARNAMEDFRAMEBUFFERIV)(getProcAddr("glClearNamedFramebufferiv"))
	if gpClearNamedFramebufferiv == nil {
		return errors.New("glClearNamedFramebufferiv")
	}
	gpClearNamedFramebufferuiv = (C.GPCLEARNAMEDFRAMEBUFFERUIV)(getProcAddr("glClearNamedFramebufferuiv"))
	if gpClearNamedFramebufferuiv == nil {
		return errors.New("glClearNamedFramebufferuiv")
	}
	gpClearStencil = (C.GPCLEARSTENCIL)(getProcAddr("glClearStencil"))
	if gpClearStencil == nil {
		return errors.New("glClearStencil")
	}
	gpClearTexImage = (C.GPCLEARTEXIMAGE)(getProcAddr("glClearTexImage"))
	if gpClearTexImage == nil {
		return errors.New("glClearTexImage")
	}
	gpClearTexSubImage = (C.GPCLEARTEXSUBIMAGE)(getProcAddr("glClearTexSubImage"))
	if gpClearTexSubImage == nil {
		return errors.New("glClearTexSubImage")
	}
	gpClientWaitSync = (C.GPCLIENTWAITSYNC)(getProcAddr("glClientWaitSync"))
	if gpClientWaitSync == nil {
		return errors.New("glClientWaitSync")
	}
	gpClipControl = (C.GPCLIPCONTROL)(getProcAddr("glClipControl"))
	if gpClipControl == nil {
		return errors.New("glClipControl")
	}
	gpColorMask = (C.GPCOLORMASK)(getProcAddr("glColorMask"))
	if gpColorMask == nil {
		return errors.New("glColorMask")
	}
	gpColorMaski = (C.GPCOLORMASKI)(getProcAddr("glColorMaski"))
	if gpColorMaski == nil {
		return errors.New("glColorMaski")
	}
	gpCompileShader = (C.GPCOMPILESHADER)(getProcAddr("glCompileShader"))
	if gpCompileShader == nil {
		return errors.New("glCompileShader")
	}
	gpCompileShaderIncludeARB = (C.GPCOMPILESHADERINCLUDEARB)(getProcAddr("glCompileShaderIncludeARB"))
	gpCompressedTexImage1D = (C.GPCOMPRESSEDTEXIMAGE1D)(getProcAddr("glCompressedTexImage1D"))
	if gpCompressedTexImage1D == nil {
		return errors.New("glCompressedTexImage1D")
	}
	gpCompressedTexImage2D = (C.GPCOMPRESSEDTEXIMAGE2D)(getProcAddr("glCompressedTexImage2D"))
	if gpCompressedTexImage2D == nil {
		return errors.New("glCompressedTexImage2D")
	}
	gpCompressedTexImage3D = (C.GPCOMPRESSEDTEXIMAGE3D)(getProcAddr("glCompressedTexImage3D"))
	if gpCompressedTexImage3D == nil {
		return errors.New("glCompressedTexImage3D")
	}
	gpCompressedTexSubImage1D = (C.GPCOMPRESSEDTEXSUBIMAGE1D)(getProcAddr("glCompressedTexSubImage1D"))
	if gpCompressedTexSubImage1D == nil {
		return errors.New("glCompressedTexSubImage1D")
	}
	gpCompressedTexSubImage2D = (C.GPCOMPRESSEDTEXSUBIMAGE2D)(getProcAddr("glCompressedTexSubImage2D"))
	if gpCompressedTexSubImage2D == nil {
		return errors.New("glCompressedTexSubImage2D")
	}
	gpCompressedTexSubImage3D = (C.GPCOMPRESSEDTEXSUBIMAGE3D)(getProcAddr("glCompressedTexSubImage3D"))
	if gpCompressedTexSubImage3D == nil {
		return errors.New("glCompressedTexSubImage3D")
	}
	gpCompressedTextureSubImage1D = (C.GPCOMPRESSEDTEXTURESUBIMAGE1D)(getProcAddr("glCompressedTextureSubImage1D"))
	if gpCompressedTextureSubImage1D == nil {
		return errors.New("glCompressedTextureSubImage1D")
	}
	gpCompressedTextureSubImage2D = (C.GPCOMPRESSEDTEXTURESUBIMAGE2D)(getProcAddr("glCompressedTextureSubImage2D"))
	if gpCompressedTextureSubImage2D == nil {
		return errors.New("glCompressedTextureSubImage2D")
	}
	gpCompressedTextureSubImage3D = (C.GPCOMPRESSEDTEXTURESUBIMAGE3D)(getProcAddr("glCompressedTextureSubImage3D"))
	if gpCompressedTextureSubImage3D == nil {
		return errors.New("glCompressedTextureSubImage3D")
	}
	gpCopyBufferSubData = (C.GPCOPYBUFFERSUBDATA)(getProcAddr("glCopyBufferSubData"))
	if gpCopyBufferSubData == nil {
		return errors.New("glCopyBufferSubData")
	}
	gpCopyImageSubData = (C.GPCOPYIMAGESUBDATA)(getProcAddr("glCopyImageSubData"))
	if gpCopyImageSubData == nil {
		return errors.New("glCopyImageSubData")
	}
	gpCopyNamedBufferSubData = (C.GPCOPYNAMEDBUFFERSUBDATA)(getProcAddr("glCopyNamedBufferSubData"))
	if gpCopyNamedBufferSubData == nil {
		return errors.New("glCopyNamedBufferSubData")
	}
	gpCopyTexImage1D = (C.GPCOPYTEXIMAGE1D)(getProcAddr("glCopyTexImage1D"))
	if gpCopyTexImage1D == nil {
		return errors.New("glCopyTexImage1D")
	}
	gpCopyTexImage2D = (C.GPCOPYTEXIMAGE2D)(getProcAddr("glCopyTexImage2D"))
	if gpCopyTexImage2D == nil {
		return errors.New("glCopyTexImage2D")
	}
	gpCopyTexSubImage1D = (C.GPCOPYTEXSUBIMAGE1D)(getProcAddr("glCopyTexSubImage1D"))
	if gpCopyTexSubImage1D == nil {
		return errors.New("glCopyTexSubImage1D")
	}
	gpCopyTexSubImage2D = (C.GPCOPYTEXSUBIMAGE2D)(getProcAddr("glCopyTexSubImage2D"))
	if gpCopyTexSubImage2D == nil {
		return errors.New("glCopyTexSubImage2D")
	}
	gpCopyTexSubImage3D = (C.GPCOPYTEXSUBIMAGE3D)(getProcAddr("glCopyTexSubImage3D"))
	if gpCopyTexSubImage3D == nil {
		return errors.New("glCopyTexSubImage3D")
	}
	gpCopyTextureSubImage1D = (C.GPCOPYTEXTURESUBIMAGE1D)(getProcAddr("glCopyTextureSubImage1D"))
	if gpCopyTextureSubImage1D == nil {
		return errors.New("glCopyTextureSubImage1D")
	}
	gpCopyTextureSubImage2D = (C.GPCOPYTEXTURESUBIMAGE2D)(getProcAddr("glCopyTextureSubImage2D"))
	if gpCopyTextureSubImage2D == nil {
		return errors.New("glCopyTextureSubImage2D")
	}
	gpCopyTextureSubImage3D = (C.GPCOPYTEXTURESUBIMAGE3D)(getProcAddr("glCopyTextureSubImage3D"))
	if gpCopyTextureSubImage3D == nil {
		return errors.New("glCopyTextureSubImage3D")
	}
	gpCreateBuffers = (C.GPCREATEBUFFERS)(getProcAddr("glCreateBuffers"))
	if gpCreateBuffers == nil {
		return errors.New("glCreateBuffers")
	}
	gpCreateFramebuffers = (C.GPCREATEFRAMEBUFFERS)(getProcAddr("glCreateFramebuffers"))
	if gpCreateFramebuffers == nil {
		return errors.New("glCreateFramebuffers")
	}
	gpCreateProgram = (C.GPCREATEPROGRAM)(getProcAddr("glCreateProgram"))
	if gpCreateProgram == nil {
		return errors.New("glCreateProgram")
	}
	gpCreateProgramPipelines = (C.GPCREATEPROGRAMPIPELINES)(getProcAddr("glCreateProgramPipelines"))
	if gpCreateProgramPipelines == nil {
		return errors.New("glCreateProgramPipelines")
	}
	gpCreateQueries = (C.GPCREATEQUERIES)(getProcAddr("glCreateQueries"))
	if gpCreateQueries == nil {
		return errors.New("glCreateQueries")
	}
	gpCreateRenderbuffers = (C.GPCREATERENDERBUFFERS)(getProcAddr("glCreateRenderbuffers"))
	if gpCreateRenderbuffers == nil {
		return errors.New("glCreateRenderbuffers")
	}
	gpCreateSamplers = (C.GPCREATESAMPLERS)(getProcAddr("glCreateSamplers"))
	if gpCreateSamplers == nil {
		return errors.New("glCreateSamplers")
	}
	gpCreateShader = (C.GPCREATESHADER)(getProcAddr("glCreateShader"))
	if gpCreateShader == nil {
		return errors.New("glCreateShader")
	}
	gpCreateShaderProgramv = (C.GPCREATESHADERPROGRAMV)(getProcAddr("glCreateShaderProgramv"))
	if gpCreateShaderProgramv == nil {
		return errors.New("glCreateShaderProgramv")
	}
	gpCreateSyncFromCLeventARB = (C.GPCREATESYNCFROMCLEVENTARB)(getProcAddr("glCreateSyncFromCLeventARB"))
	gpCreateTextures = (C.GPCREATETEXTURES)(getProcAddr("glCreateTextures"))
	if gpCreateTextures == nil {
		return errors.New("glCreateTextures")
	}
	gpCreateTransformFeedbacks = (C.GPCREATETRANSFORMFEEDBACKS)(getProcAddr("glCreateTransformFeedbacks"))
	if gpCreateTransformFeedbacks == nil {
		return errors.New("glCreateTransformFeedbacks")
	}
	gpCreateVertexArrays = (C.GPCREATEVERTEXARRAYS)(getProcAddr("glCreateVertexArrays"))
	if gpCreateVertexArrays == nil {
		return errors.New("glCreateVertexArrays")
	}
	gpCullFace = (C.GPCULLFACE)(getProcAddr("glCullFace"))
	if gpCullFace == nil {
		return errors.New("glCullFace")
	}
	gpDebugMessageCallback = (C.GPDEBUGMESSAGECALLBACK)(getProcAddr("glDebugMessageCallback"))
	if gpDebugMessageCallback == nil {
		return errors.New("glDebugMessageCallback")
	}
	gpDebugMessageCallbackARB = (C.GPDEBUGMESSAGECALLBACKARB)(getProcAddr("glDebugMessageCallbackARB"))
	gpDebugMessageCallbackKHR = (C.GPDEBUGMESSAGECALLBACKKHR)(getProcAddr("glDebugMessageCallbackKHR"))
	gpDebugMessageControl = (C.GPDEBUGMESSAGECONTROL)(getProcAddr("glDebugMessageControl"))
	if gpDebugMessageControl == nil {
		return errors.New("glDebugMessageControl")
	}
	gpDebugMessageControlARB = (C.GPDEBUGMESSAGECONTROLARB)(getProcAddr("glDebugMessageControlARB"))
	gpDebugMessageControlKHR = (C.GPDEBUGMESSAGECONTROLKHR)(getProcAddr("glDebugMessageControlKHR"))
	gpDebugMessageInsert = (C.GPDEBUGMESSAGEINSERT)(getProcAddr("glDebugMessageInsert"))
	if gpDebugMessageInsert == nil {
		return errors.New("glDebugMessageInsert")
	}
	gpDebugMessageInsertARB = (C.GPDEBUGMESSAGEINSERTARB)(getProcAddr("glDebugMessageInsertARB"))
	gpDebugMessageInsertKHR = (C.GPDEBUGMESSAGEINSERTKHR)(getProcAddr("glDebugMessageInsertKHR"))
	gpDeleteBuffers = (C.GPDELETEBUFFERS)(getProcAddr("glDeleteBuffers"))
	if gpDeleteBuffers == nil {
		return errors.New("glDeleteBuffers")
	}
	gpDeleteFramebuffers = (C.GPDELETEFRAMEBUFFERS)(getProcAddr("glDeleteFramebuffers"))
	if gpDeleteFramebuffers == nil {
		return errors.New("glDeleteFramebuffers")
	}
	gpDeleteNamedStringARB = (C.GPDELETENAMEDSTRINGARB)(getProcAddr("glDeleteNamedStringARB"))
	gpDeleteProgram = (C.GPDELETEPROGRAM)(getProcAddr("glDeleteProgram"))
	if gpDeleteProgram == nil {
		return errors.New("glDeleteProgram")
	}
	gpDeleteProgramPipelines = (C.GPDELETEPROGRAMPIPELINES)(getProcAddr("glDeleteProgramPipelines"))
	if gpDeleteProgramPipelines == nil {
		return errors.New("glDeleteProgramPipelines")
	}
	gpDeleteQueries = (C.GPDELETEQUERIES)(getProcAddr("glDeleteQueries"))
	if gpDeleteQueries == nil {
		return errors.New("glDeleteQueries")
	}
	gpDeleteRenderbuffers = (C.GPDELETERENDERBUFFERS)(getProcAddr("glDeleteRenderbuffers"))
	if gpDeleteRenderbuffers == nil {
		return errors.New("glDeleteRenderbuffers")
	}
	gpDeleteSamplers = (C.GPDELETESAMPLERS)(getProcAddr("glDeleteSamplers"))
	if gpDeleteSamplers == nil {
		return errors.New("glDeleteSamplers")
	}
	gpDeleteShader = (C.GPDELETESHADER)(getProcAddr("glDeleteShader"))
	if gpDeleteShader == nil {
		return errors.New("glDeleteShader")
	}
	gpDeleteSync = (C.GPDELETESYNC)(getProcAddr("glDeleteSync"))
	if gpDeleteSync == nil {
		return errors.New("glDeleteSync")
	}
	gpDeleteTextures = (C.GPDELETETEXTURES)(getProcAddr("glDeleteTextures"))
	if gpDeleteTextures == nil {
		return errors.New("glDeleteTextures")
	}
	gpDeleteTransformFeedbacks = (C.GPDELETETRANSFORMFEEDBACKS)(getProcAddr("glDeleteTransformFeedbacks"))
	if gpDeleteTransformFeedbacks == nil {
		return errors.New("glDeleteTransformFeedbacks")
	}
	gpDeleteVertexArrays = (C.GPDELETEVERTEXARRAYS)(getProcAddr("glDeleteVertexArrays"))
	if gpDeleteVertexArrays == nil {
		return errors.New("glDeleteVertexArrays")
	}
	gpDepthFunc = (C.GPDEPTHFUNC)(getProcAddr("glDepthFunc"))
	if gpDepthFunc == nil {
		return errors.New("glDepthFunc")
	}
	gpDepthMask = (C.GPDEPTHMASK)(getProcAddr("glDepthMask"))
	if gpDepthMask == nil {
		return errors.New("glDepthMask")
	}
	gpDepthRange = (C.GPDEPTHRANGE)(getProcAddr("glDepthRange"))
	if gpDepthRange == nil {
		return errors.New("glDepthRange")
	}
	gpDepthRangeArrayv = (C.GPDEPTHRANGEARRAYV)(getProcAddr("glDepthRangeArrayv"))
	if gpDepthRangeArrayv == nil {
		return errors.New("glDepthRangeArrayv")
	}
	gpDepthRangeIndexed = (C.GPDEPTHRANGEINDEXED)(getProcAddr("glDepthRangeIndexed"))
	if gpDepthRangeIndexed == nil {
		return errors.New("glDepthRangeIndexed")
	}
	gpDepthRangef = (C.GPDEPTHRANGEF)(getProcAddr("glDepthRangef"))
	if gpDepthRangef == nil {
		return errors.New("glDepthRangef")
	}
	gpDetachShader = (C.GPDETACHSHADER)(getProcAddr("glDetachShader"))
	if gpDetachShader == nil {
		return errors.New("glDetachShader")
	}
	gpDisable = (C.GPDISABLE)(getProcAddr("glDisable"))
	if gpDisable == nil {
		return errors.New("glDisable")
	}
	gpDisableVertexArrayAttrib = (C.GPDISABLEVERTEXARRAYATTRIB)(getProcAddr("glDisableVertexArrayAttrib"))
	if gpDisableVertexArrayAttrib == nil {
		return errors.New("glDisableVertexArrayAttrib")
	}
	gpDisableVertexAttribArray = (C.GPDISABLEVERTEXATTRIBARRAY)(getProcAddr("glDisableVertexAttribArray"))
	if gpDisableVertexAttribArray == nil {
		return errors.New("glDisableVertexAttribArray")
	}
	gpDisablei = (C.GPDISABLEI)(getProcAddr("glDisablei"))
	if gpDisablei == nil {
		return errors.New("glDisablei")
	}
	gpDispatchCompute = (C.GPDISPATCHCOMPUTE)(getProcAddr("glDispatchCompute"))
	if gpDispatchCompute == nil {
		return errors.New("glDispatchCompute")
	}
	gpDispatchComputeGroupSizeARB = (C.GPDISPATCHCOMPUTEGROUPSIZEARB)(getProcAddr("glDispatchComputeGroupSizeARB"))
	gpDispatchComputeIndirect = (C.GPDISPATCHCOMPUTEINDIRECT)(getProcAddr("glDispatchComputeIndirect"))
	if gpDispatchComputeIndirect == nil {
		return errors.New("glDispatchComputeIndirect")
	}
	gpDrawArrays = (C.GPDRAWARRAYS)(getProcAddr("glDrawArrays"))
	if gpDrawArrays == nil {
		return errors.New("glDrawArrays")
	}
	gpDrawArraysIndirect = (C.GPDRAWARRAYSINDIRECT)(getProcAddr("glDrawArraysIndirect"))
	if gpDrawArraysIndirect == nil {
		return errors.New("glDrawArraysIndirect")
	}
	gpDrawArraysInstanced = (C.GPDRAWARRAYSINSTANCED)(getProcAddr("glDrawArraysInstanced"))
	if gpDrawArraysInstanced == nil {
		return errors.New("glDrawArraysInstanced")
	}
	gpDrawArraysInstancedBaseInstance = (C.GPDRAWARRAYSINSTANCEDBASEINSTANCE)(getProcAddr("glDrawArraysInstancedBaseInstance"))
	if gpDrawArraysInstancedBaseInstance == nil {
		return errors.New("glDrawArraysInstancedBaseInstance")
	}
	gpDrawBuffer = (C.GPDRAWBUFFER)(getProcAddr("glDrawBuffer"))
	if gpDrawBuffer == nil {
		return errors.New("glDrawBuffer")
	}
	gpDrawBuffers = (C.GPDRAWBUFFERS)(getProcAddr("glDrawBuffers"))
	if gpDrawBuffers == nil {
		return errors.New("glDrawBuffers")
	}
	gpDrawElements = (C.GPDRAWELEMENTS)(getProcAddr("glDrawElements"))
	if gpDrawElements == nil {
		return errors.New("glDrawElements")
	}
	gpDrawElementsBaseVertex = (C.GPDRAWELEMENTSBASEVERTEX)(getProcAddr("glDrawElementsBaseVertex"))
	if gpDrawElementsBaseVertex == nil {
		return errors.New("glDrawElementsBaseVertex")
	}
	gpDrawElementsIndirect = (C.GPDRAWELEMENTSINDIRECT)(getProcAddr("glDrawElementsIndirect"))
	if gpDrawElementsIndirect == nil {
		return errors.New("glDrawElementsIndirect")
	}
	gpDrawElementsInstanced = (C.GPDRAWELEMENTSINSTANCED)(getProcAddr("glDrawElementsInstanced"))
	if gpDrawElementsInstanced == nil {
		return errors.New("glDrawElementsInstanced")
	}
	gpDrawElementsInstancedBaseInstance = (C.GPDRAWELEMENTSINSTANCEDBASEINSTANCE)(getProcAddr("glDrawElementsInstancedBaseInstance"))
	if gpDrawElementsInstancedBaseInstance == nil {
		return errors.New("glDrawElementsInstancedBaseInstance")
	}
	gpDrawElementsInstancedBaseVertex = (C.GPDRAWELEMENTSINSTANCEDBASEVERTEX)(getProcAddr("glDrawElementsInstancedBaseVertex"))
	if gpDrawElementsInstancedBaseVertex == nil {
		return errors.New("glDrawElementsInstancedBaseVertex")
	}
	gpDrawElementsInstancedBaseVertexBaseInstance = (C.GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE)(getProcAddr("glDrawElementsInstancedBaseVertexBaseInstance"))
	if gpDrawElementsInstancedBaseVertexBaseInstance == nil {
		return errors.New("glDrawElementsInstancedBaseVertexBaseInstance")
	}
	gpDrawRangeElements = (C.GPDRAWRANGEELEMENTS)(getProcAddr("glDrawRangeElements"))
	if gpDrawRangeElements == nil {
		return errors.New("glDrawRangeElements")
	}
	gpDrawRangeElementsBaseVertex = (C.GPDRAWRANGEELEMENTSBASEVERTEX)(getProcAddr("glDrawRangeElementsBaseVertex"))
	if gpDrawRangeElementsBaseVertex == nil {
		return errors.New("glDrawRangeElementsBaseVertex")
	}
	gpDrawTransformFeedback = (C.GPDRAWTRANSFORMFEEDBACK)(getProcAddr("glDrawTransformFeedback"))
	if gpDrawTransformFeedback == nil {
		return errors.New("glDrawTransformFeedback")
	}
	gpDrawTransformFeedbackInstanced = (C.GPDRAWTRANSFORMFEEDBACKINSTANCED)(getProcAddr("glDrawTransformFeedbackInstanced"))
	if gpDrawTransformFeedbackInstanced == nil {
		return errors.New("glDrawTransformFeedbackInstanced")
	}
	gpDrawTransformFeedbackStream = (C.GPDRAWTRANSFORMFEEDBACKSTREAM)(getProcAddr("glDrawTransformFeedbackStream"))
	if gpDrawTransformFeedbackStream == nil {
		return errors.New("glDrawTransformFeedbackStream")
	}
	gpDrawTransformFeedbackStreamInstanced = (C.GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED)(getProcAddr("glDrawTransformFeedbackStreamInstanced"))
	if gpDrawTransformFeedbackStreamInstanced == nil {
		return errors.New("glDrawTransformFeedbackStreamInstanced")
	}
	gpEnable = (C.GPENABLE)(getProcAddr("glEnable"))
	if gpEnable == nil {
		return errors.New("glEnable")
	}
	gpEnableVertexArrayAttrib = (C.GPENABLEVERTEXARRAYATTRIB)(getProcAddr("glEnableVertexArrayAttrib"))
	if gpEnableVertexArrayAttrib == nil {
		return errors.New("glEnableVertexArrayAttrib")
	}
	gpEnableVertexAttribArray = (C.GPENABLEVERTEXATTRIBARRAY)(getProcAddr("glEnableVertexAttribArray"))
	if gpEnableVertexAttribArray == nil {
		return errors.New("glEnableVertexAttribArray")
	}
	gpEnablei = (C.GPENABLEI)(getProcAddr("glEnablei"))
	if gpEnablei == nil {
		return errors.New("glEnablei")
	}
	gpEndConditionalRender = (C.GPENDCONDITIONALRENDER)(getProcAddr("glEndConditionalRender"))
	if gpEndConditionalRender == nil {
		return errors.New("glEndConditionalRender")
	}
	gpEndQuery = (C.GPENDQUERY)(getProcAddr("glEndQuery"))
	if gpEndQuery == nil {
		return errors.New("glEndQuery")
	}
	gpEndQueryIndexed = (C.GPENDQUERYINDEXED)(getProcAddr("glEndQueryIndexed"))
	if gpEndQueryIndexed == nil {
		return errors.New("glEndQueryIndexed")
	}
	gpEndTransformFeedback = (C.GPENDTRANSFORMFEEDBACK)(getProcAddr("glEndTransformFeedback"))
	if gpEndTransformFeedback == nil {
		return errors.New("glEndTransformFeedback")
	}
	gpFenceSync = (C.GPFENCESYNC)(getProcAddr("glFenceSync"))
	if gpFenceSync == nil {
		return errors.New("glFenceSync")
	}
	gpFinish = (C.GPFINISH)(getProcAddr("glFinish"))
	if gpFinish == nil {
		return errors.New("glFinish")
	}
	gpFlush = (C.GPFLUSH)(getProcAddr("glFlush"))
	if gpFlush == nil {
		return errors.New("glFlush")
	}
	gpFlushMappedBufferRange = (C.GPFLUSHMAPPEDBUFFERRANGE)(getProcAddr("glFlushMappedBufferRange"))
	if gpFlushMappedBufferRange == nil {
		return errors.New("glFlushMappedBufferRange")
	}
	gpFlushMappedNamedBufferRange = (C.GPFLUSHMAPPEDNAMEDBUFFERRANGE)(getProcAddr("glFlushMappedNamedBufferRange"))
	if gpFlushMappedNamedBufferRange == nil {
		return errors.New("glFlushMappedNamedBufferRange")
	}
	gpFramebufferParameteri = (C.GPFRAMEBUFFERPARAMETERI)(getProcAddr("glFramebufferParameteri"))
	if gpFramebufferParameteri == nil {
		return errors.New("glFramebufferParameteri")
	}
	gpFramebufferRenderbuffer = (C.GPFRAMEBUFFERRENDERBUFFER)(getProcAddr("glFramebufferRenderbuffer"))
	if gpFramebufferRenderbuffer == nil {
		return errors.New("glFramebufferRenderbuffer")
	}
	gpFramebufferTexture = (C.GPFRAMEBUFFERTEXTURE)(getProcAddr("glFramebufferTexture"))
	if gpFramebufferTexture == nil {
		return errors.New("glFramebufferTexture")
	}
	gpFramebufferTexture1D = (C.GPFRAMEBUFFERTEXTURE1D)(getProcAddr("glFramebufferTexture1D"))
	if gpFramebufferTexture1D == nil {
		return errors.New("glFramebufferTexture1D")
	}
	gpFramebufferTexture2D = (C.GPFRAMEBUFFERTEXTURE2D)(getProcAddr("glFramebufferTexture2D"))
	if gpFramebufferTexture2D == nil {
		return errors.New("glFramebufferTexture2D")
	}
	gpFramebufferTexture3D = (C.GPFRAMEBUFFERTEXTURE3D)(getProcAddr("glFramebufferTexture3D"))
	if gpFramebufferTexture3D == nil {
		return errors.New("glFramebufferTexture3D")
	}
	gpFramebufferTextureLayer = (C.GPFRAMEBUFFERTEXTURELAYER)(getProcAddr("glFramebufferTextureLayer"))
	if gpFramebufferTextureLayer == nil {
		return errors.New("glFramebufferTextureLayer")
	}
	gpFrontFace = (C.GPFRONTFACE)(getProcAddr("glFrontFace"))
	if gpFrontFace == nil {
		return errors.New("glFrontFace")
	}
	gpGenBuffers = (C.GPGENBUFFERS)(getProcAddr("glGenBuffers"))
	if gpGenBuffers == nil {
		return errors.New("glGenBuffers")
	}
	gpGenFramebuffers = (C.GPGENFRAMEBUFFERS)(getProcAddr("glGenFramebuffers"))
	if gpGenFramebuffers == nil {
		return errors.New("glGenFramebuffers")
	}
	gpGenProgramPipelines = (C.GPGENPROGRAMPIPELINES)(getProcAddr("glGenProgramPipelines"))
	if gpGenProgramPipelines == nil {
		return errors.New("glGenProgramPipelines")
	}
	gpGenQueries = (C.GPGENQUERIES)(getProcAddr("glGenQueries"))
	if gpGenQueries == nil {
		return errors.New("glGenQueries")
	}
	gpGenRenderbuffers = (C.GPGENRENDERBUFFERS)(getProcAddr("glGenRenderbuffers"))
	if gpGenRenderbuffers == nil {
		return errors.New("glGenRenderbuffers")
	}
	gpGenSamplers = (C.GPGENSAMPLERS)(getProcAddr("glGenSamplers"))
	if gpGenSamplers == nil {
		return errors.New("glGenSamplers")
	}
	gpGenTextures = (C.GPGENTEXTURES)(getProcAddr("glGenTextures"))
	if gpGenTextures == nil {
		return errors.New("glGenTextures")
	}
	gpGenTransformFeedbacks = (C.GPGENTRANSFORMFEEDBACKS)(getProcAddr("glGenTransformFeedbacks"))
	if gpGenTransformFeedbacks == nil {
		return errors.New("glGenTransformFeedbacks")
	}
	gpGenVertexArrays = (C.GPGENVERTEXARRAYS)(getProcAddr("glGenVertexArrays"))
	if gpGenVertexArrays == nil {
		return errors.New("glGenVertexArrays")
	}
	gpGenerateMipmap = (C.GPGENERATEMIPMAP)(getProcAddr("glGenerateMipmap"))
	if gpGenerateMipmap == nil {
		return errors.New("glGenerateMipmap")
	}
	gpGenerateTextureMipmap = (C.GPGENERATETEXTUREMIPMAP)(getProcAddr("glGenerateTextureMipmap"))
	if gpGenerateTextureMipmap == nil {
		return errors.New("glGenerateTextureMipmap")
	}
	gpGetActiveAtomicCounterBufferiv = (C.GPGETACTIVEATOMICCOUNTERBUFFERIV)(getProcAddr("glGetActiveAtomicCounterBufferiv"))
	if gpGetActiveAtomicCounterBufferiv == nil {
		return errors.New("glGetActiveAtomicCounterBufferiv")
	}
	gpGetActiveAttrib = (C.GPGETACTIVEATTRIB)(getProcAddr("glGetActiveAttrib"))
	if gpGetActiveAttrib == nil {
		return errors.New("glGetActiveAttrib")
	}
	gpGetActiveSubroutineName = (C.GPGETACTIVESUBROUTINENAME)(getProcAddr("glGetActiveSubroutineName"))
	if gpGetActiveSubroutineName == nil {
		return errors.New("glGetActiveSubroutineName")
	}
	gpGetActiveSubroutineUniformName = (C.GPGETACTIVESUBROUTINEUNIFORMNAME)(getProcAddr("glGetActiveSubroutineUniformName"))
	if gpGetActiveSubroutineUniformName == nil {
		return errors.New("glGetActiveSubroutineUniformName")
	}
	gpGetActiveSubroutineUniformiv = (C.GPGETACTIVESUBROUTINEUNIFORMIV)(getProcAddr("glGetActiveSubroutineUniformiv"))
	if gpGetActiveSubroutineUniformiv == nil {
		return errors.New("glGetActiveSubroutineUniformiv")
	}
	gpGetActiveUniform = (C.GPGETACTIVEUNIFORM)(getProcAddr("glGetActiveUniform"))
	if gpGetActiveUniform == nil {
		return errors.New("glGetActiveUniform")
	}
	gpGetActiveUniformBlockName = (C.GPGETACTIVEUNIFORMBLOCKNAME)(getProcAddr("glGetActiveUniformBlockName"))
	if gpGetActiveUniformBlockName == nil {
		return errors.New("glGetActiveUniformBlockName")
	}
	gpGetActiveUniformBlockiv = (C.GPGETACTIVEUNIFORMBLOCKIV)(getProcAddr("glGetActiveUniformBlockiv"))
	if gpGetActiveUniformBlockiv == nil {
		return errors.New("glGetActiveUniformBlockiv")
	}
	gpGetActiveUniformName = (C.GPGETACTIVEUNIFORMNAME)(getProcAddr("glGetActiveUniformName"))
	if gpGetActiveUniformName == nil {
		return errors.New("glGetActiveUniformName")
	}
	gpGetActiveUniformsiv = (C.GPGETACTIVEUNIFORMSIV)(getProcAddr("glGetActiveUniformsiv"))
	if gpGetActiveUniformsiv == nil {
		return errors.New("glGetActiveUniformsiv")
	}
	gpGetAttachedShaders = (C.GPGETATTACHEDSHADERS)(getProcAddr("glGetAttachedShaders"))
	if gpGetAttachedShaders == nil {
		return errors.New("glGetAttachedShaders")
	}
	gpGetAttribLocation = (C.GPGETATTRIBLOCATION)(getProcAddr("glGetAttribLocation"))
	if gpGetAttribLocation == nil {
		return errors.New("glGetAttribLocation")
	}
	gpGetBooleani_v = (C.GPGETBOOLEANI_V)(getProcAddr("glGetBooleani_v"))
	if gpGetBooleani_v == nil {
		return errors.New("glGetBooleani_v")
	}
	gpGetBooleanv = (C.GPGETBOOLEANV)(getProcAddr("glGetBooleanv"))
	if gpGetBooleanv == nil {
		return errors.New("glGetBooleanv")
	}
	gpGetBufferParameteri64v = (C.GPGETBUFFERPARAMETERI64V)(getProcAddr("glGetBufferParameteri64v"))
	if gpGetBufferParameteri64v == nil {
		return errors.New("glGetBufferParameteri64v")
	}
	gpGetBufferParameteriv = (C.GPGETBUFFERPARAMETERIV)(getProcAddr("glGetBufferParameteriv"))
	if gpGetBufferParameteriv == nil {
		return errors.New("glGetBufferParameteriv")
	}
	gpGetBufferPointerv = (C.GPGETBUFFERPOINTERV)(getProcAddr("glGetBufferPointerv"))
	if gpGetBufferPointerv == nil {
		return errors.New("glGetBufferPointerv")
	}
	gpGetBufferSubData = (C.GPGETBUFFERSUBDATA)(getProcAddr("glGetBufferSubData"))
	if gpGetBufferSubData == nil {
		return errors.New("glGetBufferSubData")
	}
	gpGetCompressedTexImage = (C.GPGETCOMPRESSEDTEXIMAGE)(getProcAddr("glGetCompressedTexImage"))
	if gpGetCompressedTexImage == nil {
		return errors.New("glGetCompressedTexImage")
	}
	gpGetCompressedTextureImage = (C.GPGETCOMPRESSEDTEXTUREIMAGE)(getProcAddr("glGetCompressedTextureImage"))
	if gpGetCompressedTextureImage == nil {
		return errors.New("glGetCompressedTextureImage")
	}
	gpGetCompressedTextureSubImage = (C.GPGETCOMPRESSEDTEXTURESUBIMAGE)(getProcAddr("glGetCompressedTextureSubImage"))
	if gpGetCompressedTextureSubImage == nil {
		return errors.New("glGetCompressedTextureSubImage")
	}
	gpGetDebugMessageLog = (C.GPGETDEBUGMESSAGELOG)(getProcAddr("glGetDebugMessageLog"))
	if gpGetDebugMessageLog == nil {
		return errors.New("glGetDebugMessageLog")
	}
	gpGetDebugMessageLogARB = (C.GPGETDEBUGMESSAGELOGARB)(getProcAddr("glGetDebugMessageLogARB"))
	gpGetDebugMessageLogKHR = (C.GPGETDEBUGMESSAGELOGKHR)(getProcAddr("glGetDebugMessageLogKHR"))
	gpGetDoublei_v = (C.GPGETDOUBLEI_V)(getProcAddr("glGetDoublei_v"))
	if gpGetDoublei_v == nil {
		return errors.New("glGetDoublei_v")
	}
	gpGetDoublev = (C.GPGETDOUBLEV)(getProcAddr("glGetDoublev"))
	if gpGetDoublev == nil {
		return errors.New("glGetDoublev")
	}
	gpGetError = (C.GPGETERROR)(getProcAddr("glGetError"))
	if gpGetError == nil {
		return errors.New("glGetError")
	}
	gpGetFloati_v = (C.GPGETFLOATI_V)(getProcAddr("glGetFloati_v"))
	if gpGetFloati_v == nil {
		return errors.New("glGetFloati_v")
	}
	gpGetFloatv = (C.GPGETFLOATV)(getProcAddr("glGetFloatv"))
	if gpGetFloatv == nil {
		return errors.New("glGetFloatv")
	}
	gpGetFragDataIndex = (C.GPGETFRAGDATAINDEX)(getProcAddr("glGetFragDataIndex"))
	if gpGetFragDataIndex == nil {
		return errors.New("glGetFragDataIndex")
	}
	gpGetFragDataLocation = (C.GPGETFRAGDATALOCATION)(getProcAddr("glGetFragDataLocation"))
	if gpGetFragDataLocation == nil {
		return errors.New("glGetFragDataLocation")
	}
	gpGetFramebufferAttachmentParameteriv = (C.GPGETFRAMEBUFFERATTACHMENTPARAMETERIV)(getProcAddr("glGetFramebufferAttachmentParameteriv"))
	if gpGetFramebufferAttachmentParameteriv == nil {
		return errors.New("glGetFramebufferAttachmentParameteriv")
	}
	gpGetFramebufferParameteriv = (C.GPGETFRAMEBUFFERPARAMETERIV)(getProcAddr("glGetFramebufferParameteriv"))
	if gpGetFramebufferParameteriv == nil {
		return errors.New("glGetFramebufferParameteriv")
	}
	gpGetGraphicsResetStatus = (C.GPGETGRAPHICSRESETSTATUS)(getProcAddr("glGetGraphicsResetStatus"))
	if gpGetGraphicsResetStatus == nil {
		return errors.New("glGetGraphicsResetStatus")
	}
	gpGetGraphicsResetStatusARB = (C.GPGETGRAPHICSRESETSTATUSARB)(getProcAddr("glGetGraphicsResetStatusARB"))
	gpGetGraphicsResetStatusKHR = (C.GPGETGRAPHICSRESETSTATUSKHR)(getProcAddr("glGetGraphicsResetStatusKHR"))
	gpGetImageHandleARB = (C.GPGETIMAGEHANDLEARB)(getProcAddr("glGetImageHandleARB"))
	gpGetInteger64i_v = (C.GPGETINTEGER64I_V)(getProcAddr("glGetInteger64i_v"))
	if gpGetInteger64i_v == nil {
		return errors.New("glGetInteger64i_v")
	}
	gpGetInteger64v = (C.GPGETINTEGER64V)(getProcAddr("glGetInteger64v"))
	if gpGetInteger64v == nil {
		return errors.New("glGetInteger64v")
	}
	gpGetIntegeri_v = (C.GPGETINTEGERI_V)(getProcAddr("glGetIntegeri_v"))
	if gpGetIntegeri_v == nil {
		return errors.New("glGetIntegeri_v")
	}
	gpGetIntegerv = (C.GPGETINTEGERV)(getProcAddr("glGetIntegerv"))
	if gpGetIntegerv == nil {
		return errors.New("glGetIntegerv")
	}
	gpGetInternalformati64v = (C.GPGETINTERNALFORMATI64V)(getProcAddr("glGetInternalformati64v"))
	if gpGetInternalformati64v == nil {
		return errors.New("glGetInternalformati64v")
	}
	gpGetInternalformativ = (C.GPGETINTERNALFORMATIV)(getProcAddr("glGetInternalformativ"))
	if gpGetInternalformativ == nil {
		return errors.New("glGetInternalformativ")
	}
	gpGetMultisamplefv = (C.GPGETMULTISAMPLEFV)(getProcAddr("glGetMultisamplefv"))
	if gpGetMultisamplefv == nil {
		return errors.New("glGetMultisamplefv")
	}
	gpGetNamedBufferParameteri64v = (C.GPGETNAMEDBUFFERPARAMETERI64V)(getProcAddr("glGetNamedBufferParameteri64v"))
	if gpGetNamedBufferParameteri64v == nil {
		return errors.New("glGetNamedBufferParameteri64v")
	}
	gpGetNamedBufferParameteriv = (C.GPGETNAMEDBUFFERPARAMETERIV)(getProcAddr("glGetNamedBufferParameteriv"))
	if gpGetNamedBufferParameteriv == nil {
		return errors.New("glGetNamedBufferParameteriv")
	}
	gpGetNamedBufferPointerv = (C.GPGETNAMEDBUFFERPOINTERV)(getProcAddr("glGetNamedBufferPointerv"))
	if gpGetNamedBufferPointerv == nil {
		return errors.New("glGetNamedBufferPointerv")
	}
	gpGetNamedBufferSubData = (C.GPGETNAMEDBUFFERSUBDATA)(getProcAddr("glGetNamedBufferSubData"))
	if gpGetNamedBufferSubData == nil {
		return errors.New("glGetNamedBufferSubData")
	}
	gpGetNamedFramebufferAttachmentParameteriv = (C.GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV)(getProcAddr("glGetNamedFramebufferAttachmentParameteriv"))
	if gpGetNamedFramebufferAttachmentParameteriv == nil {
		return errors.New("glGetNamedFramebufferAttachmentParameteriv")
	}
	gpGetNamedFramebufferParameteriv = (C.GPGETNAMEDFRAMEBUFFERPARAMETERIV)(getProcAddr("glGetNamedFramebufferParameteriv"))
	if gpGetNamedFramebufferParameteriv == nil {
		return errors.New("glGetNamedFramebufferParameteriv")
	}
	gpGetNamedRenderbufferParameteriv = (C.GPGETNAMEDRENDERBUFFERPARAMETERIV)(getProcAddr("glGetNamedRenderbufferParameteriv"))
	if gpGetNamedRenderbufferParameteriv == nil {
		return errors.New("glGetNamedRenderbufferParameteriv")
	}
	gpGetNamedStringARB = (C.GPGETNAMEDSTRINGARB)(getProcAddr("glGetNamedStringARB"))
	gpGetNamedStringivARB = (C.GPGETNAMEDSTRINGIVARB)(getProcAddr("glGetNamedStringivARB"))
	gpGetObjectLabel = (C.GPGETOBJECTLABEL)(getProcAddr("glGetObjectLabel"))
	if gpGetObjectLabel == nil {
		return errors.New("glGetObjectLabel")
	}
	gpGetObjectLabelKHR = (C.GPGETOBJECTLABELKHR)(getProcAddr("glGetObjectLabelKHR"))
	gpGetObjectPtrLabel = (C.GPGETOBJECTPTRLABEL)(getProcAddr("glGetObjectPtrLabel"))
	if gpGetObjectPtrLabel == nil {
		return errors.New("glGetObjectPtrLabel")
	}
	gpGetObjectPtrLabelKHR = (C.GPGETOBJECTPTRLABELKHR)(getProcAddr("glGetObjectPtrLabelKHR"))
	gpGetPointerv = (C.GPGETPOINTERV)(getProcAddr("glGetPointerv"))
	if gpGetPointerv == nil {
		return errors.New("glGetPointerv")
	}
	gpGetPointervKHR = (C.GPGETPOINTERVKHR)(getProcAddr("glGetPointervKHR"))
	gpGetProgramBinary = (C.GPGETPROGRAMBINARY)(getProcAddr("glGetProgramBinary"))
	if gpGetProgramBinary == nil {
		return errors.New("glGetProgramBinary")
	}
	gpGetProgramInfoLog = (C.GPGETPROGRAMINFOLOG)(getProcAddr("glGetProgramInfoLog"))
	if gpGetProgramInfoLog == nil {
		return errors.New("glGetProgramInfoLog")
	}
	gpGetProgramInterfaceiv = (C.GPGETPROGRAMINTERFACEIV)(getProcAddr("glGetProgramInterfaceiv"))
	if gpGetProgramInterfaceiv == nil {
		return errors.New("glGetProgramInterfaceiv")
	}
	gpGetProgramPipelineInfoLog = (C.GPGETPROGRAMPIPELINEINFOLOG)(getProcAddr("glGetProgramPipelineInfoLog"))
	if gpGetProgramPipelineInfoLog == nil {
		return errors.New("glGetProgramPipelineInfoLog")
	}
	gpGetProgramPipelineiv = (C.GPGETPROGRAMPIPELINEIV)(getProcAddr("glGetProgramPipelineiv"))
	if gpGetProgramPipelineiv == nil {
		return errors.New("glGetProgramPipelineiv")
	}
	gpGetProgramResourceIndex = (C.GPGETPROGRAMRESOURCEINDEX)(getProcAddr("glGetProgramResourceIndex"))
	if gpGetProgramResourceIndex == nil {
		return errors.New("glGetProgramResourceIndex")
	}
	gpGetProgramResourceLocation = (C.GPGETPROGRAMRESOURCELOCATION)(getProcAddr("glGetProgramResourceLocation"))
	if gpGetProgramResourceLocation == nil {
		return errors.New("glGetProgramResourceLocation")
	}
	gpGetProgramResourceLocationIndex = (C.GPGETPROGRAMRESOURCELOCATIONINDEX)(getProcAddr("glGetProgramResourceLocationIndex"))
	if gpGetProgramResourceLocationIndex == nil {
		return errors.New("glGetProgramResourceLocationIndex")
	}
	gpGetProgramResourceName = (C.GPGETPROGRAMRESOURCENAME)(getProcAddr("glGetProgramResourceName"))
	if gpGetProgramResourceName == nil {
		return errors.New("glGetProgramResourceName")
	}
	gpGetProgramResourceiv = (C.GPGETPROGRAMRESOURCEIV)(getProcAddr("glGetProgramResourceiv"))
	if gpGetProgramResourceiv == nil {
		return errors.New("glGetProgramResourceiv")
	}
	gpGetProgramStageiv = (C.GPGETPROGRAMSTAGEIV)(getProcAddr("glGetProgramStageiv"))
	if gpGetProgramStageiv == nil {
		return errors.New("glGetProgramStageiv")
	}
	gpGetProgramiv = (C.GPGETPROGRAMIV)(getProcAddr("glGetProgramiv"))
	if gpGetProgramiv == nil {
		return errors.New("glGetProgramiv")
	}
	gpGetQueryIndexediv = (C.GPGETQUERYINDEXEDIV)(getProcAddr("glGetQueryIndexediv"))
	if gpGetQueryIndexediv == nil {
		return errors.New("glGetQueryIndexediv")
	}
	gpGetQueryObjecti64v = (C.GPGETQUERYOBJECTI64V)(getProcAddr("glGetQueryObjecti64v"))
	if gpGetQueryObjecti64v == nil {
		return errors.New("glGetQueryObjecti64v")
	}
	gpGetQueryObjectiv = (C.GPGETQUERYOBJECTIV)(getProcAddr("glGetQueryObjectiv"))
	if gpGetQueryObjectiv == nil {
		return errors.New("glGetQueryObjectiv")
	}
	gpGetQueryObjectui64v = (C.GPGETQUERYOBJECTUI64V)(getProcAddr("glGetQueryObjectui64v"))
	if gpGetQueryObjectui64v == nil {
		return errors.New("glGetQueryObjectui64v")
	}
	gpGetQueryObjectuiv = (C.GPGETQUERYOBJECTUIV)(getProcAddr("glGetQueryObjectuiv"))
	if gpGetQueryObjectuiv == nil {
		return errors.New("glGetQueryObjectuiv")
	}
	gpGetQueryiv = (C.GPGETQUERYIV)(getProcAddr("glGetQueryiv"))
	if gpGetQueryiv == nil {
		return errors.New("glGetQueryiv")
	}
	gpGetRenderbufferParameteriv = (C.GPGETRENDERBUFFERPARAMETERIV)(getProcAddr("glGetRenderbufferParameteriv"))
	if gpGetRenderbufferParameteriv == nil {
		return errors.New("glGetRenderbufferParameteriv")
	}
	gpGetSamplerParameterIiv = (C.GPGETSAMPLERPARAMETERIIV)(getProcAddr("glGetSamplerParameterIiv"))
	if gpGetSamplerParameterIiv == nil {
		return errors.New("glGetSamplerParameterIiv")
	}
	gpGetSamplerParameterIuiv = (C.GPGETSAMPLERPARAMETERIUIV)(getProcAddr("glGetSamplerParameterIuiv"))
	if gpGetSamplerParameterIuiv == nil {
		return errors.New("glGetSamplerParameterIuiv")
	}
	gpGetSamplerParameterfv = (C.GPGETSAMPLERPARAMETERFV)(getProcAddr("glGetSamplerParameterfv"))
	if gpGetSamplerParameterfv == nil {
		return errors.New("glGetSamplerParameterfv")
	}
	gpGetSamplerParameteriv = (C.GPGETSAMPLERPARAMETERIV)(getProcAddr("glGetSamplerParameteriv"))
	if gpGetSamplerParameteriv == nil {
		return errors.New("glGetSamplerParameteriv")
	}
	gpGetShaderInfoLog = (C.GPGETSHADERINFOLOG)(getProcAddr("glGetShaderInfoLog"))
	if gpGetShaderInfoLog == nil {
		return errors.New("glGetShaderInfoLog")
	}
	gpGetShaderPrecisionFormat = (C.GPGETSHADERPRECISIONFORMAT)(getProcAddr("glGetShaderPrecisionFormat"))
	if gpGetShaderPrecisionFormat == nil {
		return errors.New("glGetShaderPrecisionFormat")
	}
	gpGetShaderSource = (C.GPGETSHADERSOURCE)(getProcAddr("glGetShaderSource"))
	if gpGetShaderSource == nil {
		return errors.New("glGetShaderSource")
	}
	gpGetShaderiv = (C.GPGETSHADERIV)(getProcAddr("glGetShaderiv"))
	if gpGetShaderiv == nil {
		return errors.New("glGetShaderiv")
	}
	gpGetString = (C.GPGETSTRING)(getProcAddr("glGetString"))
	if gpGetString == nil {
		return errors.New("glGetString")
	}
	gpGetStringi = (C.GPGETSTRINGI)(getProcAddr("glGetStringi"))
	if gpGetStringi == nil {
		return errors.New("glGetStringi")
	}
	gpGetSubroutineIndex = (C.GPGETSUBROUTINEINDEX)(getProcAddr("glGetSubroutineIndex"))
	if gpGetSubroutineIndex == nil {
		return errors.New("glGetSubroutineIndex")
	}
	gpGetSubroutineUniformLocation = (C.GPGETSUBROUTINEUNIFORMLOCATION)(getProcAddr("glGetSubroutineUniformLocation"))
	if gpGetSubroutineUniformLocation == nil {
		return errors.New("glGetSubroutineUniformLocation")
	}
	gpGetSynciv = (C.GPGETSYNCIV)(getProcAddr("glGetSynciv"))
	if gpGetSynciv == nil {
		return errors.New("glGetSynciv")
	}
	gpGetTexImage = (C.GPGETTEXIMAGE)(getProcAddr("glGetTexImage"))
	if gpGetTexImage == nil {
		return errors.New("glGetTexImage")
	}
	gpGetTexLevelParameterfv = (C.GPGETTEXLEVELPARAMETERFV)(getProcAddr("glGetTexLevelParameterfv"))
	if gpGetTexLevelParameterfv == nil {
		return errors.New("glGetTexLevelParameterfv")
	}
	gpGetTexLevelParameteriv = (C.GPGETTEXLEVELPARAMETERIV)(getProcAddr("glGetTexLevelParameteriv"))
	if gpGetTexLevelParameteriv == nil {
		return errors.New("glGetTexLevelParameteriv")
	}
	gpGetTexParameterIiv = (C.GPGETTEXPARAMETERIIV)(getProcAddr("glGetTexParameterIiv"))
	if gpGetTexParameterIiv == nil {
		return errors.New("glGetTexParameterIiv")
	}
	gpGetTexParameterIuiv = (C.GPGETTEXPARAMETERIUIV)(getProcAddr("glGetTexParameterIuiv"))
	if gpGetTexParameterIuiv == nil {
		return errors.New("glGetTexParameterIuiv")
	}
	gpGetTexParameterfv = (C.GPGETTEXPARAMETERFV)(getProcAddr("glGetTexParameterfv"))
	if gpGetTexParameterfv == nil {
		return errors.New("glGetTexParameterfv")
	}
	gpGetTexParameteriv = (C.GPGETTEXPARAMETERIV)(getProcAddr("glGetTexParameteriv"))
	if gpGetTexParameteriv == nil {
		return errors.New("glGetTexParameteriv")
	}
	gpGetTextureHandleARB = (C.GPGETTEXTUREHANDLEARB)(getProcAddr("glGetTextureHandleARB"))
	gpGetTextureImage = (C.GPGETTEXTUREIMAGE)(getProcAddr("glGetTextureImage"))
	if gpGetTextureImage == nil {
		return errors.New("glGetTextureImage")
	}
	gpGetTextureLevelParameterfv = (C.GPGETTEXTURELEVELPARAMETERFV)(getProcAddr("glGetTextureLevelParameterfv"))
	if gpGetTextureLevelParameterfv == nil {
		return errors.New("glGetTextureLevelParameterfv")
	}
	gpGetTextureLevelParameteriv = (C.GPGETTEXTURELEVELPARAMETERIV)(getProcAddr("glGetTextureLevelParameteriv"))
	if gpGetTextureLevelParameteriv == nil {
		return errors.New("glGetTextureLevelParameteriv")
	}
	gpGetTextureParameterIiv = (C.GPGETTEXTUREPARAMETERIIV)(getProcAddr("glGetTextureParameterIiv"))
	if gpGetTextureParameterIiv == nil {
		return errors.New("glGetTextureParameterIiv")
	}
	gpGetTextureParameterIuiv = (C.GPGETTEXTUREPARAMETERIUIV)(getProcAddr("glGetTextureParameterIuiv"))
	if gpGetTextureParameterIuiv == nil {
		return errors.New("glGetTextureParameterIuiv")
	}
	gpGetTextureParameterfv = (C.GPGETTEXTUREPARAMETERFV)(getProcAddr("glGetTextureParameterfv"))
	if gpGetTextureParameterfv == nil {
		return errors.New("glGetTextureParameterfv")
	}
	gpGetTextureParameteriv = (C.GPGETTEXTUREPARAMETERIV)(getProcAddr("glGetTextureParameteriv"))
	if gpGetTextureParameteriv == nil {
		return errors.New("glGetTextureParameteriv")
	}
	gpGetTextureSamplerHandleARB = (C.GPGETTEXTURESAMPLERHANDLEARB)(getProcAddr("glGetTextureSamplerHandleARB"))
	gpGetTextureSubImage = (C.GPGETTEXTURESUBIMAGE)(getProcAddr("glGetTextureSubImage"))
	if gpGetTextureSubImage == nil {
		return errors.New("glGetTextureSubImage")
	}
	gpGetTransformFeedbackVarying = (C.GPGETTRANSFORMFEEDBACKVARYING)(getProcAddr("glGetTransformFeedbackVarying"))
	if gpGetTransformFeedbackVarying == nil {
		return errors.New("glGetTransformFeedbackVarying")
	}
	gpGetTransformFeedbacki64_v = (C.GPGETTRANSFORMFEEDBACKI64_V)(getProcAddr("glGetTransformFeedbacki64_v"))
	if gpGetTransformFeedbacki64_v == nil {
		return errors.New("glGetTransformFeedbacki64_v")
	}
	gpGetTransformFeedbacki_v = (C.GPGETTRANSFORMFEEDBACKI_V)(getProcAddr("glGetTransformFeedbacki_v"))
	if gpGetTransformFeedbacki_v == nil {
		return errors.New("glGetTransformFeedbacki_v")
	}
	gpGetTransformFeedbackiv = (C.GPGETTRANSFORMFEEDBACKIV)(getProcAddr("glGetTransformFeedbackiv"))
	if gpGetTransformFeedbackiv == nil {
		return errors.New("glGetTransformFeedbackiv")
	}
	gpGetUniformBlockIndex = (C.GPGETUNIFORMBLOCKINDEX)(getProcAddr("glGetUniformBlockIndex"))
	if gpGetUniformBlockIndex == nil {
		return errors.New("glGetUniformBlockIndex")
	}
	gpGetUniformIndices = (C.GPGETUNIFORMINDICES)(getProcAddr("glGetUniformIndices"))
	if gpGetUniformIndices == nil {
		return errors.New("glGetUniformIndices")
	}
	gpGetUniformLocation = (C.GPGETUNIFORMLOCATION)(getProcAddr("glGetUniformLocation"))
	if gpGetUniformLocation == nil {
		return errors.New("glGetUniformLocation")
	}
	gpGetUniformSubroutineuiv = (C.GPGETUNIFORMSUBROUTINEUIV)(getProcAddr("glGetUniformSubroutineuiv"))
	if gpGetUniformSubroutineuiv == nil {
		return errors.New("glGetUniformSubroutineuiv")
	}
	gpGetUniformdv = (C.GPGETUNIFORMDV)(getProcAddr("glGetUniformdv"))
	if gpGetUniformdv == nil {
		return errors.New("glGetUniformdv")
	}
	gpGetUniformfv = (C.GPGETUNIFORMFV)(getProcAddr("glGetUniformfv"))
	if gpGetUniformfv == nil {
		return errors.New("glGetUniformfv")
	}
	gpGetUniformiv = (C.GPGETUNIFORMIV)(getProcAddr("glGetUniformiv"))
	if gpGetUniformiv == nil {
		return errors.New("glGetUniformiv")
	}
	gpGetUniformuiv = (C.GPGETUNIFORMUIV)(getProcAddr("glGetUniformuiv"))
	if gpGetUniformuiv == nil {
		return errors.New("glGetUniformuiv")
	}
	gpGetVertexArrayIndexed64iv = (C.GPGETVERTEXARRAYINDEXED64IV)(getProcAddr("glGetVertexArrayIndexed64iv"))
	if gpGetVertexArrayIndexed64iv == nil {
		return errors.New("glGetVertexArrayIndexed64iv")
	}
	gpGetVertexArrayIndexediv = (C.GPGETVERTEXARRAYINDEXEDIV)(getProcAddr("glGetVertexArrayIndexediv"))
	if gpGetVertexArrayIndexediv == nil {
		return errors.New("glGetVertexArrayIndexediv")
	}
	gpGetVertexArrayiv = (C.GPGETVERTEXARRAYIV)(getProcAddr("glGetVertexArrayiv"))
	if gpGetVertexArrayiv == nil {
		return errors.New("glGetVertexArrayiv")
	}
	gpGetVertexAttribIiv = (C.GPGETVERTEXATTRIBIIV)(getProcAddr("glGetVertexAttribIiv"))
	if gpGetVertexAttribIiv == nil {
		return errors.New("glGetVertexAttribIiv")
	}
	gpGetVertexAttribIuiv = (C.GPGETVERTEXATTRIBIUIV)(getProcAddr("glGetVertexAttribIuiv"))
	if gpGetVertexAttribIuiv == nil {
		return errors.New("glGetVertexAttribIuiv")
	}
	gpGetVertexAttribLdv = (C.GPGETVERTEXATTRIBLDV)(getProcAddr("glGetVertexAttribLdv"))
	if gpGetVertexAttribLdv == nil {
		return errors.New("glGetVertexAttribLdv")
	}
	gpGetVertexAttribLui64vARB = (C.GPGETVERTEXATTRIBLUI64VARB)(getProcAddr("glGetVertexAttribLui64vARB"))
	gpGetVertexAttribPointerv = (C.GPGETVERTEXATTRIBPOINTERV)(getProcAddr("glGetVertexAttribPointerv"))
	if gpGetVertexAttribPointerv == nil {
		return errors.New("glGetVertexAttribPointerv")
	}
	gpGetVertexAttribdv = (C.GPGETVERTEXATTRIBDV)(getProcAddr("glGetVertexAttribdv"))
	if gpGetVertexAttribdv == nil {
		return errors.New("glGetVertexAttribdv")
	}
	gpGetVertexAttribfv = (C.GPGETVERTEXATTRIBFV)(getProcAddr("glGetVertexAttribfv"))
	if gpGetVertexAttribfv == nil {
		return errors.New("glGetVertexAttribfv")
	}
	gpGetVertexAttribiv = (C.GPGETVERTEXATTRIBIV)(getProcAddr("glGetVertexAttribiv"))
	if gpGetVertexAttribiv == nil {
		return errors.New("glGetVertexAttribiv")
	}
	gpGetnCompressedTexImage = (C.GPGETNCOMPRESSEDTEXIMAGE)(getProcAddr("glGetnCompressedTexImage"))
	if gpGetnCompressedTexImage == nil {
		return errors.New("glGetnCompressedTexImage")
	}
	gpGetnCompressedTexImageARB = (C.GPGETNCOMPRESSEDTEXIMAGEARB)(getProcAddr("glGetnCompressedTexImageARB"))
	gpGetnTexImage = (C.GPGETNTEXIMAGE)(getProcAddr("glGetnTexImage"))
	if gpGetnTexImage == nil {
		return errors.New("glGetnTexImage")
	}
	gpGetnTexImageARB = (C.GPGETNTEXIMAGEARB)(getProcAddr("glGetnTexImageARB"))
	gpGetnUniformdv = (C.GPGETNUNIFORMDV)(getProcAddr("glGetnUniformdv"))
	if gpGetnUniformdv == nil {
		return errors.New("glGetnUniformdv")
	}
	gpGetnUniformdvARB = (C.GPGETNUNIFORMDVARB)(getProcAddr("glGetnUniformdvARB"))
	gpGetnUniformfv = (C.GPGETNUNIFORMFV)(getProcAddr("glGetnUniformfv"))
	if gpGetnUniformfv == nil {
		return errors.New("glGetnUniformfv")
	}
	gpGetnUniformfvARB = (C.GPGETNUNIFORMFVARB)(getProcAddr("glGetnUniformfvARB"))
	gpGetnUniformfvKHR = (C.GPGETNUNIFORMFVKHR)(getProcAddr("glGetnUniformfvKHR"))
	gpGetnUniformiv = (C.GPGETNUNIFORMIV)(getProcAddr("glGetnUniformiv"))
	if gpGetnUniformiv == nil {
		return errors.New("glGetnUniformiv")
	}
	gpGetnUniformivARB = (C.GPGETNUNIFORMIVARB)(getProcAddr("glGetnUniformivARB"))
	gpGetnUniformivKHR = (C.GPGETNUNIFORMIVKHR)(getProcAddr("glGetnUniformivKHR"))
	gpGetnUniformuiv = (C.GPGETNUNIFORMUIV)(getProcAddr("glGetnUniformuiv"))
	if gpGetnUniformuiv == nil {
		return errors.New("glGetnUniformuiv")
	}
	gpGetnUniformuivARB = (C.GPGETNUNIFORMUIVARB)(getProcAddr("glGetnUniformuivARB"))
	gpGetnUniformuivKHR = (C.GPGETNUNIFORMUIVKHR)(getProcAddr("glGetnUniformuivKHR"))
	gpHint = (C.GPHINT)(getProcAddr("glHint"))
	if gpHint == nil {
		return errors.New("glHint")
	}
	gpInvalidateBufferData = (C.GPINVALIDATEBUFFERDATA)(getProcAddr("glInvalidateBufferData"))
	if gpInvalidateBufferData == nil {
		return errors.New("glInvalidateBufferData")
	}
	gpInvalidateBufferSubData = (C.GPINVALIDATEBUFFERSUBDATA)(getProcAddr("glInvalidateBufferSubData"))
	if gpInvalidateBufferSubData == nil {
		return errors.New("glInvalidateBufferSubData")
	}
	gpInvalidateFramebuffer = (C.GPINVALIDATEFRAMEBUFFER)(getProcAddr("glInvalidateFramebuffer"))
	if gpInvalidateFramebuffer == nil {
		return errors.New("glInvalidateFramebuffer")
	}
	gpInvalidateNamedFramebufferData = (C.GPINVALIDATENAMEDFRAMEBUFFERDATA)(getProcAddr("glInvalidateNamedFramebufferData"))
	if gpInvalidateNamedFramebufferData == nil {
		return errors.New("glInvalidateNamedFramebufferData")
	}
	gpInvalidateNamedFramebufferSubData = (C.GPINVALIDATENAMEDFRAMEBUFFERSUBDATA)(getProcAddr("glInvalidateNamedFramebufferSubData"))
	if gpInvalidateNamedFramebufferSubData == nil {
		return errors.New("glInvalidateNamedFramebufferSubData")
	}
	gpInvalidateSubFramebuffer = (C.GPINVALIDATESUBFRAMEBUFFER)(getProcAddr("glInvalidateSubFramebuffer"))
	if gpInvalidateSubFramebuffer == nil {
		return errors.New("glInvalidateSubFramebuffer")
	}
	gpInvalidateTexImage = (C.GPINVALIDATETEXIMAGE)(getProcAddr("glInvalidateTexImage"))
	if gpInvalidateTexImage == nil {
		return errors.New("glInvalidateTexImage")
	}
	gpInvalidateTexSubImage = (C.GPINVALIDATETEXSUBIMAGE)(getProcAddr("glInvalidateTexSubImage"))
	if gpInvalidateTexSubImage == nil {
		return errors.New("glInvalidateTexSubImage")
	}
	gpIsBuffer = (C.GPISBUFFER)(getProcAddr("glIsBuffer"))
	if gpIsBuffer == nil {
		return errors.New("glIsBuffer")
	}
	gpIsEnabled = (C.GPISENABLED)(getProcAddr("glIsEnabled"))
	if gpIsEnabled == nil {
		return errors.New("glIsEnabled")
	}
	gpIsEnabledi = (C.GPISENABLEDI)(getProcAddr("glIsEnabledi"))
	if gpIsEnabledi == nil {
		return errors.New("glIsEnabledi")
	}
	gpIsFramebuffer = (C.GPISFRAMEBUFFER)(getProcAddr("glIsFramebuffer"))
	if gpIsFramebuffer == nil {
		return errors.New("glIsFramebuffer")
	}
	gpIsImageHandleResidentARB = (C.GPISIMAGEHANDLERESIDENTARB)(getProcAddr("glIsImageHandleResidentARB"))
	gpIsNamedStringARB = (C.GPISNAMEDSTRINGARB)(getProcAddr("glIsNamedStringARB"))
	gpIsProgram = (C.GPISPROGRAM)(getProcAddr("glIsProgram"))
	if gpIsProgram == nil {
		return errors.New("glIsProgram")
	}
	gpIsProgramPipeline = (C.GPISPROGRAMPIPELINE)(getProcAddr("glIsProgramPipeline"))
	if gpIsProgramPipeline == nil {
		return errors.New("glIsProgramPipeline")
	}
	gpIsQuery = (C.GPISQUERY)(getProcAddr("glIsQuery"))
	if gpIsQuery == nil {
		return errors.New("glIsQuery")
	}
	gpIsRenderbuffer = (C.GPISRENDERBUFFER)(getProcAddr("glIsRenderbuffer"))
	if gpIsRenderbuffer == nil {
		return errors.New("glIsRenderbuffer")
	}
	gpIsSampler = (C.GPISSAMPLER)(getProcAddr("glIsSampler"))
	if gpIsSampler == nil {
		return errors.New("glIsSampler")
	}
	gpIsShader = (C.GPISSHADER)(getProcAddr("glIsShader"))
	if gpIsShader == nil {
		return errors.New("glIsShader")
	}
	gpIsSync = (C.GPISSYNC)(getProcAddr("glIsSync"))
	if gpIsSync == nil {
		return errors.New("glIsSync")
	}
	gpIsTexture = (C.GPISTEXTURE)(getProcAddr("glIsTexture"))
	if gpIsTexture == nil {
		return errors.New("glIsTexture")
	}
	gpIsTextureHandleResidentARB = (C.GPISTEXTUREHANDLERESIDENTARB)(getProcAddr("glIsTextureHandleResidentARB"))
	gpIsTransformFeedback = (C.GPISTRANSFORMFEEDBACK)(getProcAddr("glIsTransformFeedback"))
	if gpIsTransformFeedback == nil {
		return errors.New("glIsTransformFeedback")
	}
	gpIsVertexArray = (C.GPISVERTEXARRAY)(getProcAddr("glIsVertexArray"))
	if gpIsVertexArray == nil {
		return errors.New("glIsVertexArray")
	}
	gpLineWidth = (C.GPLINEWIDTH)(getProcAddr("glLineWidth"))
	if gpLineWidth == nil {
		return errors.New("glLineWidth")
	}
	gpLinkProgram = (C.GPLINKPROGRAM)(getProcAddr("glLinkProgram"))
	if gpLinkProgram == nil {
		return errors.New("glLinkProgram")
	}
	gpLogicOp = (C.GPLOGICOP)(getProcAddr("glLogicOp"))
	if gpLogicOp == nil {
		return errors.New("glLogicOp")
	}
	gpMakeImageHandleNonResidentARB = (C.GPMAKEIMAGEHANDLENONRESIDENTARB)(getProcAddr("glMakeImageHandleNonResidentARB"))
	gpMakeImageHandleResidentARB = (C.GPMAKEIMAGEHANDLERESIDENTARB)(getProcAddr("glMakeImageHandleResidentARB"))
	gpMakeTextureHandleNonResidentARB = (C.GPMAKETEXTUREHANDLENONRESIDENTARB)(getProcAddr("glMakeTextureHandleNonResidentARB"))
	gpMakeTextureHandleResidentARB = (C.GPMAKETEXTUREHANDLERESIDENTARB)(getProcAddr("glMakeTextureHandleResidentARB"))
	gpMapBuffer = (C.GPMAPBUFFER)(getProcAddr("glMapBuffer"))
	if gpMapBuffer == nil {
		return errors.New("glMapBuffer")
	}
	gpMapBufferRange = (C.GPMAPBUFFERRANGE)(getProcAddr("glMapBufferRange"))
	if gpMapBufferRange == nil {
		return errors.New("glMapBufferRange")
	}
	gpMapNamedBuffer = (C.GPMAPNAMEDBUFFER)(getProcAddr("glMapNamedBuffer"))
	if gpMapNamedBuffer == nil {
		return errors.New("glMapNamedBuffer")
	}
	gpMapNamedBufferRange = (C.GPMAPNAMEDBUFFERRANGE)(getProcAddr("glMapNamedBufferRange"))
	if gpMapNamedBufferRange == nil {
		return errors.New("glMapNamedBufferRange")
	}
	gpMemoryBarrier = (C.GPMEMORYBARRIER)(getProcAddr("glMemoryBarrier"))
	if gpMemoryBarrier == nil {
		return errors.New("glMemoryBarrier")
	}
	gpMemoryBarrierByRegion = (C.GPMEMORYBARRIERBYREGION)(getProcAddr("glMemoryBarrierByRegion"))
	if gpMemoryBarrierByRegion == nil {
		return errors.New("glMemoryBarrierByRegion")
	}
	gpMinSampleShading = (C.GPMINSAMPLESHADING)(getProcAddr("glMinSampleShading"))
	if gpMinSampleShading == nil {
		return errors.New("glMinSampleShading")
	}
	gpMinSampleShadingARB = (C.GPMINSAMPLESHADINGARB)(getProcAddr("glMinSampleShadingARB"))
	gpMultiDrawArrays = (C.GPMULTIDRAWARRAYS)(getProcAddr("glMultiDrawArrays"))
	if gpMultiDrawArrays == nil {
		return errors.New("glMultiDrawArrays")
	}
	gpMultiDrawArraysIndirect = (C.GPMULTIDRAWARRAYSINDIRECT)(getProcAddr("glMultiDrawArraysIndirect"))
	if gpMultiDrawArraysIndirect == nil {
		return errors.New("glMultiDrawArraysIndirect")
	}
	gpMultiDrawArraysIndirectCountARB = (C.GPMULTIDRAWARRAYSINDIRECTCOUNTARB)(getProcAddr("glMultiDrawArraysIndirectCountARB"))
	gpMultiDrawElements = (C.GPMULTIDRAWELEMENTS)(getProcAddr("glMultiDrawElements"))
	if gpMultiDrawElements == nil {
		return errors.New("glMultiDrawElements")
	}
	gpMultiDrawElementsBaseVertex = (C.GPMULTIDRAWELEMENTSBASEVERTEX)(getProcAddr("glMultiDrawElementsBaseVertex"))
	if gpMultiDrawElementsBaseVertex == nil {
		return errors.New("glMultiDrawElementsBaseVertex")
	}
	gpMultiDrawElementsIndirect = (C.GPMULTIDRAWELEMENTSINDIRECT)(getProcAddr("glMultiDrawElementsIndirect"))
	if gpMultiDrawElementsIndirect == nil {
		return errors.New("glMultiDrawElementsIndirect")
	}
	gpMultiDrawElementsIndirectCountARB = (C.GPMULTIDRAWELEMENTSINDIRECTCOUNTARB)(getProcAddr("glMultiDrawElementsIndirectCountARB"))
	gpNamedBufferData = (C.GPNAMEDBUFFERDATA)(getProcAddr("glNamedBufferData"))
	if gpNamedBufferData == nil {
		return errors.New("glNamedBufferData")
	}
	gpNamedBufferPageCommitmentARB = (C.GPNAMEDBUFFERPAGECOMMITMENTARB)(getProcAddr("glNamedBufferPageCommitmentARB"))
	gpNamedBufferPageCommitmentEXT = (C.GPNAMEDBUFFERPAGECOMMITMENTEXT)(getProcAddr("glNamedBufferPageCommitmentEXT"))
	gpNamedBufferStorage = (C.GPNAMEDBUFFERSTORAGE)(getProcAddr("glNamedBufferStorage"))
	if gpNamedBufferStorage == nil {
		return errors.New("glNamedBufferStorage")
	}
	gpNamedBufferSubData = (C.GPNAMEDBUFFERSUBDATA)(getProcAddr("glNamedBufferSubData"))
	if gpNamedBufferSubData == nil {
		return errors.New("glNamedBufferSubData")
	}
	gpNamedFramebufferDrawBuffer = (C.GPNAMEDFRAMEBUFFERDRAWBUFFER)(getProcAddr("glNamedFramebufferDrawBuffer"))
	if gpNamedFramebufferDrawBuffer == nil {
		return errors.New("glNamedFramebufferDrawBuffer")
	}
	gpNamedFramebufferDrawBuffers = (C.GPNAMEDFRAMEBUFFERDRAWBUFFERS)(getProcAddr("glNamedFramebufferDrawBuffers"))
	if gpNamedFramebufferDrawBuffers == nil {
		return errors.New("glNamedFramebufferDrawBuffers")
	}
	gpNamedFramebufferParameteri = (C.GPNAMEDFRAMEBUFFERPARAMETERI)(getProcAddr("glNamedFramebufferParameteri"))
	if gpNamedFramebufferParameteri == nil {
		return errors.New("glNamedFramebufferParameteri")
	}
	gpNamedFramebufferReadBuffer = (C.GPNAMEDFRAMEBUFFERREADBUFFER)(getProcAddr("glNamedFramebufferReadBuffer"))
	if gpNamedFramebufferReadBuffer == nil {
		return errors.New("glNamedFramebufferReadBuffer")
	}
	gpNamedFramebufferRenderbuffer = (C.GPNAMEDFRAMEBUFFERRENDERBUFFER)(getProcAddr("glNamedFramebufferRenderbuffer"))
	if gpNamedFramebufferRenderbuffer == nil {
		return errors.New("glNamedFramebufferRenderbuffer")
	}
	gpNamedFramebufferTexture = (C.GPNAMEDFRAMEBUFFERTEXTURE)(getProcAddr("glNamedFramebufferTexture"))
	if gpNamedFramebufferTexture == nil {
		return errors.New("glNamedFramebufferTexture")
	}
	gpNamedFramebufferTextureLayer = (C.GPNAMEDFRAMEBUFFERTEXTURELAYER)(getProcAddr("glNamedFramebufferTextureLayer"))
	if gpNamedFramebufferTextureLayer == nil {
		return errors.New("glNamedFramebufferTextureLayer")
	}
	gpNamedRenderbufferStorage = (C.GPNAMEDRENDERBUFFERSTORAGE)(getProcAddr("glNamedRenderbufferStorage"))
	if gpNamedRenderbufferStorage == nil {
		return errors.New("glNamedRenderbufferStorage")
	}
	gpNamedRenderbufferStorageMultisample = (C.GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE)(getProcAddr("glNamedRenderbufferStorageMultisample"))
	if gpNamedRenderbufferStorageMultisample == nil {
		return errors.New("glNamedRenderbufferStorageMultisample")
	}
	gpNamedStringARB = (C.GPNAMEDSTRINGARB)(getProcAddr("glNamedStringARB"))
	gpObjectLabel = (C.GPOBJECTLABEL)(getProcAddr("glObjectLabel"))
	if gpObjectLabel == nil {
		return errors.New("glObjectLabel")
	}
	gpObjectLabelKHR = (C.GPOBJECTLABELKHR)(getProcAddr("glObjectLabelKHR"))
	gpObjectPtrLabel = (C.GPOBJECTPTRLABEL)(getProcAddr("glObjectPtrLabel"))
	if gpObjectPtrLabel == nil {
		return errors.New("glObjectPtrLabel")
	}
	gpObjectPtrLabelKHR = (C.GPOBJECTPTRLABELKHR)(getProcAddr("glObjectPtrLabelKHR"))
	gpPatchParameterfv = (C.GPPATCHPARAMETERFV)(getProcAddr("glPatchParameterfv"))
	if gpPatchParameterfv == nil {
		return errors.New("glPatchParameterfv")
	}
	gpPatchParameteri = (C.GPPATCHPARAMETERI)(getProcAddr("glPatchParameteri"))
	if gpPatchParameteri == nil {
		return errors.New("glPatchParameteri")
	}
	gpPauseTransformFeedback = (C.GPPAUSETRANSFORMFEEDBACK)(getProcAddr("glPauseTransformFeedback"))
	if gpPauseTransformFeedback == nil {
		return errors.New("glPauseTransformFeedback")
	}
	gpPixelStoref = (C.GPPIXELSTOREF)(getProcAddr("glPixelStoref"))
	if gpPixelStoref == nil {
		return errors.New("glPixelStoref")
	}
	gpPixelStorei = (C.GPPIXELSTOREI)(getProcAddr("glPixelStorei"))
	if gpPixelStorei == nil {
		return errors.New("glPixelStorei")
	}
	gpPointParameterf = (C.GPPOINTPARAMETERF)(getProcAddr("glPointParameterf"))
	if gpPointParameterf == nil {
		return errors.New("glPointParameterf")
	}
	gpPointParameterfv = (C.GPPOINTPARAMETERFV)(getProcAddr("glPointParameterfv"))
	if gpPointParameterfv == nil {
		return errors.New("glPointParameterfv")
	}
	gpPointParameteri = (C.GPPOINTPARAMETERI)(getProcAddr("glPointParameteri"))
	if gpPointParameteri == nil {
		return errors.New("glPointParameteri")
	}
	gpPointParameteriv = (C.GPPOINTPARAMETERIV)(getProcAddr("glPointParameteriv"))
	if gpPointParameteriv == nil {
		return errors.New("glPointParameteriv")
	}
	gpPointSize = (C.GPPOINTSIZE)(getProcAddr("glPointSize"))
	if gpPointSize == nil {
		return errors.New("glPointSize")
	}
	gpPolygonMode = (C.GPPOLYGONMODE)(getProcAddr("glPolygonMode"))
	if gpPolygonMode == nil {
		return errors.New("glPolygonMode")
	}
	gpPolygonOffset = (C.GPPOLYGONOFFSET)(getProcAddr("glPolygonOffset"))
	if gpPolygonOffset == nil {
		return errors.New("glPolygonOffset")
	}
	gpPopDebugGroup = (C.GPPOPDEBUGGROUP)(getProcAddr("glPopDebugGroup"))
	if gpPopDebugGroup == nil {
		return errors.New("glPopDebugGroup")
	}
	gpPopDebugGroupKHR = (C.GPPOPDEBUGGROUPKHR)(getProcAddr("glPopDebugGroupKHR"))
	gpPrimitiveRestartIndex = (C.GPPRIMITIVERESTARTINDEX)(getProcAddr("glPrimitiveRestartIndex"))
	if gpPrimitiveRestartIndex == nil {
		return errors.New("glPrimitiveRestartIndex")
	}
	gpProgramBinary = (C.GPPROGRAMBINARY)(getProcAddr("glProgramBinary"))
	if gpProgramBinary == nil {
		return errors.New("glProgramBinary")
	}
	gpProgramParameteri = (C.GPPROGRAMPARAMETERI)(getProcAddr("glProgramParameteri"))
	if gpProgramParameteri == nil {
		return errors.New("glProgramParameteri")
	}
	gpProgramUniform1d = (C.GPPROGRAMUNIFORM1D)(getProcAddr("glProgramUniform1d"))
	if gpProgramUniform1d == nil {
		return errors.New("glProgramUniform1d")
	}
	gpProgramUniform1dv = (C.GPPROGRAMUNIFORM1DV)(getProcAddr("glProgramUniform1dv"))
	if gpProgramUniform1dv == nil {
		return errors.New("glProgramUniform1dv")
	}
	gpProgramUniform1f = (C.GPPROGRAMUNIFORM1F)(getProcAddr("glProgramUniform1f"))
	if gpProgramUniform1f == nil {
		return errors.New("glProgramUniform1f")
	}
	gpProgramUniform1fv = (C.GPPROGRAMUNIFORM1FV)(getProcAddr("glProgramUniform1fv"))
	if gpProgramUniform1fv == nil {
		return errors.New("glProgramUniform1fv")
	}
	gpProgramUniform1i = (C.GPPROGRAMUNIFORM1I)(getProcAddr("glProgramUniform1i"))
	if gpProgramUniform1i == nil {
		return errors.New("glProgramUniform1i")
	}
	gpProgramUniform1iv = (C.GPPROGRAMUNIFORM1IV)(getProcAddr("glProgramUniform1iv"))
	if gpProgramUniform1iv == nil {
		return errors.New("glProgramUniform1iv")
	}
	gpProgramUniform1ui = (C.GPPROGRAMUNIFORM1UI)(getProcAddr("glProgramUniform1ui"))
	if gpProgramUniform1ui == nil {
		return errors.New("glProgramUniform1ui")
	}
	gpProgramUniform1uiv = (C.GPPROGRAMUNIFORM1UIV)(getProcAddr("glProgramUniform1uiv"))
	if gpProgramUniform1uiv == nil {
		return errors.New("glProgramUniform1uiv")
	}
	gpProgramUniform2d = (C.GPPROGRAMUNIFORM2D)(getProcAddr("glProgramUniform2d"))
	if gpProgramUniform2d == nil {
		return errors.New("glProgramUniform2d")
	}
	gpProgramUniform2dv = (C.GPPROGRAMUNIFORM2DV)(getProcAddr("glProgramUniform2dv"))
	if gpProgramUniform2dv == nil {
		return errors.New("glProgramUniform2dv")
	}
	gpProgramUniform2f = (C.GPPROGRAMUNIFORM2F)(getProcAddr("glProgramUniform2f"))
	if gpProgramUniform2f == nil {
		return errors.New("glProgramUniform2f")
	}
	gpProgramUniform2fv = (C.GPPROGRAMUNIFORM2FV)(getProcAddr("glProgramUniform2fv"))
	if gpProgramUniform2fv == nil {
		return errors.New("glProgramUniform2fv")
	}
	gpProgramUniform2i = (C.GPPROGRAMUNIFORM2I)(getProcAddr("glProgramUniform2i"))
	if gpProgramUniform2i == nil {
		return errors.New("glProgramUniform2i")
	}
	gpProgramUniform2iv = (C.GPPROGRAMUNIFORM2IV)(getProcAddr("glProgramUniform2iv"))
	if gpProgramUniform2iv == nil {
		return errors.New("glProgramUniform2iv")
	}
	gpProgramUniform2ui = (C.GPPROGRAMUNIFORM2UI)(getProcAddr("glProgramUniform2ui"))
	if gpProgramUniform2ui == nil {
		return errors.New("glProgramUniform2ui")
	}
	gpProgramUniform2uiv = (C.GPPROGRAMUNIFORM2UIV)(getProcAddr("glProgramUniform2uiv"))
	if gpProgramUniform2uiv == nil {
		return errors.New("glProgramUniform2uiv")
	}
	gpProgramUniform3d = (C.GPPROGRAMUNIFORM3D)(getProcAddr("glProgramUniform3d"))
	if gpProgramUniform3d == nil {
		return errors.New("glProgramUniform3d")
	}
	gpProgramUniform3dv = (C.GPPROGRAMUNIFORM3DV)(getProcAddr("glProgramUniform3dv"))
	if gpProgramUniform3dv == nil {
		return errors.New("glProgramUniform3dv")
	}
	gpProgramUniform3f = (C.GPPROGRAMUNIFORM3F)(getProcAddr("glProgramUniform3f"))
	if gpProgramUniform3f == nil {
		return errors.New("glProgramUniform3f")
	}
	gpProgramUniform3fv = (C.GPPROGRAMUNIFORM3FV)(getProcAddr("glProgramUniform3fv"))
	if gpProgramUniform3fv == nil {
		return errors.New("glProgramUniform3fv")
	}
	gpProgramUniform3i = (C.GPPROGRAMUNIFORM3I)(getProcAddr("glProgramUniform3i"))
	if gpProgramUniform3i == nil {
		return errors.New("glProgramUniform3i")
	}
	gpProgramUniform3iv = (C.GPPROGRAMUNIFORM3IV)(getProcAddr("glProgramUniform3iv"))
	if gpProgramUniform3iv == nil {
		return errors.New("glProgramUniform3iv")
	}
	gpProgramUniform3ui = (C.GPPROGRAMUNIFORM3UI)(getProcAddr("glProgramUniform3ui"))
	if gpProgramUniform3ui == nil {
		return errors.New("glProgramUniform3ui")
	}
	gpProgramUniform3uiv = (C.GPPROGRAMUNIFORM3UIV)(getProcAddr("glProgramUniform3uiv"))
	if gpProgramUniform3uiv == nil {
		return errors.New("glProgramUniform3uiv")
	}
	gpProgramUniform4d = (C.GPPROGRAMUNIFORM4D)(getProcAddr("glProgramUniform4d"))
	if gpProgramUniform4d == nil {
		return errors.New("glProgramUniform4d")
	}
	gpProgramUniform4dv = (C.GPPROGRAMUNIFORM4DV)(getProcAddr("glProgramUniform4dv"))
	if gpProgramUniform4dv == nil {
		return errors.New("glProgramUniform4dv")
	}
	gpProgramUniform4f = (C.GPPROGRAMUNIFORM4F)(getProcAddr("glProgramUniform4f"))
	if gpProgramUniform4f == nil {
		return errors.New("glProgramUniform4f")
	}
	gpProgramUniform4fv = (C.GPPROGRAMUNIFORM4FV)(getProcAddr("glProgramUniform4fv"))
	if gpProgramUniform4fv == nil {
		return errors.New("glProgramUniform4fv")
	}
	gpProgramUniform4i = (C.GPPROGRAMUNIFORM4I)(getProcAddr("glProgramUniform4i"))
	if gpProgramUniform4i == nil {
		return errors.New("glProgramUniform4i")
	}
	gpProgramUniform4iv = (C.GPPROGRAMUNIFORM4IV)(getProcAddr("glProgramUniform4iv"))
	if gpProgramUniform4iv == nil {
		return errors.New("glProgramUniform4iv")
	}
	gpProgramUniform4ui = (C.GPPROGRAMUNIFORM4UI)(getProcAddr("glProgramUniform4ui"))
	if gpProgramUniform4ui == nil {
		return errors.New("glProgramUniform4ui")
	}
	gpProgramUniform4uiv = (C.GPPROGRAMUNIFORM4UIV)(getProcAddr("glProgramUniform4uiv"))
	if gpProgramUniform4uiv == nil {
		return errors.New("glProgramUniform4uiv")
	}
	gpProgramUniformHandleui64ARB = (C.GPPROGRAMUNIFORMHANDLEUI64ARB)(getProcAddr("glProgramUniformHandleui64ARB"))
	gpProgramUniformHandleui64vARB = (C.GPPROGRAMUNIFORMHANDLEUI64VARB)(getProcAddr("glProgramUniformHandleui64vARB"))
	gpProgramUniformMatrix2dv = (C.GPPROGRAMUNIFORMMATRIX2DV)(getProcAddr("glProgramUniformMatrix2dv"))
	if gpProgramUniformMatrix2dv == nil {
		return errors.New("glProgramUniformMatrix2dv")
	}
	gpProgramUniformMatrix2fv = (C.GPPROGRAMUNIFORMMATRIX2FV)(getProcAddr("glProgramUniformMatrix2fv"))
	if gpProgramUniformMatrix2fv == nil {
		return errors.New("glProgramUniformMatrix2fv")
	}
	gpProgramUniformMatrix2x3dv = (C.GPPROGRAMUNIFORMMATRIX2X3DV)(getProcAddr("glProgramUniformMatrix2x3dv"))
	if gpProgramUniformMatrix2x3dv == nil {
		return errors.New("glProgramUniformMatrix2x3dv")
	}
	gpProgramUniformMatrix2x3fv = (C.GPPROGRAMUNIFORMMATRIX2X3FV)(getProcAddr("glProgramUniformMatrix2x3fv"))
	if gpProgramUniformMatrix2x3fv == nil {
		return errors.New("glProgramUniformMatrix2x3fv")
	}
	gpProgramUniformMatrix2x4dv = (C.GPPROGRAMUNIFORMMATRIX2X4DV)(getProcAddr("glProgramUniformMatrix2x4dv"))
	if gpProgramUniformMatrix2x4dv == nil {
		return errors.New("glProgramUniformMatrix2x4dv")
	}
	gpProgramUniformMatrix2x4fv = (C.GPPROGRAMUNIFORMMATRIX2X4FV)(getProcAddr("glProgramUniformMatrix2x4fv"))
	if gpProgramUniformMatrix2x4fv == nil {
		return errors.New("glProgramUniformMatrix2x4fv")
	}
	gpProgramUniformMatrix3dv = (C.GPPROGRAMUNIFORMMATRIX3DV)(getProcAddr("glProgramUniformMatrix3dv"))
	if gpProgramUniformMatrix3dv == nil {
		return errors.New("glProgramUniformMatrix3dv")
	}
	gpProgramUniformMatrix3fv = (C.GPPROGRAMUNIFORMMATRIX3FV)(getProcAddr("glProgramUniformMatrix3fv"))
	if gpProgramUniformMatrix3fv == nil {
		return errors.New("glProgramUniformMatrix3fv")
	}
	gpProgramUniformMatrix3x2dv = (C.GPPROGRAMUNIFORMMATRIX3X2DV)(getProcAddr("glProgramUniformMatrix3x2dv"))
	if gpProgramUniformMatrix3x2dv == nil {
		return errors.New("glProgramUniformMatrix3x2dv")
	}
	gpProgramUniformMatrix3x2fv = (C.GPPROGRAMUNIFORMMATRIX3X2FV)(getProcAddr("glProgramUniformMatrix3x2fv"))
	if gpProgramUniformMatrix3x2fv == nil {
		return errors.New("glProgramUniformMatrix3x2fv")
	}
	gpProgramUniformMatrix3x4dv = (C.GPPROGRAMUNIFORMMATRIX3X4DV)(getProcAddr("glProgramUniformMatrix3x4dv"))
	if gpProgramUniformMatrix3x4dv == nil {
		return errors.New("glProgramUniformMatrix3x4dv")
	}
	gpProgramUniformMatrix3x4fv = (C.GPPROGRAMUNIFORMMATRIX3X4FV)(getProcAddr("glProgramUniformMatrix3x4fv"))
	if gpProgramUniformMatrix3x4fv == nil {
		return errors.New("glProgramUniformMatrix3x4fv")
	}
	gpProgramUniformMatrix4dv = (C.GPPROGRAMUNIFORMMATRIX4DV)(getProcAddr("glProgramUniformMatrix4dv"))
	if gpProgramUniformMatrix4dv == nil {
		return errors.New("glProgramUniformMatrix4dv")
	}
	gpProgramUniformMatrix4fv = (C.GPPROGRAMUNIFORMMATRIX4FV)(getProcAddr("glProgramUniformMatrix4fv"))
	if gpProgramUniformMatrix4fv == nil {
		return errors.New("glProgramUniformMatrix4fv")
	}
	gpProgramUniformMatrix4x2dv = (C.GPPROGRAMUNIFORMMATRIX4X2DV)(getProcAddr("glProgramUniformMatrix4x2dv"))
	if gpProgramUniformMatrix4x2dv == nil {
		return errors.New("glProgramUniformMatrix4x2dv")
	}
	gpProgramUniformMatrix4x2fv = (C.GPPROGRAMUNIFORMMATRIX4X2FV)(getProcAddr("glProgramUniformMatrix4x2fv"))
	if gpProgramUniformMatrix4x2fv == nil {
		return errors.New("glProgramUniformMatrix4x2fv")
	}
	gpProgramUniformMatrix4x3dv = (C.GPPROGRAMUNIFORMMATRIX4X3DV)(getProcAddr("glProgramUniformMatrix4x3dv"))
	if gpProgramUniformMatrix4x3dv == nil {
		return errors.New("glProgramUniformMatrix4x3dv")
	}
	gpProgramUniformMatrix4x3fv = (C.GPPROGRAMUNIFORMMATRIX4X3FV)(getProcAddr("glProgramUniformMatrix4x3fv"))
	if gpProgramUniformMatrix4x3fv == nil {
		return errors.New("glProgramUniformMatrix4x3fv")
	}
	gpProvokingVertex = (C.GPPROVOKINGVERTEX)(getProcAddr("glProvokingVertex"))
	if gpProvokingVertex == nil {
		return errors.New("glProvokingVertex")
	}
	gpPushDebugGroup = (C.GPPUSHDEBUGGROUP)(getProcAddr("glPushDebugGroup"))
	if gpPushDebugGroup == nil {
		return errors.New("glPushDebugGroup")
	}
	gpPushDebugGroupKHR = (C.GPPUSHDEBUGGROUPKHR)(getProcAddr("glPushDebugGroupKHR"))
	gpQueryCounter = (C.GPQUERYCOUNTER)(getProcAddr("glQueryCounter"))
	if gpQueryCounter == nil {
		return errors.New("glQueryCounter")
	}
	gpReadBuffer = (C.GPREADBUFFER)(getProcAddr("glReadBuffer"))
	if gpReadBuffer == nil {
		return errors.New("glReadBuffer")
	}
	gpReadPixels = (C.GPREADPIXELS)(getProcAddr("glReadPixels"))
	if gpReadPixels == nil {
		return errors.New("glReadPixels")
	}
	gpReadnPixels = (C.GPREADNPIXELS)(getProcAddr("glReadnPixels"))
	if gpReadnPixels == nil {
		return errors.New("glReadnPixels")
	}
	gpReadnPixelsARB = (C.GPREADNPIXELSARB)(getProcAddr("glReadnPixelsARB"))
	gpReadnPixelsKHR = (C.GPREADNPIXELSKHR)(getProcAddr("glReadnPixelsKHR"))
	gpReleaseShaderCompiler = (C.GPRELEASESHADERCOMPILER)(getProcAddr("glReleaseShaderCompiler"))
	if gpReleaseShaderCompiler == nil {
		return errors.New("glReleaseShaderCompiler")
	}
	gpRenderbufferStorage = (C.GPRENDERBUFFERSTORAGE)(getProcAddr("glRenderbufferStorage"))
	if gpRenderbufferStorage == nil {
		return errors.New("glRenderbufferStorage")
	}
	gpRenderbufferStorageMultisample = (C.GPRENDERBUFFERSTORAGEMULTISAMPLE)(getProcAddr("glRenderbufferStorageMultisample"))
	if gpRenderbufferStorageMultisample == nil {
		return errors.New("glRenderbufferStorageMultisample")
	}
	gpResumeTransformFeedback = (C.GPRESUMETRANSFORMFEEDBACK)(getProcAddr("glResumeTransformFeedback"))
	if gpResumeTransformFeedback == nil {
		return errors.New("glResumeTransformFeedback")
	}
	gpSampleCoverage = (C.GPSAMPLECOVERAGE)(getProcAddr("glSampleCoverage"))
	if gpSampleCoverage == nil {
		return errors.New("glSampleCoverage")
	}
	gpSampleMaski = (C.GPSAMPLEMASKI)(getProcAddr("glSampleMaski"))
	if gpSampleMaski == nil {
		return errors.New("glSampleMaski")
	}
	gpSamplerParameterIiv = (C.GPSAMPLERPARAMETERIIV)(getProcAddr("glSamplerParameterIiv"))
	if gpSamplerParameterIiv == nil {
		return errors.New("glSamplerParameterIiv")
	}
	gpSamplerParameterIuiv = (C.GPSAMPLERPARAMETERIUIV)(getProcAddr("glSamplerParameterIuiv"))
	if gpSamplerParameterIuiv == nil {
		return errors.New("glSamplerParameterIuiv")
	}
	gpSamplerParameterf = (C.GPSAMPLERPARAMETERF)(getProcAddr("glSamplerParameterf"))
	if gpSamplerParameterf == nil {
		return errors.New("glSamplerParameterf")
	}
	gpSamplerParameterfv = (C.GPSAMPLERPARAMETERFV)(getProcAddr("glSamplerParameterfv"))
	if gpSamplerParameterfv == nil {
		return errors.New("glSamplerParameterfv")
	}
	gpSamplerParameteri = (C.GPSAMPLERPARAMETERI)(getProcAddr("glSamplerParameteri"))
	if gpSamplerParameteri == nil {
		return errors.New("glSamplerParameteri")
	}
	gpSamplerParameteriv = (C.GPSAMPLERPARAMETERIV)(getProcAddr("glSamplerParameteriv"))
	if gpSamplerParameteriv == nil {
		return errors.New("glSamplerParameteriv")
	}
	gpScissor = (C.GPSCISSOR)(getProcAddr("glScissor"))
	if gpScissor == nil {
		return errors.New("glScissor")
	}
	gpScissorArrayv = (C.GPSCISSORARRAYV)(getProcAddr("glScissorArrayv"))
	if gpScissorArrayv == nil {
		return errors.New("glScissorArrayv")
	}
	gpScissorIndexed = (C.GPSCISSORINDEXED)(getProcAddr("glScissorIndexed"))
	if gpScissorIndexed == nil {
		return errors.New("glScissorIndexed")
	}
	gpScissorIndexedv = (C.GPSCISSORINDEXEDV)(getProcAddr("glScissorIndexedv"))
	if gpScissorIndexedv == nil {
		return errors.New("glScissorIndexedv")
	}
	gpShaderBinary = (C.GPSHADERBINARY)(getProcAddr("glShaderBinary"))
	if gpShaderBinary == nil {
		return errors.New("glShaderBinary")
	}
	gpShaderSource = (C.GPSHADERSOURCE)(getProcAddr("glShaderSource"))
	if gpShaderSource == nil {
		return errors.New("glShaderSource")
	}
	gpShaderStorageBlockBinding = (C.GPSHADERSTORAGEBLOCKBINDING)(getProcAddr("glShaderStorageBlockBinding"))
	if gpShaderStorageBlockBinding == nil {
		return errors.New("glShaderStorageBlockBinding")
	}
	gpStencilFunc = (C.GPSTENCILFUNC)(getProcAddr("glStencilFunc"))
	if gpStencilFunc == nil {
		return errors.New("glStencilFunc")
	}
	gpStencilFuncSeparate = (C.GPSTENCILFUNCSEPARATE)(getProcAddr("glStencilFuncSeparate"))
	if gpStencilFuncSeparate == nil {
		return errors.New("glStencilFuncSeparate")
	}
	gpStencilMask = (C.GPSTENCILMASK)(getProcAddr("glStencilMask"))
	if gpStencilMask == nil {
		return errors.New("glStencilMask")
	}
	gpStencilMaskSeparate = (C.GPSTENCILMASKSEPARATE)(getProcAddr("glStencilMaskSeparate"))
	if gpStencilMaskSeparate == nil {
		return errors.New("glStencilMaskSeparate")
	}
	gpStencilOp = (C.GPSTENCILOP)(getProcAddr("glStencilOp"))
	if gpStencilOp == nil {
		return errors.New("glStencilOp")
	}
	gpStencilOpSeparate = (C.GPSTENCILOPSEPARATE)(getProcAddr("glStencilOpSeparate"))
	if gpStencilOpSeparate == nil {
		return errors.New("glStencilOpSeparate")
	}
	gpTexBuffer = (C.GPTEXBUFFER)(getProcAddr("glTexBuffer"))
	if gpTexBuffer == nil {
		return errors.New("glTexBuffer")
	}
	gpTexBufferRange = (C.GPTEXBUFFERRANGE)(getProcAddr("glTexBufferRange"))
	if gpTexBufferRange == nil {
		return errors.New("glTexBufferRange")
	}
	gpTexImage1D = (C.GPTEXIMAGE1D)(getProcAddr("glTexImage1D"))
	if gpTexImage1D == nil {
		return errors.New("glTexImage1D")
	}
	gpTexImage2D = (C.GPTEXIMAGE2D)(getProcAddr("glTexImage2D"))
	if gpTexImage2D == nil {
		return errors.New("glTexImage2D")
	}
	gpTexImage2DMultisample = (C.GPTEXIMAGE2DMULTISAMPLE)(getProcAddr("glTexImage2DMultisample"))
	if gpTexImage2DMultisample == nil {
		return errors.New("glTexImage2DMultisample")
	}
	gpTexImage3D = (C.GPTEXIMAGE3D)(getProcAddr("glTexImage3D"))
	if gpTexImage3D == nil {
		return errors.New("glTexImage3D")
	}
	gpTexImage3DMultisample = (C.GPTEXIMAGE3DMULTISAMPLE)(getProcAddr("glTexImage3DMultisample"))
	if gpTexImage3DMultisample == nil {
		return errors.New("glTexImage3DMultisample")
	}
	gpTexPageCommitmentARB = (C.GPTEXPAGECOMMITMENTARB)(getProcAddr("glTexPageCommitmentARB"))
	gpTexParameterIiv = (C.GPTEXPARAMETERIIV)(getProcAddr("glTexParameterIiv"))
	if gpTexParameterIiv == nil {
		return errors.New("glTexParameterIiv")
	}
	gpTexParameterIuiv = (C.GPTEXPARAMETERIUIV)(getProcAddr("glTexParameterIuiv"))
	if gpTexParameterIuiv == nil {
		return errors.New("glTexParameterIuiv")
	}
	gpTexParameterf = (C.GPTEXPARAMETERF)(getProcAddr("glTexParameterf"))
	if gpTexParameterf == nil {
		return errors.New("glTexParameterf")
	}
	gpTexParameterfv = (C.GPTEXPARAMETERFV)(getProcAddr("glTexParameterfv"))
	if gpTexParameterfv == nil {
		return errors.New("glTexParameterfv")
	}
	gpTexParameteri = (C.GPTEXPARAMETERI)(getProcAddr("glTexParameteri"))
	if gpTexParameteri == nil {
		return errors.New("glTexParameteri")
	}
	gpTexParameteriv = (C.GPTEXPARAMETERIV)(getProcAddr("glTexParameteriv"))
	if gpTexParameteriv == nil {
		return errors.New("glTexParameteriv")
	}
	gpTexStorage1D = (C.GPTEXSTORAGE1D)(getProcAddr("glTexStorage1D"))
	if gpTexStorage1D == nil {
		return errors.New("glTexStorage1D")
	}
	gpTexStorage2D = (C.GPTEXSTORAGE2D)(getProcAddr("glTexStorage2D"))
	if gpTexStorage2D == nil {
		return errors.New("glTexStorage2D")
	}
	gpTexStorage2DMultisample = (C.GPTEXSTORAGE2DMULTISAMPLE)(getProcAddr("glTexStorage2DMultisample"))
	if gpTexStorage2DMultisample == nil {
		return errors.New("glTexStorage2DMultisample")
	}
	gpTexStorage3D = (C.GPTEXSTORAGE3D)(getProcAddr("glTexStorage3D"))
	if gpTexStorage3D == nil {
		return errors.New("glTexStorage3D")
	}
	gpTexStorage3DMultisample = (C.GPTEXSTORAGE3DMULTISAMPLE)(getProcAddr("glTexStorage3DMultisample"))
	if gpTexStorage3DMultisample == nil {
		return errors.New("glTexStorage3DMultisample")
	}
	gpTexSubImage1D = (C.GPTEXSUBIMAGE1D)(getProcAddr("glTexSubImage1D"))
	if gpTexSubImage1D == nil {
		return errors.New("glTexSubImage1D")
	}
	gpTexSubImage2D = (C.GPTEXSUBIMAGE2D)(getProcAddr("glTexSubImage2D"))
	if gpTexSubImage2D == nil {
		return errors.New("glTexSubImage2D")
	}
	gpTexSubImage3D = (C.GPTEXSUBIMAGE3D)(getProcAddr("glTexSubImage3D"))
	if gpTexSubImage3D == nil {
		return errors.New("glTexSubImage3D")
	}
	gpTextureBarrier = (C.GPTEXTUREBARRIER)(getProcAddr("glTextureBarrier"))
	if gpTextureBarrier == nil {
		return errors.New("glTextureBarrier")
	}
	gpTextureBuffer = (C.GPTEXTUREBUFFER)(getProcAddr("glTextureBuffer"))
	if gpTextureBuffer == nil {
		return errors.New("glTextureBuffer")
	}
	gpTextureBufferRange = (C.GPTEXTUREBUFFERRANGE)(getProcAddr("glTextureBufferRange"))
	if gpTextureBufferRange == nil {
		return errors.New("glTextureBufferRange")
	}
	gpTextureParameterIiv = (C.GPTEXTUREPARAMETERIIV)(getProcAddr("glTextureParameterIiv"))
	if gpTextureParameterIiv == nil {
		return errors.New("glTextureParameterIiv")
	}
	gpTextureParameterIuiv = (C.GPTEXTUREPARAMETERIUIV)(getProcAddr("glTextureParameterIuiv"))
	if gpTextureParameterIuiv == nil {
		return errors.New("glTextureParameterIuiv")
	}
	gpTextureParameterf = (C.GPTEXTUREPARAMETERF)(getProcAddr("glTextureParameterf"))
	if gpTextureParameterf == nil {
		return errors.New("glTextureParameterf")
	}
	gpTextureParameterfv = (C.GPTEXTUREPARAMETERFV)(getProcAddr("glTextureParameterfv"))
	if gpTextureParameterfv == nil {
		return errors.New("glTextureParameterfv")
	}
	gpTextureParameteri = (C.GPTEXTUREPARAMETERI)(getProcAddr("glTextureParameteri"))
	if gpTextureParameteri == nil {
		return errors.New("glTextureParameteri")
	}
	gpTextureParameteriv = (C.GPTEXTUREPARAMETERIV)(getProcAddr("glTextureParameteriv"))
	if gpTextureParameteriv == nil {
		return errors.New("glTextureParameteriv")
	}
	gpTextureStorage1D = (C.GPTEXTURESTORAGE1D)(getProcAddr("glTextureStorage1D"))
	if gpTextureStorage1D == nil {
		return errors.New("glTextureStorage1D")
	}
	gpTextureStorage2D = (C.GPTEXTURESTORAGE2D)(getProcAddr("glTextureStorage2D"))
	if gpTextureStorage2D == nil {
		return errors.New("glTextureStorage2D")
	}
	gpTextureStorage2DMultisample = (C.GPTEXTURESTORAGE2DMULTISAMPLE)(getProcAddr("glTextureStorage2DMultisample"))
	if gpTextureStorage2DMultisample == nil {
		return errors.New("glTextureStorage2DMultisample")
	}
	gpTextureStorage3D = (C.GPTEXTURESTORAGE3D)(getProcAddr("glTextureStorage3D"))
	if gpTextureStorage3D == nil {
		return errors.New("glTextureStorage3D")
	}
	gpTextureStorage3DMultisample = (C.GPTEXTURESTORAGE3DMULTISAMPLE)(getProcAddr("glTextureStorage3DMultisample"))
	if gpTextureStorage3DMultisample == nil {
		return errors.New("glTextureStorage3DMultisample")
	}
	gpTextureSubImage1D = (C.GPTEXTURESUBIMAGE1D)(getProcAddr("glTextureSubImage1D"))
	if gpTextureSubImage1D == nil {
		return errors.New("glTextureSubImage1D")
	}
	gpTextureSubImage2D = (C.GPTEXTURESUBIMAGE2D)(getProcAddr("glTextureSubImage2D"))
	if gpTextureSubImage2D == nil {
		return errors.New("glTextureSubImage2D")
	}
	gpTextureSubImage3D = (C.GPTEXTURESUBIMAGE3D)(getProcAddr("glTextureSubImage3D"))
	if gpTextureSubImage3D == nil {
		return errors.New("glTextureSubImage3D")
	}
	gpTextureView = (C.GPTEXTUREVIEW)(getProcAddr("glTextureView"))
	if gpTextureView == nil {
		return errors.New("glTextureView")
	}
	gpTransformFeedbackBufferBase = (C.GPTRANSFORMFEEDBACKBUFFERBASE)(getProcAddr("glTransformFeedbackBufferBase"))
	if gpTransformFeedbackBufferBase == nil {
		return errors.New("glTransformFeedbackBufferBase")
	}
	gpTransformFeedbackBufferRange = (C.GPTRANSFORMFEEDBACKBUFFERRANGE)(getProcAddr("glTransformFeedbackBufferRange"))
	if gpTransformFeedbackBufferRange == nil {
		return errors.New("glTransformFeedbackBufferRange")
	}
	gpTransformFeedbackVaryings = (C.GPTRANSFORMFEEDBACKVARYINGS)(getProcAddr("glTransformFeedbackVaryings"))
	if gpTransformFeedbackVaryings == nil {
		return errors.New("glTransformFeedbackVaryings")
	}
	gpUniform1d = (C.GPUNIFORM1D)(getProcAddr("glUniform1d"))
	if gpUniform1d == nil {
		return errors.New("glUniform1d")
	}
	gpUniform1dv = (C.GPUNIFORM1DV)(getProcAddr("glUniform1dv"))
	if gpUniform1dv == nil {
		return errors.New("glUniform1dv")
	}
	gpUniform1f = (C.GPUNIFORM1F)(getProcAddr("glUniform1f"))
	if gpUniform1f == nil {
		return errors.New("glUniform1f")
	}
	gpUniform1fv = (C.GPUNIFORM1FV)(getProcAddr("glUniform1fv"))
	if gpUniform1fv == nil {
		return errors.New("glUniform1fv")
	}
	gpUniform1i = (C.GPUNIFORM1I)(getProcAddr("glUniform1i"))
	if gpUniform1i == nil {
		return errors.New("glUniform1i")
	}
	gpUniform1iv = (C.GPUNIFORM1IV)(getProcAddr("glUniform1iv"))
	if gpUniform1iv == nil {
		return errors.New("glUniform1iv")
	}
	gpUniform1ui = (C.GPUNIFORM1UI)(getProcAddr("glUniform1ui"))
	if gpUniform1ui == nil {
		return errors.New("glUniform1ui")
	}
	gpUniform1uiv = (C.GPUNIFORM1UIV)(getProcAddr("glUniform1uiv"))
	if gpUniform1uiv == nil {
		return errors.New("glUniform1uiv")
	}
	gpUniform2d = (C.GPUNIFORM2D)(getProcAddr("glUniform2d"))
	if gpUniform2d == nil {
		return errors.New("glUniform2d")
	}
	gpUniform2dv = (C.GPUNIFORM2DV)(getProcAddr("glUniform2dv"))
	if gpUniform2dv == nil {
		return errors.New("glUniform2dv")
	}
	gpUniform2f = (C.GPUNIFORM2F)(getProcAddr("glUniform2f"))
	if gpUniform2f == nil {
		return errors.New("glUniform2f")
	}
	gpUniform2fv = (C.GPUNIFORM2FV)(getProcAddr("glUniform2fv"))
	if gpUniform2fv == nil {
		return errors.New("glUniform2fv")
	}
	gpUniform2i = (C.GPUNIFORM2I)(getProcAddr("glUniform2i"))
	if gpUniform2i == nil {
		return errors.New("glUniform2i")
	}
	gpUniform2iv = (C.GPUNIFORM2IV)(getProcAddr("glUniform2iv"))
	if gpUniform2iv == nil {
		return errors.New("glUniform2iv")
	}
	gpUniform2ui = (C.GPUNIFORM2UI)(getProcAddr("glUniform2ui"))
	if gpUniform2ui == nil {
		return errors.New("glUniform2ui")
	}
	gpUniform2uiv = (C.GPUNIFORM2UIV)(getProcAddr("glUniform2uiv"))
	if gpUniform2uiv == nil {
		return errors.New("glUniform2uiv")
	}
	gpUniform3d = (C.GPUNIFORM3D)(getProcAddr("glUniform3d"))
	if gpUniform3d == nil {
		return errors.New("glUniform3d")
	}
	gpUniform3dv = (C.GPUNIFORM3DV)(getProcAddr("glUniform3dv"))
	if gpUniform3dv == nil {
		return errors.New("glUniform3dv")
	}
	gpUniform3f = (C.GPUNIFORM3F)(getProcAddr("glUniform3f"))
	if gpUniform3f == nil {
		return errors.New("glUniform3f")
	}
	gpUniform3fv = (C.GPUNIFORM3FV)(getProcAddr("glUniform3fv"))
	if gpUniform3fv == nil {
		return errors.New("glUniform3fv")
	}
	gpUniform3i = (C.GPUNIFORM3I)(getProcAddr("glUniform3i"))
	if gpUniform3i == nil {
		return errors.New("glUniform3i")
	}
	gpUniform3iv = (C.GPUNIFORM3IV)(getProcAddr("glUniform3iv"))
	if gpUniform3iv == nil {
		return errors.New("glUniform3iv")
	}
	gpUniform3ui = (C.GPUNIFORM3UI)(getProcAddr("glUniform3ui"))
	if gpUniform3ui == nil {
		return errors.New("glUniform3ui")
	}
	gpUniform3uiv = (C.GPUNIFORM3UIV)(getProcAddr("glUniform3uiv"))
	if gpUniform3uiv == nil {
		return errors.New("glUniform3uiv")
	}
	gpUniform4d = (C.GPUNIFORM4D)(getProcAddr("glUniform4d"))
	if gpUniform4d == nil {
		return errors.New("glUniform4d")
	}
	gpUniform4dv = (C.GPUNIFORM4DV)(getProcAddr("glUniform4dv"))
	if gpUniform4dv == nil {
		return errors.New("glUniform4dv")
	}
	gpUniform4f = (C.GPUNIFORM4F)(getProcAddr("glUniform4f"))
	if gpUniform4f == nil {
		return errors.New("glUniform4f")
	}
	gpUniform4fv = (C.GPUNIFORM4FV)(getProcAddr("glUniform4fv"))
	if gpUniform4fv == nil {
		return errors.New("glUniform4fv")
	}
	gpUniform4i = (C.GPUNIFORM4I)(getProcAddr("glUniform4i"))
	if gpUniform4i == nil {
		return errors.New("glUniform4i")
	}
	gpUniform4iv = (C.GPUNIFORM4IV)(getProcAddr("glUniform4iv"))
	if gpUniform4iv == nil {
		return errors.New("glUniform4iv")
	}
	gpUniform4ui = (C.GPUNIFORM4UI)(getProcAddr("glUniform4ui"))
	if gpUniform4ui == nil {
		return errors.New("glUniform4ui")
	}
	gpUniform4uiv = (C.GPUNIFORM4UIV)(getProcAddr("glUniform4uiv"))
	if gpUniform4uiv == nil {
		return errors.New("glUniform4uiv")
	}
	gpUniformBlockBinding = (C.GPUNIFORMBLOCKBINDING)(getProcAddr("glUniformBlockBinding"))
	if gpUniformBlockBinding == nil {
		return errors.New("glUniformBlockBinding")
	}
	gpUniformHandleui64ARB = (C.GPUNIFORMHANDLEUI64ARB)(getProcAddr("glUniformHandleui64ARB"))
	gpUniformHandleui64vARB = (C.GPUNIFORMHANDLEUI64VARB)(getProcAddr("glUniformHandleui64vARB"))
	gpUniformMatrix2dv = (C.GPUNIFORMMATRIX2DV)(getProcAddr("glUniformMatrix2dv"))
	if gpUniformMatrix2dv == nil {
		return errors.New("glUniformMatrix2dv")
	}
	gpUniformMatrix2fv = (C.GPUNIFORMMATRIX2FV)(getProcAddr("glUniformMatrix2fv"))
	if gpUniformMatrix2fv == nil {
		return errors.New("glUniformMatrix2fv")
	}
	gpUniformMatrix2x3dv = (C.GPUNIFORMMATRIX2X3DV)(getProcAddr("glUniformMatrix2x3dv"))
	if gpUniformMatrix2x3dv == nil {
		return errors.New("glUniformMatrix2x3dv")
	}
	gpUniformMatrix2x3fv = (C.GPUNIFORMMATRIX2X3FV)(getProcAddr("glUniformMatrix2x3fv"))
	if gpUniformMatrix2x3fv == nil {
		return errors.New("glUniformMatrix2x3fv")
	}
	gpUniformMatrix2x4dv = (C.GPUNIFORMMATRIX2X4DV)(getProcAddr("glUniformMatrix2x4dv"))
	if gpUniformMatrix2x4dv == nil {
		return errors.New("glUniformMatrix2x4dv")
	}
	gpUniformMatrix2x4fv = (C.GPUNIFORMMATRIX2X4FV)(getProcAddr("glUniformMatrix2x4fv"))
	if gpUniformMatrix2x4fv == nil {
		return errors.New("glUniformMatrix2x4fv")
	}
	gpUniformMatrix3dv = (C.GPUNIFORMMATRIX3DV)(getProcAddr("glUniformMatrix3dv"))
	if gpUniformMatrix3dv == nil {
		return errors.New("glUniformMatrix3dv")
	}
	gpUniformMatrix3fv = (C.GPUNIFORMMATRIX3FV)(getProcAddr("glUniformMatrix3fv"))
	if gpUniformMatrix3fv == nil {
		return errors.New("glUniformMatrix3fv")
	}
	gpUniformMatrix3x2dv = (C.GPUNIFORMMATRIX3X2DV)(getProcAddr("glUniformMatrix3x2dv"))
	if gpUniformMatrix3x2dv == nil {
		return errors.New("glUniformMatrix3x2dv")
	}
	gpUniformMatrix3x2fv = (C.GPUNIFORMMATRIX3X2FV)(getProcAddr("glUniformMatrix3x2fv"))
	if gpUniformMatrix3x2fv == nil {
		return errors.New("glUniformMatrix3x2fv")
	}
	gpUniformMatrix3x4dv = (C.GPUNIFORMMATRIX3X4DV)(getProcAddr("glUniformMatrix3x4dv"))
	if gpUniformMatrix3x4dv == nil {
		return errors.New("glUniformMatrix3x4dv")
	}
	gpUniformMatrix3x4fv = (C.GPUNIFORMMATRIX3X4FV)(getProcAddr("glUniformMatrix3x4fv"))
	if gpUniformMatrix3x4fv == nil {
		return errors.New("glUniformMatrix3x4fv")
	}
	gpUniformMatrix4dv = (C.GPUNIFORMMATRIX4DV)(getProcAddr("glUniformMatrix4dv"))
	if gpUniformMatrix4dv == nil {
		return errors.New("glUniformMatrix4dv")
	}
	gpUniformMatrix4fv = (C.GPUNIFORMMATRIX4FV)(getProcAddr("glUniformMatrix4fv"))
	if gpUniformMatrix4fv == nil {
		return errors.New("glUniformMatrix4fv")
	}
	gpUniformMatrix4x2dv = (C.GPUNIFORMMATRIX4X2DV)(getProcAddr("glUniformMatrix4x2dv"))
	if gpUniformMatrix4x2dv == nil {
		return errors.New("glUniformMatrix4x2dv")
	}
	gpUniformMatrix4x2fv = (C.GPUNIFORMMATRIX4X2FV)(getProcAddr("glUniformMatrix4x2fv"))
	if gpUniformMatrix4x2fv == nil {
		return errors.New("glUniformMatrix4x2fv")
	}
	gpUniformMatrix4x3dv = (C.GPUNIFORMMATRIX4X3DV)(getProcAddr("glUniformMatrix4x3dv"))
	if gpUniformMatrix4x3dv == nil {
		return errors.New("glUniformMatrix4x3dv")
	}
	gpUniformMatrix4x3fv = (C.GPUNIFORMMATRIX4X3FV)(getProcAddr("glUniformMatrix4x3fv"))
	if gpUniformMatrix4x3fv == nil {
		return errors.New("glUniformMatrix4x3fv")
	}
	gpUniformSubroutinesuiv = (C.GPUNIFORMSUBROUTINESUIV)(getProcAddr("glUniformSubroutinesuiv"))
	if gpUniformSubroutinesuiv == nil {
		return errors.New("glUniformSubroutinesuiv")
	}
	gpUnmapBuffer = (C.GPUNMAPBUFFER)(getProcAddr("glUnmapBuffer"))
	if gpUnmapBuffer == nil {
		return errors.New("glUnmapBuffer")
	}
	gpUnmapNamedBuffer = (C.GPUNMAPNAMEDBUFFER)(getProcAddr("glUnmapNamedBuffer"))
	if gpUnmapNamedBuffer == nil {
		return errors.New("glUnmapNamedBuffer")
	}
	gpUseProgram = (C.GPUSEPROGRAM)(getProcAddr("glUseProgram"))
	if gpUseProgram == nil {
		return errors.New("glUseProgram")
	}
	gpUseProgramStages = (C.GPUSEPROGRAMSTAGES)(getProcAddr("glUseProgramStages"))
	if gpUseProgramStages == nil {
		return errors.New("glUseProgramStages")
	}
	gpValidateProgram = (C.GPVALIDATEPROGRAM)(getProcAddr("glValidateProgram"))
	if gpValidateProgram == nil {
		return errors.New("glValidateProgram")
	}
	gpValidateProgramPipeline = (C.GPVALIDATEPROGRAMPIPELINE)(getProcAddr("glValidateProgramPipeline"))
	if gpValidateProgramPipeline == nil {
		return errors.New("glValidateProgramPipeline")
	}
	gpVertexArrayAttribBinding = (C.GPVERTEXARRAYATTRIBBINDING)(getProcAddr("glVertexArrayAttribBinding"))
	if gpVertexArrayAttribBinding == nil {
		return errors.New("glVertexArrayAttribBinding")
	}
	gpVertexArrayAttribFormat = (C.GPVERTEXARRAYATTRIBFORMAT)(getProcAddr("glVertexArrayAttribFormat"))
	if gpVertexArrayAttribFormat == nil {
		return errors.New("glVertexArrayAttribFormat")
	}
	gpVertexArrayAttribIFormat = (C.GPVERTEXARRAYATTRIBIFORMAT)(getProcAddr("glVertexArrayAttribIFormat"))
	if gpVertexArrayAttribIFormat == nil {
		return errors.New("glVertexArrayAttribIFormat")
	}
	gpVertexArrayAttribLFormat = (C.GPVERTEXARRAYATTRIBLFORMAT)(getProcAddr("glVertexArrayAttribLFormat"))
	if gpVertexArrayAttribLFormat == nil {
		return errors.New("glVertexArrayAttribLFormat")
	}
	gpVertexArrayBindingDivisor = (C.GPVERTEXARRAYBINDINGDIVISOR)(getProcAddr("glVertexArrayBindingDivisor"))
	if gpVertexArrayBindingDivisor == nil {
		return errors.New("glVertexArrayBindingDivisor")
	}
	gpVertexArrayElementBuffer = (C.GPVERTEXARRAYELEMENTBUFFER)(getProcAddr("glVertexArrayElementBuffer"))
	if gpVertexArrayElementBuffer == nil {
		return errors.New("glVertexArrayElementBuffer")
	}
	gpVertexArrayVertexBuffer = (C.GPVERTEXARRAYVERTEXBUFFER)(getProcAddr("glVertexArrayVertexBuffer"))
	if gpVertexArrayVertexBuffer == nil {
		return errors.New("glVertexArrayVertexBuffer")
	}
	gpVertexArrayVertexBuffers = (C.GPVERTEXARRAYVERTEXBUFFERS)(getProcAddr("glVertexArrayVertexBuffers"))
	if gpVertexArrayVertexBuffers == nil {
		return errors.New("glVertexArrayVertexBuffers")
	}
	gpVertexAttrib1d = (C.GPVERTEXATTRIB1D)(getProcAddr("glVertexAttrib1d"))
	if gpVertexAttrib1d == nil {
		return errors.New("glVertexAttrib1d")
	}
	gpVertexAttrib1dv = (C.GPVERTEXATTRIB1DV)(getProcAddr("glVertexAttrib1dv"))
	if gpVertexAttrib1dv == nil {
		return errors.New("glVertexAttrib1dv")
	}
	gpVertexAttrib1f = (C.GPVERTEXATTRIB1F)(getProcAddr("glVertexAttrib1f"))
	if gpVertexAttrib1f == nil {
		return errors.New("glVertexAttrib1f")
	}
	gpVertexAttrib1fv = (C.GPVERTEXATTRIB1FV)(getProcAddr("glVertexAttrib1fv"))
	if gpVertexAttrib1fv == nil {
		return errors.New("glVertexAttrib1fv")
	}
	gpVertexAttrib1s = (C.GPVERTEXATTRIB1S)(getProcAddr("glVertexAttrib1s"))
	if gpVertexAttrib1s == nil {
		return errors.New("glVertexAttrib1s")
	}
	gpVertexAttrib1sv = (C.GPVERTEXATTRIB1SV)(getProcAddr("glVertexAttrib1sv"))
	if gpVertexAttrib1sv == nil {
		return errors.New("glVertexAttrib1sv")
	}
	gpVertexAttrib2d = (C.GPVERTEXATTRIB2D)(getProcAddr("glVertexAttrib2d"))
	if gpVertexAttrib2d == nil {
		return errors.New("glVertexAttrib2d")
	}
	gpVertexAttrib2dv = (C.GPVERTEXATTRIB2DV)(getProcAddr("glVertexAttrib2dv"))
	if gpVertexAttrib2dv == nil {
		return errors.New("glVertexAttrib2dv")
	}
	gpVertexAttrib2f = (C.GPVERTEXATTRIB2F)(getProcAddr("glVertexAttrib2f"))
	if gpVertexAttrib2f == nil {
		return errors.New("glVertexAttrib2f")
	}
	gpVertexAttrib2fv = (C.GPVERTEXATTRIB2FV)(getProcAddr("glVertexAttrib2fv"))
	if gpVertexAttrib2fv == nil {
		return errors.New("glVertexAttrib2fv")
	}
	gpVertexAttrib2s = (C.GPVERTEXATTRIB2S)(getProcAddr("glVertexAttrib2s"))
	if gpVertexAttrib2s == nil {
		return errors.New("glVertexAttrib2s")
	}
	gpVertexAttrib2sv = (C.GPVERTEXATTRIB2SV)(getProcAddr("glVertexAttrib2sv"))
	if gpVertexAttrib2sv == nil {
		return errors.New("glVertexAttrib2sv")
	}
	gpVertexAttrib3d = (C.GPVERTEXATTRIB3D)(getProcAddr("glVertexAttrib3d"))
	if gpVertexAttrib3d == nil {
		return errors.New("glVertexAttrib3d")
	}
	gpVertexAttrib3dv = (C.GPVERTEXATTRIB3DV)(getProcAddr("glVertexAttrib3dv"))
	if gpVertexAttrib3dv == nil {
		return errors.New("glVertexAttrib3dv")
	}
	gpVertexAttrib3f = (C.GPVERTEXATTRIB3F)(getProcAddr("glVertexAttrib3f"))
	if gpVertexAttrib3f == nil {
		return errors.New("glVertexAttrib3f")
	}
	gpVertexAttrib3fv = (C.GPVERTEXATTRIB3FV)(getProcAddr("glVertexAttrib3fv"))
	if gpVertexAttrib3fv == nil {
		return errors.New("glVertexAttrib3fv")
	}
	gpVertexAttrib3s = (C.GPVERTEXATTRIB3S)(getProcAddr("glVertexAttrib3s"))
	if gpVertexAttrib3s == nil {
		return errors.New("glVertexAttrib3s")
	}
	gpVertexAttrib3sv = (C.GPVERTEXATTRIB3SV)(getProcAddr("glVertexAttrib3sv"))
	if gpVertexAttrib3sv == nil {
		return errors.New("glVertexAttrib3sv")
	}
	gpVertexAttrib4Nbv = (C.GPVERTEXATTRIB4NBV)(getProcAddr("glVertexAttrib4Nbv"))
	if gpVertexAttrib4Nbv == nil {
		return errors.New("glVertexAttrib4Nbv")
	}
	gpVertexAttrib4Niv = (C.GPVERTEXATTRIB4NIV)(getProcAddr("glVertexAttrib4Niv"))
	if gpVertexAttrib4Niv == nil {
		return errors.New("glVertexAttrib4Niv")
	}
	gpVertexAttrib4Nsv = (C.GPVERTEXATTRIB4NSV)(getProcAddr("glVertexAttrib4Nsv"))
	if gpVertexAttrib4Nsv == nil {
		return errors.New("glVertexAttrib4Nsv")
	}
	gpVertexAttrib4Nub = (C.GPVERTEXATTRIB4NUB)(getProcAddr("glVertexAttrib4Nub"))
	if gpVertexAttrib4Nub == nil {
		return errors.New("glVertexAttrib4Nub")
	}
	gpVertexAttrib4Nubv = (C.GPVERTEXATTRIB4NUBV)(getProcAddr("glVertexAttrib4Nubv"))
	if gpVertexAttrib4Nubv == nil {
		return errors.New("glVertexAttrib4Nubv")
	}
	gpVertexAttrib4Nuiv = (C.GPVERTEXATTRIB4NUIV)(getProcAddr("glVertexAttrib4Nuiv"))
	if gpVertexAttrib4Nuiv == nil {
		return errors.New("glVertexAttrib4Nuiv")
	}
	gpVertexAttrib4Nusv = (C.GPVERTEXATTRIB4NUSV)(getProcAddr("glVertexAttrib4Nusv"))
	if gpVertexAttrib4Nusv == nil {
		return errors.New("glVertexAttrib4Nusv")
	}
	gpVertexAttrib4bv = (C.GPVERTEXATTRIB4BV)(getProcAddr("glVertexAttrib4bv"))
	if gpVertexAttrib4bv == nil {
		return errors.New("glVertexAttrib4bv")
	}
	gpVertexAttrib4d = (C.GPVERTEXATTRIB4D)(getProcAddr("glVertexAttrib4d"))
	if gpVertexAttrib4d == nil {
		return errors.New("glVertexAttrib4d")
	}
	gpVertexAttrib4dv = (C.GPVERTEXATTRIB4DV)(getProcAddr("glVertexAttrib4dv"))
	if gpVertexAttrib4dv == nil {
		return errors.New("glVertexAttrib4dv")
	}
	gpVertexAttrib4f = (C.GPVERTEXATTRIB4F)(getProcAddr("glVertexAttrib4f"))
	if gpVertexAttrib4f == nil {
		return errors.New("glVertexAttrib4f")
	}
	gpVertexAttrib4fv = (C.GPVERTEXATTRIB4FV)(getProcAddr("glVertexAttrib4fv"))
	if gpVertexAttrib4fv == nil {
		return errors.New("glVertexAttrib4fv")
	}
	gpVertexAttrib4iv = (C.GPVERTEXATTRIB4IV)(getProcAddr("glVertexAttrib4iv"))
	if gpVertexAttrib4iv == nil {
		return errors.New("glVertexAttrib4iv")
	}
	gpVertexAttrib4s = (C.GPVERTEXATTRIB4S)(getProcAddr("glVertexAttrib4s"))
	if gpVertexAttrib4s == nil {
		return errors.New("glVertexAttrib4s")
	}
	gpVertexAttrib4sv = (C.GPVERTEXATTRIB4SV)(getProcAddr("glVertexAttrib4sv"))
	if gpVertexAttrib4sv == nil {
		return errors.New("glVertexAttrib4sv")
	}
	gpVertexAttrib4ubv = (C.GPVERTEXATTRIB4UBV)(getProcAddr("glVertexAttrib4ubv"))
	if gpVertexAttrib4ubv == nil {
		return errors.New("glVertexAttrib4ubv")
	}
	gpVertexAttrib4uiv = (C.GPVERTEXATTRIB4UIV)(getProcAddr("glVertexAttrib4uiv"))
	if gpVertexAttrib4uiv == nil {
		return errors.New("glVertexAttrib4uiv")
	}
	gpVertexAttrib4usv = (C.GPVERTEXATTRIB4USV)(getProcAddr("glVertexAttrib4usv"))
	if gpVertexAttrib4usv == nil {
		return errors.New("glVertexAttrib4usv")
	}
	gpVertexAttribBinding = (C.GPVERTEXATTRIBBINDING)(getProcAddr("glVertexAttribBinding"))
	if gpVertexAttribBinding == nil {
		return errors.New("glVertexAttribBinding")
	}
	gpVertexAttribDivisor = (C.GPVERTEXATTRIBDIVISOR)(getProcAddr("glVertexAttribDivisor"))
	if gpVertexAttribDivisor == nil {
		return errors.New("glVertexAttribDivisor")
	}
	gpVertexAttribFormat = (C.GPVERTEXATTRIBFORMAT)(getProcAddr("glVertexAttribFormat"))
	if gpVertexAttribFormat == nil {
		return errors.New("glVertexAttribFormat")
	}
	gpVertexAttribI1i = (C.GPVERTEXATTRIBI1I)(getProcAddr("glVertexAttribI1i"))
	if gpVertexAttribI1i == nil {
		return errors.New("glVertexAttribI1i")
	}
	gpVertexAttribI1iv = (C.GPVERTEXATTRIBI1IV)(getProcAddr("glVertexAttribI1iv"))
	if gpVertexAttribI1iv == nil {
		return errors.New("glVertexAttribI1iv")
	}
	gpVertexAttribI1ui = (C.GPVERTEXATTRIBI1UI)(getProcAddr("glVertexAttribI1ui"))
	if gpVertexAttribI1ui == nil {
		return errors.New("glVertexAttribI1ui")
	}
	gpVertexAttribI1uiv = (C.GPVERTEXATTRIBI1UIV)(getProcAddr("glVertexAttribI1uiv"))
	if gpVertexAttribI1uiv == nil {
		return errors.New("glVertexAttribI1uiv")
	}
	gpVertexAttribI2i = (C.GPVERTEXATTRIBI2I)(getProcAddr("glVertexAttribI2i"))
	if gpVertexAttribI2i == nil {
		return errors.New("glVertexAttribI2i")
	}
	gpVertexAttribI2iv = (C.GPVERTEXATTRIBI2IV)(getProcAddr("glVertexAttribI2iv"))
	if gpVertexAttribI2iv == nil {
		return errors.New("glVertexAttribI2iv")
	}
	gpVertexAttribI2ui = (C.GPVERTEXATTRIBI2UI)(getProcAddr("glVertexAttribI2ui"))
	if gpVertexAttribI2ui == nil {
		return errors.New("glVertexAttribI2ui")
	}
	gpVertexAttribI2uiv = (C.GPVERTEXATTRIBI2UIV)(getProcAddr("glVertexAttribI2uiv"))
	if gpVertexAttribI2uiv == nil {
		return errors.New("glVertexAttribI2uiv")
	}
	gpVertexAttribI3i = (C.GPVERTEXATTRIBI3I)(getProcAddr("glVertexAttribI3i"))
	if gpVertexAttribI3i == nil {
		return errors.New("glVertexAttribI3i")
	}
	gpVertexAttribI3iv = (C.GPVERTEXATTRIBI3IV)(getProcAddr("glVertexAttribI3iv"))
	if gpVertexAttribI3iv == nil {
		return errors.New("glVertexAttribI3iv")
	}
	gpVertexAttribI3ui = (C.GPVERTEXATTRIBI3UI)(getProcAddr("glVertexAttribI3ui"))
	if gpVertexAttribI3ui == nil {
		return errors.New("glVertexAttribI3ui")
	}
	gpVertexAttribI3uiv = (C.GPVERTEXATTRIBI3UIV)(getProcAddr("glVertexAttribI3uiv"))
	if gpVertexAttribI3uiv == nil {
		return errors.New("glVertexAttribI3uiv")
	}
	gpVertexAttribI4bv = (C.GPVERTEXATTRIBI4BV)(getProcAddr("glVertexAttribI4bv"))
	if gpVertexAttribI4bv == nil {
		return errors.New("glVertexAttribI4bv")
	}
	gpVertexAttribI4i = (C.GPVERTEXATTRIBI4I)(getProcAddr("glVertexAttribI4i"))
	if gpVertexAttribI4i == nil {
		return errors.New("glVertexAttribI4i")
	}
	gpVertexAttribI4iv = (C.GPVERTEXATTRIBI4IV)(getProcAddr("glVertexAttribI4iv"))
	if gpVertexAttribI4iv == nil {
		return errors.New("glVertexAttribI4iv")
	}
	gpVertexAttribI4sv = (C.GPVERTEXATTRIBI4SV)(getProcAddr("glVertexAttribI4sv"))
	if gpVertexAttribI4sv == nil {
		return errors.New("glVertexAttribI4sv")
	}
	gpVertexAttribI4ubv = (C.GPVERTEXATTRIBI4UBV)(getProcAddr("glVertexAttribI4ubv"))
	if gpVertexAttribI4ubv == nil {
		return errors.New("glVertexAttribI4ubv")
	}
	gpVertexAttribI4ui = (C.GPVERTEXATTRIBI4UI)(getProcAddr("glVertexAttribI4ui"))
	if gpVertexAttribI4ui == nil {
		return errors.New("glVertexAttribI4ui")
	}
	gpVertexAttribI4uiv = (C.GPVERTEXATTRIBI4UIV)(getProcAddr("glVertexAttribI4uiv"))
	if gpVertexAttribI4uiv == nil {
		return errors.New("glVertexAttribI4uiv")
	}
	gpVertexAttribI4usv = (C.GPVERTEXATTRIBI4USV)(getProcAddr("glVertexAttribI4usv"))
	if gpVertexAttribI4usv == nil {
		return errors.New("glVertexAttribI4usv")
	}
	gpVertexAttribIFormat = (C.GPVERTEXATTRIBIFORMAT)(getProcAddr("glVertexAttribIFormat"))
	if gpVertexAttribIFormat == nil {
		return errors.New("glVertexAttribIFormat")
	}
	gpVertexAttribIPointer = (C.GPVERTEXATTRIBIPOINTER)(getProcAddr("glVertexAttribIPointer"))
	if gpVertexAttribIPointer == nil {
		return errors.New("glVertexAttribIPointer")
	}
	gpVertexAttribL1d = (C.GPVERTEXATTRIBL1D)(getProcAddr("glVertexAttribL1d"))
	if gpVertexAttribL1d == nil {
		return errors.New("glVertexAttribL1d")
	}
	gpVertexAttribL1dv = (C.GPVERTEXATTRIBL1DV)(getProcAddr("glVertexAttribL1dv"))
	if gpVertexAttribL1dv == nil {
		return errors.New("glVertexAttribL1dv")
	}
	gpVertexAttribL1ui64ARB = (C.GPVERTEXATTRIBL1UI64ARB)(getProcAddr("glVertexAttribL1ui64ARB"))
	gpVertexAttribL1ui64vARB = (C.GPVERTEXATTRIBL1UI64VARB)(getProcAddr("glVertexAttribL1ui64vARB"))
	gpVertexAttribL2d = (C.GPVERTEXATTRIBL2D)(getProcAddr("glVertexAttribL2d"))
	if gpVertexAttribL2d == nil {
		return errors.New("glVertexAttribL2d")
	}
	gpVertexAttribL2dv = (C.GPVERTEXATTRIBL2DV)(getProcAddr("glVertexAttribL2dv"))
	if gpVertexAttribL2dv == nil {
		return errors.New("glVertexAttribL2dv")
	}
	gpVertexAttribL3d = (C.GPVERTEXATTRIBL3D)(getProcAddr("glVertexAttribL3d"))
	if gpVertexAttribL3d == nil {
		return errors.New("glVertexAttribL3d")
	}
	gpVertexAttribL3dv = (C.GPVERTEXATTRIBL3DV)(getProcAddr("glVertexAttribL3dv"))
	if gpVertexAttribL3dv == nil {
		return errors.New("glVertexAttribL3dv")
	}
	gpVertexAttribL4d = (C.GPVERTEXATTRIBL4D)(getProcAddr("glVertexAttribL4d"))
	if gpVertexAttribL4d == nil {
		return errors.New("glVertexAttribL4d")
	}
	gpVertexAttribL4dv = (C.GPVERTEXATTRIBL4DV)(getProcAddr("glVertexAttribL4dv"))
	if gpVertexAttribL4dv == nil {
		return errors.New("glVertexAttribL4dv")
	}
	gpVertexAttribLFormat = (C.GPVERTEXATTRIBLFORMAT)(getProcAddr("glVertexAttribLFormat"))
	if gpVertexAttribLFormat == nil {
		return errors.New("glVertexAttribLFormat")
	}
	gpVertexAttribLPointer = (C.GPVERTEXATTRIBLPOINTER)(getProcAddr("glVertexAttribLPointer"))
	if gpVertexAttribLPointer == nil {
		return errors.New("glVertexAttribLPointer")
	}
	gpVertexAttribP1ui = (C.GPVERTEXATTRIBP1UI)(getProcAddr("glVertexAttribP1ui"))
	if gpVertexAttribP1ui == nil {
		return errors.New("glVertexAttribP1ui")
	}
	gpVertexAttribP1uiv = (C.GPVERTEXATTRIBP1UIV)(getProcAddr("glVertexAttribP1uiv"))
	if gpVertexAttribP1uiv == nil {
		return errors.New("glVertexAttribP1uiv")
	}
	gpVertexAttribP2ui = (C.GPVERTEXATTRIBP2UI)(getProcAddr("glVertexAttribP2ui"))
	if gpVertexAttribP2ui == nil {
		return errors.New("glVertexAttribP2ui")
	}
	gpVertexAttribP2uiv = (C.GPVERTEXATTRIBP2UIV)(getProcAddr("glVertexAttribP2uiv"))
	if gpVertexAttribP2uiv == nil {
		return errors.New("glVertexAttribP2uiv")
	}
	gpVertexAttribP3ui = (C.GPVERTEXATTRIBP3UI)(getProcAddr("glVertexAttribP3ui"))
	if gpVertexAttribP3ui == nil {
		return errors.New("glVertexAttribP3ui")
	}
	gpVertexAttribP3uiv = (C.GPVERTEXATTRIBP3UIV)(getProcAddr("glVertexAttribP3uiv"))
	if gpVertexAttribP3uiv == nil {
		return errors.New("glVertexAttribP3uiv")
	}
	gpVertexAttribP4ui = (C.GPVERTEXATTRIBP4UI)(getProcAddr("glVertexAttribP4ui"))
	if gpVertexAttribP4ui == nil {
		return errors.New("glVertexAttribP4ui")
	}
	gpVertexAttribP4uiv = (C.GPVERTEXATTRIBP4UIV)(getProcAddr("glVertexAttribP4uiv"))
	if gpVertexAttribP4uiv == nil {
		return errors.New("glVertexAttribP4uiv")
	}
	gpVertexAttribPointer = (C.GPVERTEXATTRIBPOINTER)(getProcAddr("glVertexAttribPointer"))
	if gpVertexAttribPointer == nil {
		return errors.New("glVertexAttribPointer")
	}
	gpVertexBindingDivisor = (C.GPVERTEXBINDINGDIVISOR)(getProcAddr("glVertexBindingDivisor"))
	if gpVertexBindingDivisor == nil {
		return errors.New("glVertexBindingDivisor")
	}
	gpViewport = (C.GPVIEWPORT)(getProcAddr("glViewport"))
	if gpViewport == nil {
		return errors.New("glViewport")
	}
	gpViewportArrayv = (C.GPVIEWPORTARRAYV)(getProcAddr("glViewportArrayv"))
	if gpViewportArrayv == nil {
		return errors.New("glViewportArrayv")
	}
	gpViewportIndexedf = (C.GPVIEWPORTINDEXEDF)(getProcAddr("glViewportIndexedf"))
	if gpViewportIndexedf == nil {
		return errors.New("glViewportIndexedf")
	}
	gpViewportIndexedfv = (C.GPVIEWPORTINDEXEDFV)(getProcAddr("glViewportIndexedfv"))
	if gpViewportIndexedfv == nil {
		return errors.New("glViewportIndexedfv")
	}
	gpWaitSync = (C.GPWAITSYNC)(getProcAddr("glWaitSync"))
	if gpWaitSync == nil {
		return errors.New("glWaitSync")
	}
	if len(missList) > 0 {
		fmt.Println("GL API Missing List: ", missList)
	}
	return nil
}

func getProcAddress(namea string) unsafe.Pointer {
	cname := C.CString(namea)
	defer C.free(unsafe.Pointer(cname))
	return C.ProcAddress(cname)
}
