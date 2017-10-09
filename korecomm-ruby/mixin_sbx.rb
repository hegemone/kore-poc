class Engine
  def handle_egress(msg)
    puts "Engine#handle_egress -> [#{msg}]"
  end
end

module Kore
  module Comm
    module Plugin
      module ClassMethods
        attr_reader :matchers
        def match(pattern, opts)
          @matchers[pattern] = opts[:method]
        end
        def self.extended(base)
          base.instance_exec do
            @matchers = {}
          end
        end
      end
      module InstanceMethods
        attr_accessor :engine
        def handle_ingress(msg)
          self.class.matchers.each do |pattern, f|
            if !!pattern.match(msg)
              method(f).call(msg)
            end
          end
        end
        def send_message(msg)
          @engine.handle_egress(msg)
        end
      end
      def self.included(base)
        base.send :include, InstanceMethods
        base.extend ClassMethods
      end
    end
  end
end

class FooPlugin
  include Kore::Comm::Plugin
  match /bacon/i, :method => :command_bacon
  def command_bacon(arg)
    self.send_message 'Send bacon!'
  end
end

foo = FooPlugin.new
foo.engine = Engine.new
foo.handle_ingress("bacon foo")
